package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRole
//@description: 创建Role记录
//@param: Role model.Role
//@return: err error

func CreateRole(Role mp.Role) (err error) {
	err = global.GVA_DB.Create(&Role).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRole
//@description: 删除Role记录
//@param: Role model.Role
//@return: err error

func DeleteRole(Role mp.Role) (err error) {
	err = global.GVA_DB.Delete(Role).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRoleByIds
//@description: 批量删除Role记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRoleByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Role{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRole
//@description: 更新Role记录
//@param: Role *model.Role
//@return: err error

func UpdateRole(Role *mp.Role) (err error) {
	err = global.GVA_DB.Save(Role).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRole
//@description: 根据id获取Role记录
//@param: id uint
//@return: err error, Role model.Role

func GetRole(id uint) (err error, Role mp.Role) {
	err = global.GVA_DB.Where("id = ?", id).First(&Role).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRoleInfoList
//@description: 分页获取Role记录
//@param: info request.RoleSearch
//@return: err error, list interface{}, total int64

func GetRoleInfoList(info rp.RoleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Role{})
    var Roles []mp.Role
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Roles).Error
	return err, Roles, total
}