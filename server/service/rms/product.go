package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateProduct
//@description: 创建Product记录
//@param: Product model.Product
//@return: err error

func CreateProduct(Product mp.Product) (err error) {
	err = global.GVA_DB.Create(&Product).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProduct
//@description: 删除Product记录
//@param: Product model.Product
//@return: err error

func DeleteProduct(Product mp.Product) (err error) {
	err = global.GVA_DB.Delete(Product).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProductByIds
//@description: 批量删除Product记录
//@param: ids request.IdsReq
//@return: err error

func DeleteProductByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Product{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateProduct
//@description: 更新Product记录
//@param: Product *model.Product
//@return: err error

func UpdateProduct(Product *mp.Product) (err error) {
	err = global.GVA_DB.Save(Product).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProduct
//@description: 根据id获取Product记录
//@param: id uint
//@return: err error, Product model.Product

func GetProduct(id uint) (err error, Product mp.Product) {
	err = global.GVA_DB.Where("id = ?", id).First(&Product).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProductInfoList
//@description: 分页获取Product记录
//@param: info request.ProductSearch
//@return: err error, list interface{}, total int64

func GetProductInfoList(info rp.ProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Product{})
    var Products []mp.Product
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Products).Error
	return err, Products, total
}