package scale_rights

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/scale_rights"
	rp "gin-vue-admin/model/request/internal_system/scale_rights"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFuturePositionSizeDetail
//@description: 创建FuturePositionSizeDetail记录
//@param: futurePositionSizeDetail model.FuturePositionSizeDetail
//@return: err error

func CreateFuturePositionSizeDetail(futurePositionSizeDetail mp.FuturePositionSizeDetail) (err error) {
	err = global.GVA_DB.Create(&futurePositionSizeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFuturePositionSizeDetail
//@description: 删除FuturePositionSizeDetail记录
//@param: futurePositionSizeDetail model.FuturePositionSizeDetail
//@return: err error

func DeleteFuturePositionSizeDetail(futurePositionSizeDetail mp.FuturePositionSizeDetail) (err error) {
	err = global.GVA_DB.Delete(&futurePositionSizeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFuturePositionSizeDetailByIds
//@description: 批量删除FuturePositionSizeDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFuturePositionSizeDetailByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.FuturePositionSizeDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFuturePositionSizeDetail
//@description: 更新FuturePositionSizeDetail记录
//@param: futurePositionSizeDetail *model.FuturePositionSizeDetail
//@return: err error

func UpdateFuturePositionSizeDetail(futurePositionSizeDetail mp.FuturePositionSizeDetail) (err error) {
	err = global.GVA_DB.Save(&futurePositionSizeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFuturePositionSizeDetail
//@description: 根据id获取FuturePositionSizeDetail记录
//@param: id uint
//@return: err error, futurePositionSizeDetail model.FuturePositionSizeDetail

func GetFuturePositionSizeDetail(id uint) (err error, futurePositionSizeDetail mp.FuturePositionSizeDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&futurePositionSizeDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFuturePositionSizeDetailInfoList
//@description: 分页获取FuturePositionSizeDetail记录
//@param: info request.FuturePositionSizeDetailSearch
//@return: err error, list interface{}, total int64

func GetFuturePositionSizeDetailInfoList(info rp.FuturePositionSizeDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.FuturePositionSizeDetail{})
	var futurePositionSizeDetails []mp.FuturePositionSizeDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	if info.TradeType != "" {
		db = db.Where("trade_type LIKE ?", "%"+info.TradeType+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futurePositionSizeDetails).Error
	return err, futurePositionSizeDetails, total
}
