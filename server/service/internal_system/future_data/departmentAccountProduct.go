package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDepartmentAccountProduct
//@description: 创建DepartmentAccountProduct记录
//@param: departmentAccountProduct model.DepartmentAccountProduct
//@return: err error

func CreateDepartmentAccountProduct(departmentAccountProduct model.DepartmentAccountProduct) (err error) {
	err = global.GVA_DB.Create(&departmentAccountProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDepartmentAccountProduct
//@description: 删除DepartmentAccountProduct记录
//@param: departmentAccountProduct model.DepartmentAccountProduct
//@return: err error

func DeleteDepartmentAccountProduct(departmentAccountProduct model.DepartmentAccountProduct) (err error) {
	err = global.GVA_DB.Delete(&departmentAccountProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDepartmentAccountProductByIds
//@description: 批量删除DepartmentAccountProduct记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDepartmentAccountProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DepartmentAccountProduct{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDepartmentAccountProduct
//@description: 更新DepartmentAccountProduct记录
//@param: departmentAccountProduct *model.DepartmentAccountProduct
//@return: err error

func UpdateDepartmentAccountProduct(departmentAccountProduct model.DepartmentAccountProduct) (err error) {
	err = global.GVA_DB.Save(&departmentAccountProduct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepartmentAccountProduct
//@description: 根据id获取DepartmentAccountProduct记录
//@param: id uint
//@return: err error, departmentAccountProduct model.DepartmentAccountProduct

func GetDepartmentAccountProduct(id uint) (err error, departmentAccountProduct model.DepartmentAccountProduct) {
	err = global.GVA_DB.Where("id = ?", id).First(&departmentAccountProduct).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepartmentAccountProductInfoList
//@description: 分页获取DepartmentAccountProduct记录
//@param: info request.DepartmentAccountProductSearch
//@return: err error, list interface{}, total int64

func GetDepartmentAccountProductInfoList(info request.DepartmentAccountProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.DepartmentAccountProduct{})
	var departmentAccountProducts []model.DepartmentAccountProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.DepartmentCode != "" {
		db = db.Where("`department_code` = ?", info.DepartmentCode)
	}
	if info.AccountId != "" {
		db = db.Where("`account_id` = ?", info.AccountId)
	}
	if !info.EffectiveDate.IsZero() {
		db = db.Where("`effective_date` = ?", info.EffectiveDate)
	}
	if !info.ExpirationDate.IsZero() {
		db = db.Where("`expiration_date` = ?", info.ExpirationDate)
	}
	if info.ProductCode != "" {
		db = db.Where("`product_code` = ?", info.ProductCode)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&departmentAccountProducts).Error
	return err, departmentAccountProducts, total
}
