package dce

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/position_reward/dce"
	request "gin-vue-admin/model/request/internal_system/position_reward/dce"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs001DceBrokerPositionDaily
//@description: 创建Us001DceBrokerPositionDaily记录
//@param: us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
//@return: err error

func CreateUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Create(&us001DceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs001DceBrokerPositionDaily
//@description: 删除Us001DceBrokerPositionDaily记录
//@param: us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
//@return: err error

func DeleteUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&us001DceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs001DceBrokerPositionDailyByIds
//@description: 批量删除Us001DceBrokerPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs001DceBrokerPositionDailyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Us001DceBrokerPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs001DceBrokerPositionDaily
//@description: 更新Us001DceBrokerPositionDaily记录
//@param: us001DceBrokerPositionDaily *model.Us001DceBrokerPositionDaily
//@return: err error

func UpdateUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Save(&us001DceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs001DceBrokerPositionDaily
//@description: 根据id获取Us001DceBrokerPositionDaily记录
//@param: id uint
//@return: err error, us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily

func GetUs001DceBrokerPositionDaily(id uint) (err error, us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&us001DceBrokerPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs001DceBrokerPositionDailyInfoList
//@description: 分页获取Us001DceBrokerPositionDaily记录
//@param: info request.Us001DceBrokerPositionDailySearch
//@return: err error, list interface{}, total int64

func GetUs001DceBrokerPositionDailyInfoList(info request.Us001DceBrokerPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Us001DceBrokerPositionDaily{}).Order("time desc")
	var us001DceBrokerPositionDailys []model.Us001DceBrokerPositionDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("`time` = ?", info.Time)
	}
	if info.BrokerId != "" {
		db = db.Where("`broker_id` = ?", info.BrokerId)
	}
	if info.ProductCode != "" {
		db = db.Where("`product_code` = ?", info.ProductCode)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us001DceBrokerPositionDailys).Error
	return err, us001DceBrokerPositionDailys, total
}

//批量创建
func CreateUs001DceBrokerPositionDailyList(list []model.Us001DceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Model(&model.Us001DceBrokerPositionDaily{}).CreateInBatches(list, len(list)).Error
	return err
}

func ParseUs001DceBrokerPositionExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"日期", "期货公司", "品种", "比较期一般法人客户做市商合约日均持仓量", "比较期一般法人客户非做市商合约日均持仓量 ", "比较期其他类型客户做市商日均持仓量", "比较期其他类型客户非做市商日均持仓量", "上缴手续费"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return err
	}
	menus := make([]model.Us001DceBrokerPositionDaily, 0)
	// err = file.RemoveRow("Sheet1", 1)
	// if err != nil {
	// 	return err
	// }
	// err = file.RemoveCol("Sheet1", "I")
	// if err != nil {
	// 	return err
	// }
	// err = file.RemoveCol("Sheet1", "I")
	// if err != nil {
	// 	return err
	// }
	if err := file.SaveAs(global.GVA_CONFIG.Excel.Dir + "DceBrokerPositionDaily.xlsx"); err != nil {
		fmt.Println(err)
	}
	file, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "DceBrokerPositionDaily.xlsx")
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
		corporateMarketPosition, _ := strconv.ParseFloat(row[3], 64)
		corporateNonMarketPosition, _ := strconv.ParseFloat(row[4], 64)
		otherMarketPosition, _ := strconv.ParseFloat(row[5], 64)
		otherNonMarketPosition, _ := strconv.ParseFloat(row[6], 64)
		payFee, _ := strconv.ParseFloat(row[7], 64)
		tradingDate, _ := time.Parse("2006-01-02", row[0])
		menu := model.Us001DceBrokerPositionDaily{
			Time:                       tradingDate,
			BrokerId:                   row[1],
			ProductCode:                strings.ToUpper(row[2]),
			CorporateMarketPosition:    float64(corporateMarketPosition),
			CorporateNonMarketPosition: float64(corporateNonMarketPosition),
			OtherMarketPosition:        float64(otherMarketPosition),
			OtherNonMarketPosition:     float64(otherNonMarketPosition),
			PayFee:                     float64(payFee),
		}
		menus = append(menus, menu)
	}
	err = CreateUs001DceBrokerPositionDailyList(menus)
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
