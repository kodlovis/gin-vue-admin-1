package future_inventory

import (
	"errors"
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
//@function: CreateUs004FutureStockLme
//@description: 创建Us004FutureStockLme记录
//@param: us004FutureStockLme model.Us004FutureStockLme
//@return: err error

func CreateUs004FutureStockLme(us004FutureStockLme mif.Us004FutureStockLme) (err error) {
	err = global.GVA_DB.Create(&us004FutureStockLme).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureStockLme
//@description: 删除Us004FutureStockLme记录
//@param: us004FutureStockLme model.Us004FutureStockLme
//@return: err error

func DeleteUs004FutureStockLme(us004FutureStockLme mif.Us004FutureStockLme) (err error) {
	err = global.GVA_DB.Delete(&us004FutureStockLme).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureStockLmeByIds
//@description: 批量删除Us004FutureStockLme记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs004FutureStockLmeByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.Us004FutureStockLme{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs004FutureStockLme
//@description: 更新Us004FutureStockLme记录
//@param: us004FutureStockLme *model.Us004FutureStockLme
//@return: err error

func UpdateUs004FutureStockLme(us004FutureStockLme mif.Us004FutureStockLme) (err error) {
	err = global.GVA_DB.Save(&us004FutureStockLme).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureStockLme
//@description: 根据id获取Us004FutureStockLme记录
//@param: id uint
//@return: err error, us004FutureStockLme model.Us004FutureStockLme

func GetUs004FutureStockLme(id uint) (err error, us004FutureStockLme mif.Us004FutureStockLme) {
	err = global.GVA_DB.Where("id = ?", id).First(&us004FutureStockLme).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureStockLmeInfoList
//@description: 分页获取Us004FutureStockLme记录
//@param: info request.Us004FutureStockLmeSearch
//@return: err error, list interface{}, total int64

func GetUs004FutureStockLmeInfoList(info rif.Us004FutureStockLmeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.Us004FutureStockLme{})
	var us004FutureStockLmes []mif.Us004FutureStockLme
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.MarketDay.IsZero() {
		db = db.Where("`market_day` = ?", info.MarketDay)
	}
	if info.FuturesVariety != "" {
		db = db.Where("`futures_variety` = ?", info.FuturesVariety)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us004FutureStockLmes).Error
	return err, us004FutureStockLmes, total
}
func ParseLMEInventoryExcel2InfoList() error {
	skipHeader := true
	fixedHeader := []string{"product", "Gen Val7", "Close 1", "Date"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "InventoryExcel.xlsx")
	if err != nil {
		return err
	}
	menus := make([]mif.Us004FutureStockLme, 0)

	err = file.RemoveCol("Sheet1", "B")
	err = file.RemoveRow("Sheet1", 1)
	err = file.RemoveRow("Sheet1", 1)
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
		volume, _ := strconv.ParseFloat(row[2], 0)
		cancelVolume, _ := strconv.ParseFloat(row[1], 0)
		time, _ := time.Parse("2006-01-02", row[3])

		menu := mif.Us004FutureStockLme{
			MarketDay:        time,
			FuturesVariety:   row[0],
			WarehouseReceipt: volume,
			CancelStock:      cancelVolume,
		}
		menus = append(menus, menu)
	}
	err = CreateLEMInventoryList(menus)
	if err != nil {
		return err
	}
	err = os.Remove(global.GVA_CONFIG.Excel.Dir + "InventoryExcel.xlsx")
	if err != nil {
		return err
	}
	return err
}

//批量创建
func CreateLEMInventoryList(list []mif.Us004FutureStockLme) (err error) {
	err = global.GVA_DB.Model(&mif.Us004FutureStockLme{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(list, len(list)).Error
	return
}
