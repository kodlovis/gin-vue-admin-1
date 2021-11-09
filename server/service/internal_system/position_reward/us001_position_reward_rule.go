package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePositionRewardRule
//@description: 创建PositionRewardRule记录
//@param: positionRewardRule model.PositionRewardRule
//@return: err error

func CreatePositionRewardRule(positionRewardRule mp.PositionRewardRule) (err error) {
	err = global.GVA_DB.Create(&positionRewardRule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePositionRewardRule
//@description: 删除PositionRewardRule记录
//@param: positionRewardRule model.PositionRewardRule
//@return: err error

func DeletePositionRewardRule(positionRewardRule mp.PositionRewardRule) (err error) {
	err = global.GVA_DB.Delete(&positionRewardRule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePositionRewardRuleByIds
//@description: 批量删除PositionRewardRule记录
//@param: ids request.IdsReq
//@return: err error

func DeletePositionRewardRuleByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.PositionRewardRule{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePositionRewardRule
//@description: 更新PositionRewardRule记录
//@param: positionRewardRule *model.PositionRewardRule
//@return: err error

func UpdatePositionRewardRule(positionRewardRule mp.PositionRewardRule) (err error) {
	err = global.GVA_DB.Save(&positionRewardRule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPositionRewardRule
//@description: 根据id获取PositionRewardRule记录
//@param: id uint
//@return: err error, positionRewardRule model.PositionRewardRule

func GetPositionRewardRule(id uint) (err error, positionRewardRule mp.PositionRewardRule) {
	err = global.GVA_DB.Where("id = ?", id).First(&positionRewardRule).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPositionRewardRuleInfoList
//@description: 分页获取PositionRewardRule记录
//@param: info request.PositionRewardRuleSearch
//@return: err error, list interface{}, total int64

func GetPositionRewardRuleInfoList(info rp.PositionRewardRuleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.PositionRewardRule{})
	var positionRewardRules []mp.PositionRewardRule
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.EffectiveDate != "" {
		db = db.Where("effective_date = ?", info.EffectiveDate)
	}
	if info.ExpirationDate != "" {
		db = db.Where("expiration_date = ?", info.ExpirationDate)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&positionRewardRules).Error
	return err, positionRewardRules, total
}
