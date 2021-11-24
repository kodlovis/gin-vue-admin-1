package rms

import (
	"gin-vue-admin/global"
	rp "gin-vue-admin/model/request/rms"
	mp "gin-vue-admin/model/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateProductLeveragePrivateEquity
//@description: 创建ProductLeveragePrivateEquity记录
//@param: productLeveragePrivateEquity model.ProductLeveragePrivateEquity
//@return: err error

func CreateProductLeveragePrivateEquity(productLeveragePrivateEquity mp.ProductLeveragePrivateEquity) (err error) {
	err = global.GVA_DB.Create(&productLeveragePrivateEquity).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProductLeveragePrivateEquity
//@description: 删除ProductLeveragePrivateEquity记录
//@param: productLeveragePrivateEquity model.ProductLeveragePrivateEquity
//@return: err error

func DeleteProductLeveragePrivateEquity(productLeveragePrivateEquity mp.ProductLeveragePrivateEquity) (err error) {
	err = global.GVA_DB.Delete(&productLeveragePrivateEquity).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProductLeveragePrivateEquityByIds
//@description: 批量删除ProductLeveragePrivateEquity记录
//@param: ids request.IdsReq
//@return: err error

func DeleteProductLeveragePrivateEquityByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.ProductLeveragePrivateEquity{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateProductLeveragePrivateEquity
//@description: 更新ProductLeveragePrivateEquity记录
//@param: productLeveragePrivateEquity *model.ProductLeveragePrivateEquity
//@return: err error

func UpdateProductLeveragePrivateEquity(productLeveragePrivateEquity mp.ProductLeveragePrivateEquity) (err error) {
	err = global.GVA_DB.Save(&productLeveragePrivateEquity).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProductLeveragePrivateEquity
//@description: 根据id获取ProductLeveragePrivateEquity记录
//@param: id uint
//@return: err error, productLeveragePrivateEquity model.ProductLeveragePrivateEquity

func GetProductLeveragePrivateEquity(id uint) (err error, productLeveragePrivateEquity mp.ProductLeveragePrivateEquity) {
	err = global.GVA_DB.Where("id = ?", id).First(&productLeveragePrivateEquity).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProductLeveragePrivateEquityInfoList
//@description: 分页获取ProductLeveragePrivateEquity记录
//@param: info request.ProductLeveragePrivateEquitySearch
//@return: err error, list interface{}, total int64

func GetProductLeveragePrivateEquityInfoList(info rp.ProductLeveragePrivateEquitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.ProductLeveragePrivateEquity{}).Preload("ProductInfo").Order("product_code, leverage")
	var productLeveragePrivateEquitys []mp.ProductLeveragePrivateEquity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&productLeveragePrivateEquitys).Error
	return err, productLeveragePrivateEquitys, total
}
