<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="部门">
                    <el-select v-model="setInfo.department_value" class="m-2" placeholder="Select" @change="getdata">
                        <el-option v-for="item in department_options" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="填报日期">
                    <el-date-picker v-model="setInfo.up_date" type="date" placeholder="Pick a day" :size="'default'"
                        @change="getdata" />
                </el-form-item>

                <!-- <el-form-item>
                    <el-button type="primary" icon="copyDocument" @click="getdata">查询</el-button>
                </el-form-item> -->
            </el-form>
        </div>
        <div class="gva-table-box">
            <el-table :data="tableData[0].detail" border style="width: 100%;" class="table-info bgf" width="1080"
                :header-cell-style="rowClass" cell-class-name="lm-cell" empty-text="暂无数据" tooltip-effect="dark">
                <template v-for="(item, index) in tableData">
                    <el-table-column :key="index" :label="item.title" v-if="item.title == '类型'" fixed="left"
                        width="245">
                        <template slot-scope="scope">
                            <div>{{ item.detail[scope.$index] }}</div>
                        </template>
                    </el-table-column>
                </template>

                <template v-for="(item, index) in tableData">
                    <el-table-column :key="index" :label="item.title" v-if="item.title != '类型'">
                        <template slot-scope="scope">
                            <div>{{ item.detail[scope.$index] }}</div>
                        </template>
                    </el-table-column>
                </template>
            </el-table>
        </div>
    </div>
</template>

<script>
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import {
    getCycleConfirmDate
} from "@/api/internalSystem/fund/fund_confirm_date"; //  此处请自行替换地址
import {

    getCycleData,
} from "@/api/internalSystem/fund/fund_budget_data"; //  此处请自行替换地址
export default {
    name: "BudgetReport",
    mixins: [infoList],

    data() {
        return {
            cycleList: [],
            department_options: [
                {
                    value: "有色事业部",
                    label: "有色事业部",
                },
                {
                    value: "黑色事业部",
                    label: "黑色事业部",
                },
                {
                    value: "农产事业部",
                    label: "农产事业部",
                },
                {
                    value: "原料事业部",
                    label: "原料事业部",
                },
                     {
                    value: "钢材事业部",
                    label: "钢材事业部",
                },
                {
                    value: "总经理二室",
                    label: "总经理二室",
                },
                {
                    value: "总经理一室",
                    label: "总经理一室",
                },
                {
                    value: "杉贸本部",
                    label: "杉贸本部",
                },
                {
                    value: "交易部",
                    label: "交易部",
                },
                {
                    value: "智维钢材",
                    label: "智维钢材",
                },
                {
                    value: "智维本部",
                    label: "智维本部",
                },
            ],
            setInfo: {
                department_value: "有色事业部",
                up_date: new Date().getTime(),
            },
            isChooseCols: false,
            tableData: [
                {
                    title: "类型",
                    isshow: true,
                    detail: [

                    ]
                },
            ],
        };
    },
    filters: {
        formatDate: function (time) {
            if (time != null && time != "") {
                var date = new Date(time);
                return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
            } else {
                return "";
            }
        },
        formatBoolean: function (bool) {
            if (bool != null) {
                return bool ? "是" : "否";
            } else {
                return "";
            }
        },
    },
    methods: {
        async getdata() {
            this.tableData = this.$options.data().tableData
            //获取周期
            var start_dict = {
                'start_date': formatTimeToStr(this.setInfo.up_date, "yyyy-MM-dd")
            }
            let res = await getCycleConfirmDate(start_dict)
            var i = 1
            const datelist = []
            var end_dateStr
            var start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000, "yyyy-MM-dd")
            if (res.data.list) {
                end_dateStr = formatTimeToStr(new Date(res.data.list[0].end_date).getTime() + 86400000, "yyyy-MM-dd")
            } else {
                end_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * 22, "yyyy-MM-dd")

            }
            while (start_dateStr != end_dateStr) {
                i = i + 1
                var month_data = formatTimeToStr(new Date(start_dateStr), "MM-dd")
                datelist.push(month_data)
                start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * i, "yyyy-MM-dd")
            }
            for (var k in datelist) {
                var table_dict = {
                    title: datelist[k],
                    isshow: true,
                    detail: [],
                }
                this.tableData.push(table_dict)
            }
            this.cycleList = datelist
            //获取类型
            switch (this.setInfo.department_value) {
                case '杉贸本部':
                    this.tableData[0].detail = [
                        "支出-融资还款",
                        "支出-借款拆出",
                        "支出-拆借还款",
                        "支出-融资存款",
                        "支出-回报存款",
                        "支出-财务手续费",
                        "支出-其他",
                        "支出小计",

                        "收入-保证金及收益退回",
                        "收入-借款拆入",
                        "收入-拆借还款",
                        "收入-融资贴现",
                        "收入-回报贴现",
                        "收入-其他",
                        "收入小计"
                    ]
                    break
                case '总经理一室':
                    this.tableData[0].detail = [
                        "支出-业务存款",
                        "支出-业务还款",
                        "支出-财务手续费",
                        "支出-其他",
                        "支出小计",

                        "收入-利息收入",
                        "收入-咨询费",
                        "收入-回报贴现",
                        "收入-其他",
                        "收入小计"
                    ]
                    break
                case '智维本部':
                    this.tableData[0].detail = [
                        "支出-借款拆出",
                        "支出-拆借还款",
                        "支出-其他",
                        "支出小计",

                        "收入-借款拆入",
                        "收入-拆借还款",
                        "收入-其他",
                        "收入小计"
                    ]
                    break
                case '全部':
                    this.tableData[0].detail = [
                        "支出-有色事业部",
                        "支出-黑色事业部",
                        "支出-农产事业部",
                        "支出-原料事业部",
                        "支出-钢材事业部",
                        "支出-总经理一室",
                        "支出-总经理二室",
                        "支出-交易部",
                        "支出-杉贸本部",
                        "支出小计",
                        "收入-有色事业部",
                        "收入-黑色事业部",
                        "收入-农产事业部",
                        "收入-原料事业部",
                        "收入-钢材事业部",
                        "收入-总经理一室",
                        "收入-总经理二室",
                        "收入-交易部",
                        "收入-杉贸本部",
                        "收入小计",
                        "杉贸-当日总收支",
                        "杉贸-当日资金余额",
                        "智维-当日资金余额",
                        "合计资金余额",
                        "可质押仓单余额-杉贸",
                        "可质押仓单余额-智维",
                        "合计可质押余额",
                        "杉贸银行授信余额",
                        "资源总计(资金+可质押+授信余额)"
                    ]
                    break
                default:
                    this.tableData[0].detail = [
                        "支出-现货采购-采购货款",
                        "支出-现货采购-卖货解质押",
                        "支出-现货采购-流量业务",
                        "支出-现货采购-协助集团融资借款",
                        "支出-现货采购-到期付汇",
                        "支出-现货采购-开证保证金",
                        "支出-现货采购-税金", "",
                        "支出-期货入金-开仓",
                        "支出-期货入金-银行解质押",
                        "支出-期货入金-卖交割解质押",
                        "支出-期货入金-买交割",
                        "支出-期货入金-提保入金",

                        "支出小计",
                        "收入-现货销售-正常销售",
                        "收入-现货销售-买货质押",
                        "收入-现货销售-流量业务",
                        "收入-现货销售-协助集团融资借款回款",
                        "收入-现货销售-开保证金及收益退回", "",
                        "收入-期货出金-平仓",
                        "收入-期货出金-银行质押",
                        "收入-期货出金-买交割质押",
                        "收入-期货出金-卖交割",
                        "收入-期货出金-释放保证金",
                        "收入小计"
                    ]
            }
            //table数据绑定
            if (this.setInfo.department_value != "全部") {
                let confirm_values = {
                    department_code: this.setInfo.department_value,
                    create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                }
                let res_data = await getCycleData(confirm_values)
                for (var j in res_data.data.list) {
                    var list = res_data.data.list[j]
                    var list_key = this.cycleList.indexOf(formatTimeToStr(new Date(list.in_data), "MM-dd")) + 1
                    if (list_key > 0) {
                        this.$set(this.tableData[list_key].detail, this.fill_confirm_data(list), list.in_value)
                    }
                }
                //绑定加合数据
                if (this.setInfo.department_value == "总经理一室" || this.setInfo.department_value == "杉贸本部" || this.setInfo.department_value == "智维本部") {
                    if (this.setInfo.department_value == '杉贸本部') {
                        for (var q in this.tableData) {
                            if (q > 0) {
                                this.$set(this.tableData[q].detail, 7, this.sum_details(this.tableData[q].detail)['out_sum'])
                                this.$set(this.tableData[q].detail, 14, this.sum_details(this.tableData[q].detail)['in_sum'])
                            }
                        }
                    }
                    if (this.setInfo.department_value == '总经理一室') {
                        for (var w in this.tableData) {
                            if (w > 0) {
                                this.$set(this.tableData[w].detail, 4, this.sum_details(this.tableData[w].detail)['out_sum'])
                                this.$set(this.tableData[w].detail, 9, this.sum_details(this.tableData[w].detail)['in_sum'])
                            }
                        }
                    }
                    if (this.setInfo.department_value == '智维本部') {
                        for (var e in this.tableData) {
                            if (e > 0) {
                                this.$set(this.tableData[e].detail, 3, this.sum_details(this.tableData[e].detail)['out_sum'])
                                this.$set(this.tableData[e].detail, 7, this.sum_details(this.tableData[e].detail)['in_sum'])
                            }
                        }
                    }
                } else {
                    for (var p in this.tableData) {
                        if (p > 0) {
                            this.$set(this.tableData[p].detail, 13, this.sum_details(this.tableData[p].detail)['out_sum'])
                            this.$set(this.tableData[p].detail, 25, this.sum_details(this.tableData[p].detail)['in_sum'])
                        }
                    }
                }
            }





        },
        //填写数据确认table
        fill_confirm_data(data) {
            if (this.setInfo.department_value == "总经理一室" || this.setInfo.department_value == "杉贸本部" || this.setInfo.department_value == "智维本部") {
                if (this.setInfo.department_value == '杉贸本部')
                    switch (data.type_code) {
                                            case 'rzhk':
                            return 0
                        case 'jkcc':
                            return 1
                        case 'outcjhk':
                            return 2
                        case 'rzck':
                            return 3
                        case 'hbck':
                            return 4
                        case 'cwsxf':
                            return 5
                            case 'outqt':
                            return 6

                        case 'bzjjsyth':
                            return 8
                        case 'jkcr':
                            return 9
                        case 'incjhk':
                            return 10
                        case 'rztx':
                            return 11
                        case 'hbtx':
                            return 12
                             case 'inqt':
                            return 13
                        default:
                            return
                    }
                if (this.setInfo.department_value == '总经理一室')
                    switch (data.type_code) {
                       case 'ywck':
                            return 0
                        case 'ywhk':
                            return 1
                        case 'cwsxf':
                            return 2
                             case 'outqt':
                            return 3

                        case 'lxsr':
                            return 5
                        case 'zxf':
                            return 6
                        case 'hbtx':
                            return 7
                             case 'inqt':
                            return 8
                        default:
                            return
                    }
                if (this.setInfo.department_value == '智维本部')
                    switch (data.type_code) {
                                   case 'jkcc':
                            return 0
                        case 'outcjhk':
                            return 1
                             case 'outqt':
                            return 2

                        case 'jkcr':
                            return 4
                        case 'incjhk':
                            return 5
                             case 'inqt':
                            return 6

                        default:
                            return
                    }
            } else {
                switch (data.type_code) {
                    case 'cghk':
                        return 0
                    case 'mhjzy':
                        return 1
                    case 'outllyw':
                        return 2
                    case 'xzjtrzjk':
                        return 3
                    case 'dqfh':
                        return 4
                    case 'kzbzj':
                        return 5
                    case 'sj':
                        return 6

                    case 'kc':
                        return 8
                    case 'yhjzy':
                        return 9
                    case 'mjgjzy':
                        return 10
                    case 'outmjg':
                        return 11
                    case 'tbrj':
                        return 12

                    case 'zcxs':
                        return 14
                    case 'mhzy':
                        return 15
                    case 'inllyw':
                        return 16
                    case 'xzjtjkhk':
                        return 17
                    case 'kzbzjth':
                        return 18

                    case 'pc':
                        return 20
                    case 'yhzy':
                        return 21
                    case 'mjgzy':
                        return 22
                    case 'inmjg':
                        return 23
                    case 'sfbzj':
                        return 24
                    default:
                        return
                }
            }

        },
        //计算加合
        sum_details(details_data) {
            switch (this.setInfo.department_value) {
                case '杉贸本部':
                    var out_sum2 = 0
                    var in_sum2 = 0
                    for (var i2 = 0; i2 < 7; i2++) {
                        if (details_data[i2]) {
                            out_sum2 += details_data[i2]
                        }
                    }
                    for (var j2 = 8; j2 < 14; j2++) {
                        if (details_data[j2]) {
                            in_sum2 += details_data[j2]
                        }
                    }
                    var res_dict2 = {
                        in_sum: in_sum2,
                        out_sum: out_sum2
                    }
                    return res_dict2
                case '总经理一室':
                    var out_sum3 = 0
                    var in_sum3 = 0
                    for (var i3 = 0; i3 < 4; i3++) {
                        if (details_data[i3]) {
                            out_sum3 += details_data[i3]
                        }
                    }
                    for (var j3 = 5; j3 < 9; j3++) {
                        if (details_data[j3]) {
                            in_sum3 += details_data[j3]
                        }
                    }
                    var res_dict3 = {
                        in_sum: in_sum3,
                        out_sum: out_sum3
                    }
                    return res_dict3
                case '智维本部':
                    var out_sum4 = 0
                    var in_sum4 = 0
                    for (var i4 = 0; i4 < 3; i4++) {
                        if (details_data[i4]) {
                            out_sum4 += details_data[i4]
                        }
                    }
                    for (var j4 = 4; j4 < 7; j4++) {
                        if (details_data[j4]) {
                            in_sum4 += details_data[j4]
                        }
                    }
                    var res_dict4 = {
                        in_sum: in_sum4,
                        out_sum: out_sum4
                    }
                    return res_dict4
                default:
                    var out_sum = 0
                    var in_sum = 0
                    for (var i = 0; i < 13; i++) {
                        if (details_data[i]) {
                            out_sum += details_data[i]
                        }
                    }
                    for (var j = 14; j < 25; j++) {
                        if (details_data[j]) {
                            in_sum += details_data[j]
                        }
                    }
                    var res_dict = {
                        in_sum: in_sum,
                        out_sum: out_sum
                    }
                    return res_dict
            }
        },

        cellStyle() {
            return "text-align:center";
        },
        rowClass() {
            return "text-align:center";
        },

    },
    created() {
        this.getdata()
    }
};
</script>

<style>
.is-selected {
    color: #1989fa;
}

.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.grid-content {
    border-radius: 4px;
    min-height: 36px;
}

.myTable {
    border-collapse: collapse;
    margin: 0 auto;
    text-align: center;
}


.myTable td,
.myTable th {
    border: 1px solid #cad9ea;
    color: #666;
    height: 60px;
}
</style>
