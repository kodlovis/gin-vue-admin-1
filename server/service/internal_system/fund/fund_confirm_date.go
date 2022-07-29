package fund

import (
	"fmt"
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFundConfirm
//@description: 创建FundConfirm记录
//@param: fundConfirmDate model.FundConfirm
//@return: err error

func CreateFundConfirm(fundConfirmDate model.FundConfirm) (err error) {
	err = global.GVA_DB.Create(&fundConfirmDate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFundConfirm
//@description: 删除FundConfirm记录
//@param: fundConfirmDate model.FundConfirm
//@return: err error

func DeleteFundConfirm(fundConfirmDate model.FundConfirm) (err error) {
	err = global.GVA_DB.Delete(&fundConfirmDate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFundConfirmByIds
//@description: 批量删除FundConfirm记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFundConfirmByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.FundConfirm{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFundConfirm
//@description: 更新FundConfirm记录
//@param: fundConfirmDate *model.FundConfirm
//@return: err error

func UpdateFundConfirm(fundConfirmDate model.FundConfirm) (err error) {
	err = global.GVA_DB.Save(&fundConfirmDate).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFundConfirm
//@description: 根据id获取FundConfirm记录
//@param: id uint
//@return: err error, fundConfirmDate model.FundConfirm

func GetFundConfirm(id uint) (err error, fundConfirmDate model.FundConfirm) {
	err = global.GVA_DB.Where("id = ?", id).First(&fundConfirmDate).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFundConfirmInfoList
//@description: 分页获取FundConfirm记录
//@param: info request.FundConfirmSearch
//@return: err error, list interface{}, total int64

func GetFundConfirmInfoList(info request.FundConfirmSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.FundConfirm{})
	var fundConfirmDates []model.FundConfirm
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fundConfirmDates).Error
	return err, fundConfirmDates, total
}
func GetLast() (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.FundConfirm{})
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02")
	var fundConfirmDates []model.FundConfirm
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("to_char( start_date,'yyyy-mm-dd')::date = ?", addTime).Last(&fundConfirmDates).Error
	return err, fundConfirmDates
}

func GetCycle(addTime string) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.FundConfirm{})
	var fundConfirmDates []model.FundConfirm
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("to_char( start_date,'yyyy-mm-dd')::date = ?", addTime).Last(&fundConfirmDates).Error
	if err != nil {
		fundConfirmDates = nil
		err = nil
	}
	return err, fundConfirmDates
}

func GetSetting(addTime string) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.ConfirmSetting{})
	var fundConfirmDates []model.ConfirmSetting
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("to_char( confirm_date,'yyyy-mm-dd')::date = ?", addTime).Find(&fundConfirmDates).Error
	return err, fundConfirmDates
}

func UpSetting(confirm_data model.ConfirmSetting) (err error) {
	// 创建db
	db := global.GVA_DB.Model(&model.ConfirmSetting{})
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("to_char( confirm_date,'yyyy-mm-dd')::date = ?", confirm_data.Confirm_date).Update("is_ok", confirm_data.Is_ok).Error
	fmt.Println(err)
	return err
}
