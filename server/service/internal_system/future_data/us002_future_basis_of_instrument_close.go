package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
	"math"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs002FutureBasisOfInstrumentClose
//@description: 创建Us002FutureBasisOfInstrumentClose记录
//@param: us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
//@return: err error

func CreateUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose) (err error) {
	err = global.GVA_DB.Create(&us002FutureBasisOfInstrumentClose).Error
	sql := "REFRESH MATERIALIZED VIEW future.mv_future_position_detail WITH DATA;"
	go global.GVA_DB.Exec(sql)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs002FutureBasisOfInstrumentClose
//@description: 删除Us002FutureBasisOfInstrumentClose记录
//@param: us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
//@return: err error

func DeleteUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose) (err error) {
	err = global.GVA_DB.Delete(&us002FutureBasisOfInstrumentClose).Error
	sql := "REFRESH MATERIALIZED VIEW future.mv_future_position_detail WITH DATA;"
	go global.GVA_DB.Exec(sql)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs002FutureBasisOfInstrumentCloseByIds
//@description: 批量删除Us002FutureBasisOfInstrumentClose记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs002FutureBasisOfInstrumentCloseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Us002FutureBasisOfInstrumentClose{}, "id in ?", ids.Ids).Error
	sql := "REFRESH MATERIALIZED VIEW future.mv_future_position_detail WITH DATA;"
	go global.GVA_DB.Exec(sql)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs002FutureBasisOfInstrumentClose
//@description: 更新Us002FutureBasisOfInstrumentClose记录
//@param: us002FutureBasisOfInstrumentClose *model.Us002FutureBasisOfInstrumentClose
//@return: err error

func UpdateUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose) (err error) {
	err = global.GVA_DB.Save(&us002FutureBasisOfInstrumentClose).Error
	sql := "REFRESH MATERIALIZED VIEW future.mv_future_position_detail WITH DATA;"
	go global.GVA_DB.Exec(sql)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs002FutureBasisOfInstrumentClose
//@description: 根据id获取Us002FutureBasisOfInstrumentClose记录
//@param: id uint
//@return: err error, us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose

func GetUs002FutureBasisOfInstrumentClose(id uint) (err error, us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose) {
	err = global.GVA_DB.Where("id = ?", id).First(&us002FutureBasisOfInstrumentClose).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs002FutureBasisOfInstrumentCloseInfoList
//@description: 分页获取Us002FutureBasisOfInstrumentClose记录
//@param: info request.Us002FutureBasisOfInstrumentCloseSearch
//@return: err error, list interface{}, total int64

func GetUs002FutureBasisOfInstrumentCloseInfoList(info request.Us002FutureBasisOfInstrumentCloseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Us002FutureBasisOfInstrumentClose{})
	var us002FutureBasisOfInstrumentCloses []model.Us002FutureBasisOfInstrumentClose
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductCode != "" {
		db = db.Where("product_code = ?", info.ProductCode)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us002FutureBasisOfInstrumentCloses).Error
	return err, us002FutureBasisOfInstrumentCloses, total
}

//获取近月交割合约信息
func PositionDeliveryMonthInstrumentlIST(info request.PositionDeliveryMonthInstrument) (err error, list interface{}, total int64) {

	//获取当前用户所属的部门list
	var deparmentList []string
	sql := "SELECT department_code FROM master_data.us010_user_department where user_id=?"
	err = global.GVA_DB.Raw(sql, info.UserID).Scan(&deparmentList).Error
	//拼接sql适用的字符串
	var dept string
	dept = "(" + dept
	for i := 0; i < len(deparmentList); i++ {
		dept = dept + "'" + deparmentList[i] + "'"
		if i+1 < len(deparmentList) {
			dept = dept + ","
		}
	}
	dept = dept + ")"

	var timeStr string
	var positionDeliveryMonthInstrument []model.PositionDeliveryMonthInstrument
	if info.Time != "" {
		timeStr = info.Time
	} else {
		current_time := time.Now().AddDate(0, 0, 0)
		timeStr = current_time.Format("2006-01-02 15:04:05")
	}
	//查询当前账号归属部门的近月合约list
	asql := "SELECT * from master_data.fn_get_delivery_month_instrument_detail_by_dept(to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date,to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date) where department_code  in " + dept + ""
	err = global.GVA_DB.Raw(asql).Scan(&positionDeliveryMonthInstrument).Error
	total = int64(len(positionDeliveryMonthInstrument))
	//获取近月合约品种的远月合约list
	var instrument []model.FarMonthInstrument
	nsql := "SELECT instrument,product from master_data.fn_get_no_risk_far_month_instrument(to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date)"
	err = global.GVA_DB.Raw(nsql).Scan(&instrument).Error
	//将近月合约信息和远月合约list组合起来
	for i := 0; i < len(positionDeliveryMonthInstrument); i++ {
		for j := 0; j < len(instrument); j++ {
			if positionDeliveryMonthInstrument[i].Product == instrument[j].Product {
				positionDeliveryMonthInstrument[i].FarMonthInstrument = append(positionDeliveryMonthInstrument[i].FarMonthInstrument, instrument[j])
			}
		}
	}
	return err, positionDeliveryMonthInstrument, total
	// sql := "with source as (SELECT '2022-06-13',instrument, exchange_id,  product,launch_date, last_trade_date, month, last_ddate FROM future.future_instruments " +
	// 	"where last_trade_date>='2022-06-13' and to_char(last_trade_date,'yyyy-mm')='2022-06' and exchange_id!='CFX' and product!='SCTAS' order by instrument) " +
	// 	"select s.*, case when nt.month is null then 1 else 0 end as is_throw from source s left join future.product_not_throw nt on s.month =nt.month and s.product=nt.product " +
	// 	"WHERE instrument  IN (SELECT instrument from future.fuc_future_position_detail('2022-06-13', '2022-06-13'))"
}

//该方法暂不使用
func GetBenchmarkInstrumentList(info request.PositionDeliveryMonthInstrument) (err error, list interface{}, total int64) {
	var instrument []model.FarMonthInstrument
	current_time := time.Now().AddDate(0, 0, -2)
	timeStr := current_time.Format("2006-01-02 15:04:05")
	sql := "with source as (SELECT to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date as time,instrument, exchange_id,  product,launch_date, last_trade_date, month, last_ddate FROM future.future_instruments " +
		"where last_trade_date>=to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date and to_char(last_trade_date,'yyyy-mm')=to_char('" + timeStr + "'::date,'yyyy-mm') and exchange_id!='CFX' and product!='SCTAS' order by instrument), " +
		"near_in as (select s.*, case when nt.month is null then 1 else 0 end as is_throw from source s left join future.product_not_throw nt on s.month =nt.month and s.product=nt.product " +
		"WHERE instrument  IN (SELECT instrument from future.fuc_future_position_detail(to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date, to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date))) " +
		"SELECT instrument,product FROM future.future_instruments where product in (select product from near_in ) and last_trade_date>to_char('" + timeStr + "'::date,'yyyy-mm-dd')::date and instrument not in (select instrument from near_in )order by instrument"
	err = global.GVA_DB.Raw(sql).Scan(&instrument).Error
	total = int64(len(instrument))
	return err, instrument, total
}

//计算无风险价差返回给前段
func GetNoRiskValue(info request.NoRiskValue) (err error, list interface{}) {
	var noRiskValue []request.NoRiskValue
	sql := "with source as(select time, openinterest_x, openinterest_y, (-1)*safe_spread as safe_spread, current_spread, spread_rate,ticker_f,ticker_n " +
		"from future.safe_spread(to_char('" + info.Time + "'::date,'yyyy-mm-dd')::date,to_char('" + info.Time + "'::date,'yyyy-mm-dd')::date) c " +
		"where upper(c.ticker_n) ='" + info.Instrument + "' and upper(c.ticker_f) ='" + info.FarMonthInstrument + "' order by c.time, c.product) " +
		"select time,safe_spread,ticker_f,ticker_n from source"
	err = global.GVA_DB.Raw(sql).Scan(&noRiskValue).Error
	for i := 0; i < len(noRiskValue); i++ {
		noRiskValue[i].NoRiskValue = (math.Floor(noRiskValue[i].NoRiskValue*10 + 0.5)) / 10
	}
	return err, noRiskValue
}

func CreateBasisOfInstrumentCloseList(BasisOInstrumentList request.BasisOInstrumentList) (err error) {
	err = global.GVA_DB.Model(&model.Us002FutureBasisOfInstrumentClose{}).CreateInBatches(BasisOInstrumentList.Us002FutureBasisOfInstrumentClose, len(BasisOInstrumentList.Us002FutureBasisOfInstrumentClose)).Error
	sql := "REFRESH MATERIALIZED VIEW future.mv_future_position_detail WITH DATA;"
	go global.GVA_DB.Exec(sql)
	return err
}
