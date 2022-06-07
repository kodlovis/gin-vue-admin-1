package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDepartment
//@description: 创建Department记录
//@param: department model.Department
//@return: err error

func CreateDepartment(department model.Department) (err error) {
	err = global.GVA_DB.Create(&department).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDepartment
//@description: 删除Department记录
//@param: department model.Department
//@return: err error

func DeleteDepartment(department model.Department) (err error) {
	err = global.GVA_DB.Delete(&department).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDepartmentByIds
//@description: 批量删除Department记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDepartmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Department{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDepartment
//@description: 更新Department记录
//@param: department *model.Department
//@return: err error

func UpdateDepartment(department model.Department) (err error) {
	err = global.GVA_DB.Save(&department).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepartment
//@description: 根据id获取Department记录
//@param: id uint
//@return: err error, department model.Department

func GetDepartment(id uint) (err error, department model.Department) {
	err = global.GVA_DB.Where("id = ?", id).First(&department).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepartmentInfoList
//@description: 分页获取Department记录
//@param: info request.DepartmentSearch
//@return: err error, list interface{}, total int64

func GetDepartmentInfoList(info request.DepartmentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Department{})
	var departments []model.Department
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.DepartmentName != "" {
		db = db.Where("`department_name` = ?", info.DepartmentName)
	}
	if info.DepartmentCode != "" {
		db = db.Where("`department_code` = ?", info.DepartmentCode)
	}
	if info.CompanyCode != "" {
		db = db.Where("`company_code` = ?", info.CompanyCode)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&departments).Error
	return err, departments, total
}
