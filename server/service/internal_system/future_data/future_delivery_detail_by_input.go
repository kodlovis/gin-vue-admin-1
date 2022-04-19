package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFutureDeliveryDetailByInput
//@description: 创建FutureDeliveryDetailByInput记录
//@param: futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
//@return: err error

func CreateFutureDeliveryDetailByInput(futureDeliveryDetailByInput model.FutureDeliveryDetailByInput) (err error) {
	err = global.GVA_DB.Create(&futureDeliveryDetailByInput).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDeliveryDetailByInput
//@description: 删除FutureDeliveryDetailByInput记录
//@param: futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
//@return: err error

func DeleteFutureDeliveryDetailByInput(futureDeliveryDetailByInput model.FutureDeliveryDetailByInput) (err error) {
	err = global.GVA_DB.Delete(&futureDeliveryDetailByInput).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDeliveryDetailByInputByIds
//@description: 批量删除FutureDeliveryDetailByInput记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFutureDeliveryDetailByInputByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.FutureDeliveryDetailByInput{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFutureDeliveryDetailByInput
//@description: 更新FutureDeliveryDetailByInput记录
//@param: futureDeliveryDetailByInput *model.FutureDeliveryDetailByInput
//@return: err error

func UpdateFutureDeliveryDetailByInput(futureDeliveryDetailByInput model.FutureDeliveryDetailByInput) (err error) {
	err = global.GVA_DB.Save(&futureDeliveryDetailByInput).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDeliveryDetailByInput
//@description: 根据id获取FutureDeliveryDetailByInput记录
//@param: id uint
//@return: err error, futureDeliveryDetailByInput model.FutureDeliveryDetailByInput

func GetFutureDeliveryDetailByInput(id uint) (err error, futureDeliveryDetailByInput model.FutureDeliveryDetailByInput) {
	err = global.GVA_DB.Where("id = ?", id).First(&futureDeliveryDetailByInput).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDeliveryDetailByInputInfoList
//@description: 分页获取FutureDeliveryDetailByInput记录
//@param: info request.FutureDeliveryDetailByInputSearch
//@return: err error, list interface{}, total int64

func GetFutureDeliveryDetailByInputInfoList(info request.FutureDeliveryDetailByInputSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.FutureDeliveryDetailByInput{}).Preload("AccountInfo").Order("time desc, department_name")
	var futureDeliveryDetailByInputs []model.FutureDeliveryDetailByInput
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.DepartmentName != "" {
		db = db.Where("department_name LIKE ?", "%"+info.DepartmentName+"%")
	}
	if info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+info.ProductName+"%")
	}
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futureDeliveryDetailByInputs).Error
	return err, futureDeliveryDetailByInputs, total
}
