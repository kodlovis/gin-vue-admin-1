package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMac
//@description: 创建Mac记录
//@param: Mac model.Mac
//@return: err error

func CreateMac(Mac mp.Mac) (err error) {
	err = global.GVA_DB.Create(&Mac).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMac
//@description: 删除Mac记录
//@param: Mac model.Mac
//@return: err error

func DeleteMac(Mac mp.Mac) (err error) {
	err = global.GVA_DB.Delete(Mac).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMacByIds
//@description: 批量删除Mac记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMacByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Mac{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMac
//@description: 更新Mac记录
//@param: Mac *model.Mac
//@return: err error

func UpdateMac(list []mp.Mac) (err error) {
	for i := 0; i < len(list); i++ {
		err = global.GVA_DB.Save(list[i]).Error
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMac
//@description: 根据id获取Mac记录
//@param: id uint
//@return: err error, Mac model.Mac

func GetMac(id uint) (err error, Mac mp.Mac) {
	err = global.GVA_DB.Where("id = ?", id).First(&Mac).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMacInfoList
//@description: 分页获取Mac记录
//@param: info request.MacSearch
//@return: err error, list interface{}, total int64

func GetMacInfoList(info rp.MacSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Mac{})
    var Macs []mp.Mac
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Macs).Error
	return err, Macs, total
}