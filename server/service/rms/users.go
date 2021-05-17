package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUsers
//@description: 创建Users记录
//@param: Users model.Users
//@return: err error

func CreateUsers(Users mp.Users) (err error) {
	err = global.GVA_DB.Create(&Users).Error
	return err
}

func GetLastUser() (err error, User mp.Users) {
	err = global.GVA_DB.Last(&User).Error
	return
}

func CreateUserMac(list []mp.UserMac) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Create(&list[i]).Error
	}
	return err
}
func CreateUserProduct(list []mp.UserProduct) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Create(&list[i]).Error
	}
	return err
}
func RemoveUserMacProduct(Users mp.Users) (err error) {
	err = global.GVA_DB.Table("users_mac").Where("users_id = ?", Users.ID).Delete(&[]mp.UserMac{}).Error
	err = global.GVA_DB.Table("users_product").Where("users_id = ?", Users.ID).Delete(&[]mp.UserProduct{}).Error
	return err
}
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsers
//@description: 删除Users记录
//@param: Users model.Users
//@return: err error

func DeleteUsers(Users mp.Users) (err error) {
	err = global.GVA_DB.Delete(Users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsersByIds
//@description: 批量删除Users记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUsersByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Users{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUsers
//@description: 更新Users记录
//@param: Users *model.Users
//@return: err error

func UpdateUsers(Users *mp.Users) (err error) {
	err = global.GVA_DB.Save(Users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsers
//@description: 根据id获取Users记录
//@param: id uint
//@return: err error, Users model.Users

func GetUsers(id uint) (err error, Users mp.Users) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Macs").Preload("Products").Preload("Role").First(&Users).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsersInfoList
//@description: 分页获取Users记录
//@param: info request.UsersSearch
//@return: err error, list interface{}, total int64

func GetUsersInfoList(info rp.UsersSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Users{})
    var Users []mp.Users
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Users).Error
	err = db.Preload("Macs").Preload("Products").Preload("Role").Find(&Users).Error
	return err, Users, total
}

func GetUIList(userid string, password string, mac string) (err error, list interface{}, total int64, code int64) {
	var info rp.UsersSearch
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Users{}).Where("name = ? and password = ?", userid,password)
    var Users []mp.Users
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Users).Error
	err = db.Preload("Macs", "mac_address = ?", mac).Preload("Products").Preload("Role").Find(&Users).Error
	var user mp.Users
	err = db.Where("name = ?", userid).First(&user).Error
	if user.Password != password{
		code=1
	}
	return err, Users, total,code
}
func GetUIListByServer() (err error, list interface{}, total int64) {
	var info rp.UsersSearch
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Users{})
    var Users []mp.Users
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Users).Error
	err = db.Preload("Macs").Preload("Products").Preload("Role").Find(&Users).Error
	return err, Users, total
}