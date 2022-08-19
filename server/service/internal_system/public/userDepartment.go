package public

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/public"
	request "gin-vue-admin/model/request/internal_system/public"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUserDepartment
//@description: 创建UserDepartment记录
//@param: userDepartment model.UserDepartment
//@return: err error

func CreateUserDepartment(userDepartment model.UserDepartment) (err error) {
	err = global.GVA_DB.Create(&userDepartment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUserDepartment
//@description: 删除UserDepartment记录
//@param: userDepartment model.UserDepartment
//@return: err error

func DeleteUserDepartment(userDepartment model.UserDepartment) (err error) {
	err = global.GVA_DB.Delete(&userDepartment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUserDepartmentByIds
//@description: 批量删除UserDepartment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUserDepartmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserDepartment{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUserDepartment
//@description: 更新UserDepartment记录
//@param: userDepartment *model.UserDepartment
//@return: err error

func UpdateUserDepartment(userDepartment model.UserDepartment) (err error) {
	err = global.GVA_DB.Save(&userDepartment).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserDepartment
//@description: 根据id获取UserDepartment记录
//@param: id uint
//@return: err error, userDepartment model.UserDepartment

func GetUserDepartment(id uint) (err error, userDepartment model.UserDepartment) {
	err = global.GVA_DB.Where("id = ?", id).First(&userDepartment).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserDepartmentInfoList
//@description: 分页获取UserDepartment记录
//@param: info request.UserDepartmentSearch
//@return: err error, list interface{}, total int64

func GetUserDepartmentInfoList(info request.UserDepartmentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.UserDepartment{}).Preload("DepartmentInfo")
	var userDepartments []model.UserDepartment
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userDepartments).Error
	return err, userDepartments, total
}

func GetUserDepartmentListByID(info request.UserDepartmentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.UserDepartment{}).Preload("DepartmentInfo").Where("user_id = ?", info.UserId)
	var userDepartments []model.UserDepartment
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userDepartments).Error
	return err, userDepartments, total
}
