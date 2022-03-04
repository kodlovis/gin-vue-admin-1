package future_inventory

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_inventory"
	rif "gin-vue-admin/model/request/internal_system/future_inventory"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gorm.io/gorm/clause"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs004FutureInventoryDaily
//@description: 创建Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily model.Us004FutureInventoryDaily
//@return: err error

func CreateUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Create(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureInventoryDaily
//@description: 删除Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily model.Us004FutureInventoryDaily
//@return: err error

func DeleteUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Delete(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureInventoryDailyByIds
//@description: 批量删除Us004FutureInventoryDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs004FutureInventoryDailyByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.Us004FutureInventoryDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs004FutureInventoryDaily
//@description: 更新Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily *model.Us004FutureInventoryDaily
//@return: err error

func UpdateUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Save(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureInventoryDaily
//@description: 根据id获取Us004FutureInventoryDaily记录
//@param: id uint
//@return: err error, us004FutureInventoryDaily model.Us004FutureInventoryDaily

func GetUs004FutureInventoryDaily(id uint) (err error, us004FutureInventoryDaily mif.Us004FutureInventoryDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&us004FutureInventoryDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureInventoryDailyInfoList
//@description: 分页获取Us004FutureInventoryDaily记录
//@param: info request.Us004FutureInventoryDailySearch
//@return: err error, list interface{}, total int64

func GetUs004FutureInventoryDailyInfoList(info rif.Us004FutureInventoryDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.Us004FutureInventoryDaily{}).Order("time desc")
	var us004FutureInventoryDailys []mif.Us004FutureInventoryDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	if info.ExchangeId != "" {
		db = db.Where("exchange_id LIKE ?", "%"+info.ExchangeId+"%")
	}
	if info.Comment != "" {
		db = db.Where("comment LIKE ?", "%"+info.Comment+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us004FutureInventoryDailys).Error
	return err, us004FutureInventoryDailys, total
}
func GetUs004FutureInventoryDailyType(info rif.Us004FutureInventoryDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.Us004FutureInventoryDaily{})
	var us004FutureInventoryDailys []mif.Us004FutureInventoryDaily
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Select("comment").Distinct("comment").Find(&us004FutureInventoryDailys).Error
	return err, us004FutureInventoryDailys, total
}

func ParseInventoryExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"日期", "数量", "品种", "库存类型"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "InventoryExcel.xlsx")
	if err != nil {
		return err
	}
	menus := make([]mif.Us004FutureInventoryDaily, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return err
	}
	db := global.GVA_DB.Model(&mif.Us004ProductAssociationDetail{})
	var us004ProductAssociationDetails []mif.Us004ProductAssociationDetail
	err = db.Find(&us004ProductAssociationDetails).Error
	var productCode string
	var unit string
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
				err = ParseLMEInventoryExcel2InfoList()
				if err != nil {
					return err
				}
				return nil
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		// amount, _ := strconv.ParseInt(row[6], 0, 0)
		// direction, _ := strconv.ParseInt(row[5], 0, 0)
		// hedgeFlag, _ := strconv.ParseInt(row[4], 0, 0)
		volume, _ := strconv.ParseFloat(row[1], 0)
		time, _ := time.Parse("2006-01-02", row[0])

		// if row[1] == "鲁证期货" {
		// 	row[1] = "0275"
		// }
		for i := 0; i < len(us004ProductAssociationDetails); i++ {
			if row[2] == us004ProductAssociationDetails[i].SubProductName {
				productCode = us004ProductAssociationDetails[i].ProductName
				unit = us004ProductAssociationDetails[i].Unit
			}
		}
		menu := mif.Us004FutureInventoryDaily{
			Time:        time,
			Volume:      float32(volume),
			Comment:     row[2],
			ProductCode: productCode,
			Unit:        unit,
			Type:        row[3],
		}
		menus = append(menus, menu)
	}
	err = CreateInventoryList(menus)
	err = os.Remove(global.GVA_CONFIG.Excel.Dir + "InventoryExcel.xlsx")
	if err != nil {
		return err
	}
	return err
}

//批量创建
func CreateInventoryList(list []mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Model(&mif.Us004FutureInventoryDaily{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(list, len(list)).Error
	return
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
