package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRiskSetting
//@description: 创建RiskSetting记录
//@param: RiskSetting model.RiskSetting
//@return: err error

func CreateRiskSetting(RiskSetting mp.RiskSetting) (err error) {
	err = global.GVA_DB.Create(&RiskSetting).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRiskSetting
//@description: 删除RiskSetting记录
//@param: RiskSetting model.RiskSetting
//@return: err error

func DeleteRiskSetting(RiskSetting mp.RiskSetting) (err error) {
	err = global.GVA_DB.Delete(RiskSetting).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRiskSettingByIds
//@description: 批量删除RiskSetting记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRiskSettingByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.RiskSetting{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRiskSetting
//@description: 更新RiskSetting记录
//@param: RiskSetting *model.RiskSetting
//@return: err error

func UpdateRiskSetting(RiskSetting *mp.RiskSetting) (err error) {
	err = global.GVA_DB.Save(RiskSetting).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRiskSetting
//@description: 根据id获取RiskSetting记录
//@param: id uint
//@return: err error, RiskSetting model.RiskSetting

func GetRiskSetting(id uint) (err error, RiskSetting mp.RiskSetting) {
	err = global.GVA_DB.Where("id = ?", id).First(&RiskSetting).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRiskSettingInfoList
//@description: 分页获取RiskSetting记录
//@param: info request.RiskSettingSearch
//@return: err error, list interface{}, total int64

func GetRiskSettingInfoList(info rp.RiskSettingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.RiskSetting{})
    var RiskSettings []mp.RiskSetting
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&RiskSettings).Error
	return err, RiskSettings, total
}