package czce

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/position_reward/czce"
	request "gin-vue-admin/model/request/internal_system/position_reward/czce"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs001CzceBrokerPositionDaily
//@description: 创建Us001CzceBrokerPositionDaily记录
//@param: us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
//@return: err error

func CreateUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Create(&us001CzceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs001CzceBrokerPositionDaily
//@description: 删除Us001CzceBrokerPositionDaily记录
//@param: us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
//@return: err error

func DeleteUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&us001CzceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs001CzceBrokerPositionDailyByIds
//@description: 批量删除Us001CzceBrokerPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs001CzceBrokerPositionDailyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Us001CzceBrokerPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs001CzceBrokerPositionDaily
//@description: 更新Us001CzceBrokerPositionDaily记录
//@param: us001CzceBrokerPositionDaily *model.Us001CzceBrokerPositionDaily
//@return: err error

func UpdateUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Save(&us001CzceBrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs001CzceBrokerPositionDaily
//@description: 根据id获取Us001CzceBrokerPositionDaily记录
//@param: id uint
//@return: err error, us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily

func GetUs001CzceBrokerPositionDaily(id uint) (err error, us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&us001CzceBrokerPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs001CzceBrokerPositionDailyInfoList
//@description: 分页获取Us001CzceBrokerPositionDaily记录
//@param: info request.Us001CzceBrokerPositionDailySearch
//@return: err error, list interface{}, total int64

func GetUs001CzceBrokerPositionDailyInfoList(info request.Us001CzceBrokerPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Us001CzceBrokerPositionDaily{}).Order("time desc")
	var us001CzceBrokerPositionDailys []model.Us001CzceBrokerPositionDaily
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
	err = db.Limit(limit).Offset(offset).Find(&us001CzceBrokerPositionDailys).Error
	for i := 0; i < len(us001CzceBrokerPositionDailys); i++ {
		us001CzceBrokerPositionDailys[i].CustomerDailyPosition = (math.Floor(us001CzceBrokerPositionDailys[i].CustomerDailyPosition*10 + 0.5)) / 10
		us001CzceBrokerPositionDailys[i].InstitutionalDailyPosition = (math.Floor(us001CzceBrokerPositionDailys[i].InstitutionalDailyPosition*10 + 0.5)) / 10
		us001CzceBrokerPositionDailys[i].SpecialLawCustomerDailyPosition = (math.Floor(us001CzceBrokerPositionDailys[i].SpecialLawCustomerDailyPosition*10 + 0.5)) / 10
		us001CzceBrokerPositionDailys[i].CorporateCustomerDailyPosition = (math.Floor(us001CzceBrokerPositionDailys[i].CorporateCustomerDailyPosition*10 + 0.5)) / 10
		us001CzceBrokerPositionDailys[i].DailyAveragePosition = (math.Floor(us001CzceBrokerPositionDailys[i].DailyAveragePosition*10 + 0.5)) / 10
	}
	return err, us001CzceBrokerPositionDailys, total
}

//批量创建
func CreateUs001CzceBrokerPositionDailyList(list []model.Us001CzceBrokerPositionDaily) (err error) {
	err = global.GVA_DB.Model(&model.Us001CzceBrokerPositionDaily{}).CreateInBatches(list, len(list)).Error
	return err
}

func ParseUs001CzceBrokerPositionExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"日期", "期货公司", "代码", "个人客户日均持仓(第二项）", "机构客户日均持仓(第二项）", "特法客户日均持仓(第二项）", "法人客户日均持仓（第三项）", "今日日均持仓(第二项）", "重点合约当月日均持仓(第四项）", "法人持仓占比乘数(第三项）", "本月截止今日上交手续费"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return err
	}
	menus := make([]model.Us001CzceBrokerPositionDaily, 0)
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
	if err := file.SaveAs(global.GVA_CONFIG.Excel.Dir + "CzceBrokerPositionDaily.xlsx"); err != nil {
		fmt.Println(err)
	}
	file, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "CzceBrokerPositionDaily.xlsx")
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
		customerDailyPosition, _ := strconv.ParseFloat(row[3], 64)
		institutionalDailyPosition, _ := strconv.ParseFloat(row[4], 64)
		specialLawCustomerDailyPosition, _ := strconv.ParseFloat(row[5], 64)
		corporateCustomerDailyPosition, _ := strconv.ParseFloat(row[6], 64)
		dailyAveragePosition, _ := strconv.ParseFloat(row[7], 64)
		keyContractMonthlyAveragePosition, _ := strconv.ParseFloat(row[8], 64)
		theHandlingFeeDueTodayThisMonth, _ := strconv.ParseFloat(row[10], 64)
		corporatePositionMultiplier, _ := strconv.ParseFloat(row[9], 64)
		tradingDate, _ := time.Parse("2006-01-02", row[0])
		menu := model.Us001CzceBrokerPositionDaily{
			Time:                              tradingDate,
			BrokerId:                          row[1],
			ProductCode:                       strings.ToUpper(row[2]),
			CustomerDailyPosition:             float64(customerDailyPosition),
			InstitutionalDailyPosition:        float64(institutionalDailyPosition),
			SpecialLawCustomerDailyPosition:   float64(specialLawCustomerDailyPosition),
			CorporateCustomerDailyPosition:    float64(corporateCustomerDailyPosition),
			DailyAveragePosition:              float64(dailyAveragePosition),
			KeyContractMonthlyAveragePosition: float64(keyContractMonthlyAveragePosition),
			TheHandlingFeeDueTodayThisMonth:   float64(theHandlingFeeDueTodayThisMonth),
			CorporatePositionMultiplier:       float64(corporatePositionMultiplier),
			ExchangeId:                        "ZCE",
		}
		menus = append(menus, menu)
	}
	err = CreateUs001CzceBrokerPositionDailyList(menus)
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
