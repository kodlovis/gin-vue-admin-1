package position_reward

import (
	"errors"
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAccountPositionDaily
//@description: 创建AccountPositionDaily记录
//@param: accountPositionDaily model.AccountPositionDaily
//@return: err error

func CreateAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Create(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountPositionDaily
//@description: 删除AccountPositionDaily记录
//@param: accountPositionDaily model.AccountPositionDaily
//@return: err error

func DeleteAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountPositionDailyByIds
//@description: 批量删除AccountPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAccountPositionDailyByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.AccountPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAccountPositionDaily
//@description: 更新AccountPositionDaily记录
//@param: accountPositionDaily *model.AccountPositionDaily
//@return: err error

func UpdateAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Save(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountPositionDaily
//@description: 根据id获取AccountPositionDaily记录
//@param: id uint
//@return: err error, accountPositionDaily model.AccountPositionDaily

func GetAccountPositionDaily(id uint) (err error, accountPositionDaily mp.AccountPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&accountPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountPositionDailyInfoList
//@description: 分页获取AccountPositionDaily记录
//@param: info request.AccountPositionDailySearch
//@return: err error, list interface{}, total int64

func GetAccountPositionDailyInfoList(info rp.AccountPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.AccountPositionDaily{}).Preload("AccountInfo").Order("trading_date desc")
	var us001AccountPositionDailys []mp.AccountPositionDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.TradingDate.IsZero() {
		db = db.Where("trading_date = ?", info.TradingDate)
	}
	if info.BrokerId != "" {
		db = db.Where("broker_id LIKE ?", "%"+info.BrokerId+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us001AccountPositionDailys).Error
	return err, us001AccountPositionDailys, total
}
func ParseAccountPositionExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"交易日期", "期货公司", "账号", "合约", "买卖方向", "投机/套保", "手数"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return err
	}
	menus := make([]mp.AccountPositionDaily, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return err
		}
		if skipHeader {
			if compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		amount, _ := strconv.ParseInt(row[6], 0, 0)
		direction, _ := strconv.ParseInt(row[5], 0, 0)
		hedgeFlag, _ := strconv.ParseInt(row[4], 0, 0)

		tradingDate, _ := time.Parse("2006-01-02", row[0])

		if row[1] == "鲁证期货" {
			row[1] = "0275"
		}
		menu := mp.AccountPositionDaily{
			TradingDate: tradingDate,
			BrokerId:    row[1],
			AccountId:   row[2],
			Instrument:  row[3],
			Direction:   int(direction),
			HedgeFlag:   int(hedgeFlag),
			Amount:      int(amount),
		}
		menus = append(menus, menu)
	}
	err = CreateAccountPositionDailyList(menus)
	if err != nil {
		return err
	}
	return err
}

//批量创建
func CreateAccountPositionDailyList(list []mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Model(&mp.AccountPositionDaily{}).CreateInBatches(list, len(list)).Error
	return
}
