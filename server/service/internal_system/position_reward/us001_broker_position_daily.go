package position_reward

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBrokerPositionDaily
//@description: 创建BrokerPositionDaily记录
//@param: BrokerPositionDaily model.BrokerPositionDaily
//@return: err error

func CreateBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Create(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBrokerPositionDaily
//@description: 删除BrokerPositionDaily记录
//@param: BrokerPositionDaily model.BrokerPositionDaily
//@return: err error

func DeleteBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBrokerPositionDailyByIds
//@description: 批量删除BrokerPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBrokerPositionDailyByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.BrokerPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBrokerPositionDaily
//@description: 更新BrokerPositionDaily记录
//@param: BrokerPositionDaily *model.BrokerPositionDaily
//@return: err error

func UpdateBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Save(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBrokerPositionDaily
//@description: 根据id获取BrokerPositionDaily记录
//@param: id uint
//@return: err error, BrokerPositionDaily model.BrokerPositionDaily

func GetBrokerPositionDaily(id uint) (err error, BrokerPositionDaily mp.BrokerPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&BrokerPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBrokerPositionDailyInfoList
//@description: 分页获取BrokerPositionDaily记录
//@param: info request.BrokerPositionDailySearch
//@return: err error, list interface{}, total int64

func GetBrokerPositionDailyInfoList(info rp.BrokerPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.BrokerPositionDaily{}).Preload("ProductInfo").Order("trading_date desc")
	var BrokerPositionDailys []mp.BrokerPositionDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.TradingDate.IsZero() {
		db = db.Where("trading_date = ?", info.TradingDate)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&BrokerPositionDailys).Error
	return err, BrokerPositionDailys, total
}

//批量创建
func CreateBrokerPositionDailyList(list []mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Model(&mp.BrokerPositionDaily{}).CreateInBatches(list, len(list)).Error
	return err
}

func ParseBrokerPositionExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"日期", "期货公司", "品种", "总上缴手续费", "量化客户上缴", "最新日均持仓量", "最新最大可操作量", "最新持仓"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return err
	}
	menus := make([]mp.BrokerPositionDaily, 0)
	err = file.RemoveRow("Sheet1", 1)
	if err != nil {
		return err
	}
	err = file.RemoveCol("Sheet1", "I")
	if err != nil {
		return err
	}
	err = file.RemoveCol("Sheet1", "I")
	if err != nil {
		return err
	}
	if err := file.SaveAs(global.GVA_CONFIG.Excel.Dir + "BrokerPositionDaily.xlsx"); err != nil {
		fmt.Println(err)
	}
	file, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "BrokerPositionDaily.xlsx")
	if err != nil {
		return err
	}
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
		totalTradingFee, _ := strconv.ParseFloat(row[3], 64)
		specialTradingFee, _ := strconv.ParseFloat(row[4], 64)
		averageDailyPosition, _ := strconv.ParseFloat(row[5], 64)
		maximumToOpen, _ := strconv.ParseFloat(row[6], 64)
		currentPosition, _ := strconv.ParseFloat(row[7], 64)
		tradingDate, _ := time.Parse("2006-01-02", row[0])
		if row[1] == "鲁证" {
			row[1] = "0275"
		}
		menu := mp.BrokerPositionDaily{
			TradingDate:          tradingDate,
			BrokerId:             row[1],
			ProductCode:          strings.ToUpper(row[2]),
			TotalTradingFee:      float64(totalTradingFee),
			SpecialTradingFee:    float64(specialTradingFee),
			AverageDailyPosition: float64(averageDailyPosition),
			MaximumToOpen:        float64(maximumToOpen),
			CurrentPosition:      float64(currentPosition),
		}
		menus = append(menus, menu)
	}
	err = CreateBrokerPositionDailyList(menus)
	if err != nil {
		return err
	}
	return err
}

func compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
