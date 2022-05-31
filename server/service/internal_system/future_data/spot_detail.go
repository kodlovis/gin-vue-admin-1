package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_data"
	rif "gin-vue-admin/model/request/internal_system/future_data"
	"math"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSpotDetail
//@description: 创建SpotDetail记录
//@param: spotDetail model.SpotDetail
//@return: err error
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
func CreateSpotDetail(spotDetail rif.SpotDetail) (err error) {

	//v2, err := strconv.ParseFloat(spotDetail.ProfitByTrade, 64)
	//var spot mif.SpotDetail
	if IsNum(spotDetail.ProfitByTrade) == true {
		v2, err := strconv.ParseFloat(spotDetail.ProfitByTrade, 64)
		if err == nil {
			var spot mif.SpotDetail
			spot.ProfitByTrade = v2
			spot.ID = spotDetail.ID
			spot.Time = spotDetail.Time
			spot.ProductName = spotDetail.ProductName
			spot.AccountId = spotDetail.AccountId
			spot.ProfitByFloat = spotDetail.ProfitByFloat
			spot.TradeFee = spotDetail.TradeFee
			spot.DepartmentName = spotDetail.DepartmentName
			err = global.GVA_DB.Create(&spot).Error
			return err
		}
	} else if IsNum(spotDetail.ProfitByTrade) != true {
		var spot mif.SpotDetailWithoutTrade
		spot.ID = spotDetail.ID
		spot.Time = spotDetail.Time
		spot.ProductName = spotDetail.ProductName
		spot.AccountId = spotDetail.AccountId
		spot.ProfitByFloat = spotDetail.ProfitByFloat
		spot.TradeFee = spotDetail.TradeFee
		spot.DepartmentName = spotDetail.DepartmentName
		err = global.GVA_DB.Create(&spot).Error
		return err
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSpotDetail
//@description: 删除SpotDetail记录
//@param: spotDetail model.SpotDetail
//@return: err error

func DeleteSpotDetail(spotDetail mif.SpotDetail) (err error) {
	err = global.GVA_DB.Delete(&spotDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSpotDetailByIds
//@description: 批量删除SpotDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSpotDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.SpotDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSpotDetail
//@description: 更新SpotDetail记录
//@param: spotDetail *model.SpotDetail
//@return: err error

func UpdateSpotDetail(spotDetail rif.SpotDetail) (err error) {
	if IsNum(spotDetail.ProfitByTrade) == true {
		v2, err := strconv.ParseFloat(spotDetail.ProfitByTrade, 64)
		if err == nil {
			var spot mif.SpotDetail
			spot.ProfitByTrade = v2
			spot.ID = spotDetail.ID
			spot.Time = spotDetail.Time
			spot.ProductName = spotDetail.ProductName
			spot.AccountId = spotDetail.AccountId
			spot.ProfitByFloat = spotDetail.ProfitByFloat
			spot.TradeFee = spotDetail.TradeFee
			spot.DepartmentName = spotDetail.DepartmentName
			err = global.GVA_DB.Save(&spot).Error
			return err
		}
	} else if IsNum(spotDetail.ProfitByTrade) != true {
		//err = global.GVA_DB.Save(&spotDetail).Error
		//err := global.GVA_DB.Model(mif.SpotDetail{}).Where("id = ?", spotDetail.ID).Updates(mif.SpotDetail{Time: spotDetail.Time, ProductName: spotDetail.ProductName, AccountId: spotDetail.AccountId,
		//	ProfitByTrade: "", ProfitByFloat: spotDetail.ProfitByFloat, TradeFee: spotDetail.TradeFee, DepartmentName: spotDetail.DepartmentName})

		err := global.GVA_DB.Model(mif.SpotDetail{}).Where("id = ?", spotDetail.ID).Select("profit_by_trade").Updates(map[string]interface{}{"profit_by_trade": nil, "time": spotDetail.Time,
			"product_name": spotDetail.ProductName, "account_id": spotDetail.AccountId, "profit_by_float": spotDetail.ProfitByFloat, "trade_fee": spotDetail.TradeFee, "department_name": spotDetail.DepartmentName})

		// err := global.GVA_DB.Raw("update future.spot_detail set time=?, ProductName:=?, AccountId=?,
		// 	ProfitByTrade is null, ProfitByFloat=?, TradeFee=?, DepartmentName=? where id =?",).Row().Scan()
		print(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSpotDetail
//@description: 根据id获取SpotDetail记录
//@param: id uint
//@return: err error, spotDetail model.SpotDetail

func GetSpotDetail(id uint) (err error, spotDetails mif.SpotDetailWithStringTrade) {
	err = global.GVA_DB.Where("id = ?", id).Order("time desc, department_name, product_name, account_id, profit_by_float, profit_by_trade, trade_fee").First(&spotDetails).Error
	spotDetails.ProfitByFloat = (math.Floor(spotDetails.ProfitByFloat*10 + 0.5)) / 10
	if IsNum(spotDetails.ProfitByTrade) == true {
		v2, err := strconv.ParseFloat(spotDetails.ProfitByTrade, 64)
		if err == nil {
			strKm := strconv.FormatFloat((math.Floor(v2*10+0.5))/10, 'f', 1, 64)
			spotDetails.ProfitByTrade = strKm
		}
	} else if IsNum(spotDetails.ProfitByTrade) != true {
		spotDetails.ProfitByTrade = ""
	}
	spotDetails.TradeFee = (math.Floor(spotDetails.TradeFee*10 + 0.5)) / 10
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSpotDetailInfoList
//@description: 分页获取SpotDetail记录
//@param: info request.SpotDetailSearch
//@return: err error, list interface{}, total int64

func GetSpotDetailInfoList(info rif.SpotDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.SpotDetail{}).Preload("AccountInfo").Order("time desc, department_name, product_name, account_id, profit_by_float, profit_by_trade, trade_fee")
	var spotDetails []mif.SpotDetailWithStringTrade
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductName != "" {
		db = db.Where("product_name like ?", "%"+info.ProductName+"%")
	}
	if info.AccountId != "" {
		db = db.Where("account_id LIKE ?", "%"+info.AccountId+"%")
	}
	if info.DepartmentName != "" {
		db = db.Where("department_name = ?", info.DepartmentName)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&spotDetails).Error
	for i := 0; i < len(spotDetails); i++ {
		spotDetails[i].ProfitByFloat = (math.Floor(spotDetails[i].ProfitByFloat*10 + 0.5)) / 10
		if IsNum(spotDetails[i].ProfitByTrade) == true {
			v2, err := strconv.ParseFloat(spotDetails[i].ProfitByTrade, 64)
			if err == nil {
				strKm := strconv.FormatFloat((math.Floor(v2*10+0.5))/10, 'f', 1, 64)
				spotDetails[i].ProfitByTrade = strKm
			}
		} else if IsNum(spotDetails[i].ProfitByTrade) != true {
			spotDetails[i].ProfitByTrade = "已继承当月累计盈亏"
		}
		spotDetails[i].TradeFee = (math.Floor(spotDetails[i].TradeFee*10 + 0.5)) / 10
	}
	return err, spotDetails, total
}
