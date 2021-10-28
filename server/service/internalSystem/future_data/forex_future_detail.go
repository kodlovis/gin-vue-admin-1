package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/future_data"
	rif "gin-vue-admin/model/request/internalSystem/future_data"
	"math"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateForexFutureDetail
//@description: 创建ForexFutureDetail记录
//@param: forexFutureDetail model.ForexFutureDetail
//@return: err error

func CreateForexFutureDetail(forexFutureDetail mif.ForexFutureDetail) (err error) {
	err = global.GVA_DB.Create(&forexFutureDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteForexFutureDetail
//@description: 删除ForexFutureDetail记录
//@param: forexFutureDetail model.ForexFutureDetail
//@return: err error

func DeleteForexFutureDetail(forexFutureDetail mif.ForexFutureDetail) (err error) {
	err = global.GVA_DB.Delete(&forexFutureDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteForexFutureDetailByIds
//@description: 批量删除ForexFutureDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteForexFutureDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.ForexFutureDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateForexFutureDetail
//@description: 更新ForexFutureDetail记录
//@param: forexFutureDetail *model.ForexFutureDetail
//@return: err error

func UpdateForexFutureDetail(forexFutureDetail mif.ForexFutureDetail) (err error) {
	err = global.GVA_DB.Save(&forexFutureDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetForexFutureDetail
//@description: 根据id获取ForexFutureDetail记录
//@param: id uint
//@return: err error, forexFutureDetail model.ForexFutureDetail

func GetForexFutureDetail(id uint) (err error, forexFutureDetail mif.ForexFutureDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&forexFutureDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetForexFutureDetailInfoList
//@description: 分页获取ForexFutureDetail记录
//@param: info request.ForexFutureDetailSearch
//@return: err error, list interface{}, total int64

func GetForexFutureDetailInfoList(info rif.ForexFutureDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.ForexFutureDetail{}).Order("time desc, department_name")
	var forexFutureDetails []mif.ForexFutureDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+info.ProductName+"%")
	}
	if info.AccountId != "" {
		db = db.Where("account_id LIKE ?", "%"+info.AccountId+"%")
	}
	if info.DepartmentName != "" {
		db = db.Where("department_name LIKE ?", "%"+info.DepartmentName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&forexFutureDetails).Error
	for i := 0; i < len(forexFutureDetails); i++ {
		forexFutureDetails[i].ProfitByFloat = (math.Floor(forexFutureDetails[i].ProfitByFloat*10 + 0.5)) / 10
		forexFutureDetails[i].ProfitByTrade = (math.Floor(forexFutureDetails[i].ProfitByTrade*10 + 0.5)) / 10
		forexFutureDetails[i].TradeFee = (math.Floor(forexFutureDetails[i].TradeFee*10 + 0.5)) / 10
	}
	return err, forexFutureDetails, total
}
