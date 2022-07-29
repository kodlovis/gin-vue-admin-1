<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <!-- <el-form-item label="部门">
                    <el-select v-model="setInfo.department_value" class="m-2" placeholder="Select" @change="getdata">
                        <el-option v-for="item in department_options" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item> -->
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
                :row-class-name="tableRowClassName" :header-cell-style="rowClass" cell-class-name="lm-cell"
                empty-text="暂无数据" tooltip-effect="dark" :span-method="objectSpanMethod">
                <template v-for="(item, index) in tableData">
                    <el-table-column :key="index" :label="item.title" v-if="['平台', '大类', '部门'].includes(item.title)"
                        fixed="left" width="200">
                        <template slot-scope="scope">
                            <div>{{ item.detail[scope.$index] }}</div>
                        </template>
                    </el-table-column>
                </template>
                <!-- v-if="!['平台', '大类', '部门'].includes(item.title)" -->
                <template v-for="(item, index) in tableData">
                    <el-table-column :key="index" :label="item.title" v-if="!['平台', '大类', '部门'].includes(item.title)">
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
    getOneRemainingSum,
    getwarehouse
} from "@/api/internalSystem/fund/remainingSum";
import {
    getCycleConfirmDate
} from "@/api/internalSystem/fund/fund_confirm_date"; //  此处请自行替换地址
import {
    getAll,
    getCycleData
} from "@/api/internalSystem/fund/fund_budget_data";

export default {
    name: "BudgetReport",
    mixins: [infoList],

    data() {
        return {
            //授信余额
            credit: 0,
            //授信余额变化列表
            credit_list: [],
            //初始资金余额
            fund_sm: 0,
            fund_zw: 0,
            in_out: {},
            //质押仓单列表
            warehouse_list: [],

            fund_impawn: [],
            cycleList: [],
            out_type: ['cghk', 'mhjzy', 'outllyw', 'xzjtrzjk', 'dqfh', 'kzbzj', 'sj', 'kc', 'yhjzy', 'mjgjzy', 'outmjg', 'tbrj', 'jjfz', 'ywck', 'ywhk', 'cwsxf', 'rzhk', 'jkcc', 'outcjhk', 'rzck', 'hbck', 'outqt'],
            setInfo: {
                department_value: "全部",
                up_date: new Date().getTime(),
            },
            isChooseCols: false,
            tableData: [
                {
                    title: "平台",
                    isshow: true,
                    detail: [
                        "杉贸",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",

                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",
                        "",

                        "智维",
                        "",
                        "",//支出小计
                        "",
                        "",
                        "",//收入小计
                        "",//当日资金余额
                        "合计1资金余额",//合计资金余额


                        "可质押仓单",//杉贸余额
                        "",//智维余额
                        "合计2可质押余额",//合计2质押余额


                        "授信",//授信余额

                        "资源总计(资金+可质押+授信)",//资源总计


                    ]
                },
                {
                    title: "大类",
                    isshow: true,
                    detail: [
                        "支出",
                        "",
                        "",
                        "",

                        "",

                        "",
                        "",
                        "",
                        "",

                        "小计",

                        "",
                        "",
                        "",
                        "",

                        "收入",

                        "",
                        "",
                        "",
                        "",

                        "小计",
                        "当日总收支",
                        "当日资金余额",



                        "支出",
                        "",
                        "小计",//支出小计
                        "收入",
                        "",
                        "小计",//收入小计
                        "当日资金余额",//当日资金余额
                        "",//合计资金余额


                        "杉贸余额",//杉贸余额
                        "智维余额",//智维余额
                        "",//合计2质押余额

                        "授信余额",//授信余额

                        "",//资源总计

                    ]
                },
                {
                    title: "部门",
                    isshow: true,
                    detail: [
                        "有色事业部",
                        "黑色事业部",
                        "农产事业部",
                        "原料事业部",
                        "钢材事业部",
                        "总经理一室",
                        "总经理二室",
                        "交易部",
                        "杉贸本部",
                        "",//支出小计
                        "有色事业部",
                        "黑色事业部",
                        "农产事业部",
                        "原料事业部",
                        "钢材事业部",
                        "总经理一室",
                        "总经理二室",
                        "交易部",
                        "杉贸本部",
                        "",//收入小计
                        "",//当日总收支
                        "",//当日资金余额


                        "钢材事业部",
                        "智维本部",
                        "",//支出小计
                        "钢材事业部",
                        "智维本部",
                        "",//收入小计
                        "",//当日资金余额
                        "",//合计资金余额


                        "",//杉贸余额
                        "",//智维余额
                        "",//合计2质押余额


                        "",//授信余额

                        "",//资源总计


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
            }
            else {
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
        tableRowClassName({ rowIndex }) {
            //  console.log({ row, column, rowIndex, columnIndex })
            if ([34].includes(rowIndex)) {
                return 'success-row';
            } else if ([32, 29,33].includes(rowIndex)) {
                return 'warning-row';
            }else if ([21,28].includes(rowIndex)){
                return 'warnings-row';
            }
            return '';
        },
        objectSpanMethod({ rowIndex, columnIndex }) {
            // console.log({ row, column, rowIndex, columnIndex })
            if (rowIndex == 0 && columnIndex == 0) {
                return [22, 1]
            }
            if (rowIndex == 0 && columnIndex == 1) {
                return {
                    rowspan: 9,
                    colspan: 1,
                }
            }
            if ([1, 2, 3, 4, 5, 6, 7, 8].includes(rowIndex) && columnIndex == 1) {
                return {
                    rowspan: 0,
                    colspan: 0,
                }
            }

            if (rowIndex == 22 && columnIndex == 0) {
                return {
                    rowspan: 7,
                    colspan: 1,
                }
            }
            if (rowIndex == 22 && columnIndex == 1) {
                return {
                    rowspan: 2,
                    colspan: 1,
                }
            }
            if (rowIndex == 25 && columnIndex == 1) {
                return {
                    rowspan: 2,
                    colspan: 1,
                }
            }
            if (rowIndex == 30 && columnIndex == 0) {
                return {
                    rowspan: 2,
                    colspan: 1,
                }
            }
            if (rowIndex == 33 && columnIndex == 0) {
                return {
                    rowspan: 1,
                    colspan: 1,
                }
            }
            if (rowIndex == 29 && columnIndex == 0) {
                return {
                    rowspan: 1,
                    colspan: 3,
                }
            }
            if (rowIndex == 29 && [1, 2].includes(columnIndex)) {
                return {
                    rowspan: 0,
                    colspan: 0,
                }
            }
            if (rowIndex == 32 && columnIndex == 0) {
                return {
                    rowspan: 1,
                    colspan: 3,
                }
            }
            if (rowIndex == 32 && [1, 2].includes(columnIndex)) {
                return {
                    rowspan: 0,
                    colspan: 0,
                }
            }
            if (rowIndex == 34 && columnIndex == 0) {
                return {
                    rowspan: 1,
                    colspan: 3,
                }
            }
            if (rowIndex == 34 && [1, 2].includes(columnIndex)) {
                return {
                    rowspan: 0,
                    colspan: 0,
                }
            }
            // if (rowIndex == 33 && columnIndex == 0) {
            //     return {
            //         rowspan: 1,
            //         colspan: 1,
            //     }
            // }
            // if ((33 <= rowIndex <= 34 && columnIndex == 0)||( rowIndex == 29 && columnIndex == 0)||(rowIndex == 32 && columnIndex == 0)) {
            //     return {
            //         rowspan: 1,
            //         colspan: 1,
            //     }
            // }
            if ((0 < rowIndex <= 21 && columnIndex == 0) || (rowIndex == 23 && columnIndex == 1) || (rowIndex == 26 && columnIndex == 1) || (22 < rowIndex <= 28 && columnIndex == 0) || (30 < rowIndex <= 31 && columnIndex == 0)) {
                return { rowspan: 0, colspan: 0 };
            }


        },

        async getdata() {
            // 获取初始各个余额

            this.tableData = this.$options.data().tableData
            this.gerSum()
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
                    detail: new Array(34),
                }
                this.tableData.push(table_dict)
            }
            this.cycleList = datelist
            // table数据绑定
            let confirm_values = {
                create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            let res_data = await getAll(confirm_values)
            // console.log("该日所有数据", res_data, this.cycleList)
            this.cycleList.forEach(((item, i) => {
                // console.log(item, i)
                var out_ys = new Number
                var out_hs = new Number
                var out_nc = new Number
                var out_yl = new Number
                var out_gc = new Number
                var out_z1 = new Number
                var out_z2 = new Number
                var out_jy = new Number
                var out_sm = new Number

                var in_ys = new Number
                var in_hs = new Number
                var in_nc = new Number
                var in_yl = new Number
                var in_gc = new Number
                var in_z1 = new Number
                var in_z2 = new Number
                var in_jy = new Number
                var in_sm = new Number

                var out_zwgc = new Number
                var out_zw = new Number

                var in_zwgc = new Number
                var in_zw = new Number
                //每一次循环筛选出本日期的所有数据
                var daylist = []
                for (var o in res_data.data.list) {
                    if (item == formatTimeToStr(new Date(res_data.data.list[o].in_data), "MM-dd")) {
                        daylist.push(res_data.data.list[o])
                    }
                }
                // console.log("我是每日数据", daylist)
                if (daylist.length != 0) {
                    for (var l in daylist) {
                        if (daylist[l].department_code == "有色事业部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_ys = daylist[l].in_value + out_ys
                            } else {
                                in_ys = daylist[l].in_value + in_ys
                            }
                        }
                        if (daylist[l].department_code == "黑色事业部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_hs = daylist[l].in_value + out_hs
                            } else {
                                in_hs = daylist[l].in_value + in_hs
                            }
                        }
                        if (daylist[l].department_code == "农产事业部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_nc = daylist[l].in_value + out_nc
                            } else {
                                in_nc = daylist[l].in_value + in_nc
                            }
                        }
                        if (daylist[l].department_code == "原料事业部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_yl = daylist[l].in_value + out_yl
                            } else {
                                in_yl = daylist[l].in_value + in_yl
                            }
                        }
                        if (daylist[l].department_code == "钢材事业部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_gc = daylist[l].in_value + out_gc
                            } else {
                                in_gc = daylist[l].in_value + in_gc
                            }
                        }
                        if (daylist[l].department_code == "总经理二室") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_z2 = daylist[l].in_value + out_z2
                            } else {
                                in_z2 = daylist[l].in_value + in_z2
                            }
                        }
                        if (daylist[l].department_code == "总经理一室") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_z1 = daylist[l].in_value + out_z1
                            } else {
                                in_z1 = daylist[l].in_value + in_z1
                            }
                        }
                        if (daylist[l].department_code == "杉贸本部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_sm = daylist[l].in_value + out_sm
                            } else {
                                in_sm = daylist[l].in_value + in_sm
                            }
                        }
                        if (daylist[l].department_code == "交易部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_jy = daylist[l].in_value + out_jy
                            } else {
                                in_jy = daylist[l].in_value + in_jy
                            }
                        }
                        if (daylist[l].department_code == "智维钢材") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_zwgc = daylist[l].in_value + out_zwgc
                            } else {
                                in_zwgc = daylist[l].in_value + in_zwgc
                            }
                        }
                        if (daylist[l].department_code == "智维本部") {
                            if (this.out_type.includes(daylist[l].type_code)) {
                                out_zw = daylist[l].in_value + out_zw
                            } else {
                                in_zw = daylist[l].in_value + in_zw
                            }
                        }
                    }
                }
                var ssmy_out = out_ys + out_hs + out_nc + out_yl + out_gc + out_z1 + out_z2 + out_jy + out_sm
                var ssmy_in = in_ys + in_hs + in_nc + in_yl + in_gc + in_z1 + in_z2 + in_jy + in_sm
                var ssmy_all = ssmy_in - ssmy_out
                var zw_out = out_zwgc + out_zw
                var zw_in = in_zwgc + in_zw
                this.$set(this.tableData[i + 3].detail, 0, out_ys)
                this.$set(this.tableData[i + 3].detail, 1, out_hs)
                this.$set(this.tableData[i + 3].detail, 2, out_nc)
                this.$set(this.tableData[i + 3].detail, 3, out_yl)
                this.$set(this.tableData[i + 3].detail, 4, out_gc)
                this.$set(this.tableData[i + 3].detail, 5, out_z1)
                this.$set(this.tableData[i + 3].detail, 6, out_z2)
                this.$set(this.tableData[i + 3].detail, 7, out_jy)
                this.$set(this.tableData[i + 3].detail, 8, out_sm)
                this.$set(this.tableData[i + 3].detail, 9, ssmy_out)
                this.$set(this.tableData[i + 3].detail, 10, in_ys)
                this.$set(this.tableData[i + 3].detail, 11, in_hs)
                this.$set(this.tableData[i + 3].detail, 12, in_nc)
                this.$set(this.tableData[i + 3].detail, 13, in_yl)
                this.$set(this.tableData[i + 3].detail, 14, in_gc)
                this.$set(this.tableData[i + 3].detail, 15, in_z1)
                this.$set(this.tableData[i + 3].detail, 16, in_z2)
                this.$set(this.tableData[i + 3].detail, 17, in_jy)
                this.$set(this.tableData[i + 3].detail, 18, in_sm)
                this.$set(this.tableData[i + 3].detail, 19, ssmy_in)
                this.$set(this.tableData[i + 3].detail, 20, ssmy_all)
                this.$set(this.tableData[i + 3].detail, 22, out_zwgc)
                this.$set(this.tableData[i + 3].detail, 23, out_zw)
                this.$set(this.tableData[i + 3].detail, 24, zw_out)
                this.$set(this.tableData[i + 3].detail, 25, in_zwgc)
                this.$set(this.tableData[i + 3].detail, 26, in_zw)
                this.$set(this.tableData[i + 3].detail, 27, zw_in)
            }))





        },
        async gerSum() {
            //获取余额
            var credit_dict = {
                type: "授信余额",
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                Company: "杉贸"
            }
            let credit_res = await getOneRemainingSum(credit_dict)
            var sm_dict = {
                type: "资金余额",
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                Company: "杉贸"
            }
            let sm_res = await getOneRemainingSum(sm_dict)
            var zw_dict = {
                type: "资金余额",
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                Company: "智维"
            }
            let zw_res = await getOneRemainingSum(zw_dict)
            // 获取授信变化数据
            let confirm_value = {
                department_code: "资金部门",
                create_role: "资金专员",
                create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            let res = await getCycleData(confirm_value)
            //获取质押仓单余额
            let warehouse_value = {
                type: "可质押余额",
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            let warehouse_res = await getwarehouse(warehouse_value)
            console.log(warehouse_res)
            this.warehouse_list = warehouse_res.data.list

            this.credit = credit_res.data.list[0].inValue
            this.fund_sm = sm_res.data.list[0].inValue
            this.fund_zw = zw_res.data.list[0].inValue
            this.credit_list = res.data.list
            this.cycleList.forEach(((item, i) => {
                this.$set(this.tableData[i + 3]?.detail, 33, this.count_remaining_sum(this.tableData[i + 3]?.title))
                this.$set(this.tableData[3]?.detail, 21, this.fund_sm + this.tableData[3]?.detail[20])
                // console.log(i + 4, this.tableData.length)
                this.$set(this.tableData[3]?.detail, 28, this.fund_zw + this.tableData[3]?.detail[27] - this.tableData[3]?.detail[23])
                if (i + 4 < this.tableData.length) {
                    this.$set(this.tableData[i + 4]?.detail, 21, this.tableData[i + 3]?.detail[21] + this.tableData[i + 4]?.detail[20])
                    this.$set(this.tableData[i + 4]?.detail, 28, this.tableData[i + 3]?.detail[28] + this.tableData[i + 4]?.detail[27] - this.tableData[i + 4]?.detail[23])
                }
                this.$set(this.tableData[i + 3]?.detail, 29, this.tableData[i + 3]?.detail[28] + this.tableData[i + 3]?.detail[21])

                this.$set(this.tableData[i + 3]?.detail, 30, this.getwarehouse("杉贸",this.tableData[i + 3]?.title))
                    this.$set(this.tableData[i + 3]?.detail, 31, this.getwarehouse("智维",this.tableData[i + 3]?.title))
                     this.$set(this.tableData[i + 3]?.detail, 32, this.tableData[i + 3]?.detail[30]+this.tableData[i + 3]?.detail[31])
                    this.$set(this.tableData[i + 3]?.detail, 34, this.tableData[i + 3]?.detail[29]+this.tableData[i + 3]?.detail[32]+this.tableData[i + 3]?.detail[33])
            }))

            // console.log(this.credit, this.fund_sm, this.fund_zw)
        },
        getwarehouse(type, day) {

            var now_value = 99999999999999;
            var now_key = 0
            for (let i = 0; i < this.warehouse_list.length; i++) {
                if(this.warehouse_list[i].company==type && ( new Date(day).getTime()- new Date(formatTimeToStr(new Date(this.warehouse_list[i].upDate), "MM-dd")).getTime())<=now_value &&  new Date(formatTimeToStr(new Date(this.warehouse_list[i].upDate), "MM-dd")).getTime()<=new Date(day).getTime()){
                    now_value = new Date(day).getTime() - new Date(formatTimeToStr(new Date(this.warehouse_list[i].upDate), "MM-dd")).getTime()
                    now_key = i
                }
            }
            return this.warehouse_list[now_key].inValue
        },

        cellStyle() {
            return "text-align:center";
        },
        rowClass() {
            return "text-align:center";
        },
        count_remaining_sum(data) {
            // console.log(data)
            var add = 0
            var min = 0
            // console.log(this.credit_list)
            for (var p in this.credit_list) {
                if (new Date(formatTimeToStr(new Date(this.credit_list[p]['in_data']), "MM-dd")).getTime() <= new Date(data).getTime()) {
                    if (this.credit_list[p]['type_code'] == "sxzy") {
                        //授信占用
                        min = min + this.credit_list[p]['in_value']
                    } else {
                        add = add + this.credit_list[p]['in_value']
                    }
                }
            }
            // console.log(add, min, this.credit_list, this.credit)
            return this.credit + add - min
        },

    },
    created() {
        // this.gerSum()
        this.getdata()
    },

};
</script>

<style>
.el-table .warning-row {
    background: rgb(255, 240, 212);
}

.el-table .success-row {
    background: #e1f5d7;
}
.el-table .warnings-row {
    background: rgb(204, 214, 240);
}


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

.admin-box .el-table td {
    padding: 0;
}

.el-table td,
.el-table th {
    padding: 0 0;
    min-width: 0;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    text-overflow: ellipsis;
    vertical-align: middle;
    position: relative;
    text-align: center;
}

.el-table .cell,
.el-table--border td:first-child .cell,
.el-table--border th:first-child .cell {
    padding-left: 0px;
}

.el-table .cell {
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: normal;
    word-break: break-all;
    line-height: 23px;
    padding-right: 0px;
}


.myTable td,
.myTable th {
    border: 1px solid #cad9ea;
    color: #666;
    height: 60px;
}
</style>
