<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="填报日期">
                    <el-date-picker v-model="setInfo.up_date" type="date" placeholder="Pick a day" :size="'default'"
                        @change="getCycle" />
                </el-form-item>

                <el-form-item>
                    <el-popover placement="bottom" width="240" v-model="copy_visible">
                        <!-- <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> -->
                        <el-form-item>
                            <el-date-picker v-model="down_date" type="date" placeholder="Pick a day"
                                :size="'default'" />
                        </el-form-item>
                        <!-- </el-form> -->

                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="copy_visible=false">取消</el-button>
                            <el-button type="primary" size="mini" @click="copyCycle">确定</el-button>
                        </div>
                        <el-button slot="reference" icon="copyDocument" type="primary">复制</el-button>

                    </el-popover>
                </el-form-item>
                <el-form-item label="实际授信余额">
                    <el-input autocomplete="off" v-model="setInfo.remaining_sum" type="number" disabled></el-input>
                </el-form-item>

                <el-form-item>
                    <el-popover placement="bottom" width="240" v-model="change_visible">
                        <!-- <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> -->
                        <el-form-item>
                            <el-form-item label="实际授信余额">
                                <el-input autocomplete="off" v-model.number="setInfo.remaining_sum"></el-input>
                            </el-form-item>
                        </el-form-item>
                        <!-- </el-form> -->

                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="change_visible=false">取消</el-button>
                            <el-button type="primary" size="mini" @click="save_remaining_sum">确定</el-button>
                        </div>
                        <el-button slot="reference" icon="copyDocument" type="primary">变更</el-button>

                    </el-popover>
                </el-form-item>


            </el-form>
        </div>
        <div class="gva-table-box">
            <el-calendar>
                <template #dateCell="{ data }">
                    <div class="confirm" :class="addClass(data.day)">
                        <ul>
                            <li style="padding-top:2px;">

                                <p :class="data.isSelected ? 'is-selected' : ''" @click="openDraw(data)">
                                    {{ data.day.split("-").slice(1).join("-") }}
                                    {{ data.isSelected ? "✔️" : "" }}
                                </p>
                            </li>
                            <li v-if="(new Date(data.day).getTime() > new Date().getTime()) && in_out[data.day] != undefined"
                                style="padding-top:2px;" :class="addClassText(data.day, '支出')">
                                授信释放：{{ in_out[data.day] == undefined ? '0' : in_out[data.day][1] }}万</li>
                            <li v-if="(new Date(data.day).getTime() > new Date().getTime()) && in_out[data.day] != undefined"
                                style="padding-top:2px ;" :class="addClassText(data.day, '支出')">
                                授信占用：{{ in_out[data.day] == undefined ? '0' : in_out[data.day][0] }}万</li>

                            <li v-show="(new Date(data.day).getTime() > new Date().getTime()) && in_out[data.day] != undefined"
                                style="padding-top:2px ;" :class="addClassText(data.day, '收入')">
                                授信余额：{{ count_remaining_sum(data.day) }}万</li>
                        </ul>
                    </div>
                </template>
            </el-calendar>
        </div>
        <el-dialog :visible.sync="bank_show" width="450px">
            <template #title>
                <h4>
                    授信余额录入&nbsp;&nbsp;{{ choseDate }}&nbsp;&nbsp;&nbsp;{{
                            setInfo.department_value
                    }}&nbsp;&nbsp;&nbsp;{{ setInfo.role_value }}
                </h4>
            </template>
            <el-form :model="form">
                <el-form-item label="授信释放：" label-width="120px">
                    <el-input v-model="form.sxsf" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="授信占用：" label-width="120px">
                    <el-input v-model="form.sxzy" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancelClick">取 消</el-button>
                <el-button type="primary" @click="confirmClick">提 交</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import {
    createBudgetDatas,
    clickBudgetData,
    getCycleData,
} from "@/api/internalSystem/fund/fund_budget_data"; //  此处请自行替换地址
import {
    createRemainingSum,
    getOneRemainingSum,
} from "@/api/internalSystem/fund/remainingSum";
import {
    getCycleConfirmDate,
} from "@/api/internalSystem/fund/fund_confirm_date";  //  此处请自行替换地址

export default {
    name: "BankData",
    mixins: [infoList],

    data() {
        return {
            //录入数据参数
            form: { "sxsf": 0, "sxzy": 0 },
            bank_show: false,
            down_date: new Date().getTime() - 86400000,
            copy_visible: false,
            change_visible: false,
            choseDate: new Date(),
            date_list: [],
            change_remaining_sum: 0,
            setInfo: {
                department_value: "资金部门",
                role_value: "资金专员",
                up_date: new Date().getTime(),
                remaining_sum:0
            },
            in_out: {},
            reslist: [],
            active: false,
            is_allow: 0,


        };
    },
    filters: {
        formatDate: function (time) {
            if (time != null && time != "") {
                var date = new Date(time);
                return formatTimeToStr(date, "yyyy-MM-dd");
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
        //计算每日余额
        count_remaining_sum(data) {
            // console.log(data)
            var add = 0
            var min = 0
            // console.log(this.reslist,new Date(data))
            for (var i in this.reslist) {
                if (new Date(formatTimeToStr(new Date(this.reslist[i]['in_data']), "yyyy-MM-dd")).getTime() <= new Date(formatTimeToStr(new Date(data), "yyyy-MM-dd")).getTime()) {
                    if (this.reslist[i]['type_code'] == "sxzy") {
                        //授信占用
                        min = min + this.reslist[i]['in_value']
                    } else {
                        add = add + this.reslist[i]['in_value']
                    }
                }
            }
            // console.log(add, min, this.setInfo.remaining_sum)
            return this.setInfo.remaining_sum + add - min
        },
        //保存每日授信余额值
        async save_remaining_sum() {
            var remaining_dict = {
                inValue: this.setInfo.remaining_sum,
                company: "杉贸",
                createUser: JSON.parse(localStorage.getItem("vuex"))['user']['userInfo']['nickName'],
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                type: "授信余额",
            }
            let res = await createRemainingSum(remaining_dict);
           if (res.code == 0) {
                this.$message({
                    type: "success",
                    message: "余额变更成功"
                })
                this.active = true
                this.confirm_dialog = false
            } else {
                this.$message({
                    type: "error",
                    message: "余额变更失败"
                })
            }
        },
        // 获取当前授信余额
        async get_remaining_sum() {
            var search_dict = {
                type: "授信余额",
                upDate: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                Company: "杉贸"
            }
            let res = await getOneRemainingSum(search_dict)
            
            if (formatTimeToStr(new Date(res.data.list[0]['upDate']), "yyyy-MM-dd") != formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd")) {
                let search_value = {
                    department_code: this.setInfo.department_value,
                    create_role: this.setInfo.role_value,
                    in_data: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                    up_date: formatTimeToStr(new Date(res.data.list[0]['upDate']), "yyyy-MM-dd"),
                }
                let res_list = await clickBudgetData(search_value)
                var add = 0
                var min = 0
                for (var i in res_list.data.list) {
                        if (res_list.data.list[i]['type_code'] == "sxzy") {
                            //授信占用
                            min = min + res_list.data.list[i]['in_value']
                        } else {
                            add = add + res_list.data.list[i]['in_value']
                        }
                }
                this.setInfo.remaining_sum =res.data.list[0]['inValue'] + add - min
            }else{
                 this.setInfo.remaining_sum =res.data.list[0]['inValue']
            }


        },
        //复制数据
        async copyCycle() {
            let confirm_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_date: formatTimeToStr(new Date(this.down_date), "yyyy-MM-dd"),
            }
            let res = await getCycleData(confirm_value)
            var up_cycle_data = []
            for (var i in res.data.list) {
                if (new Date(formatTimeToStr(new Date(res.data.list[i]['in_data']), "yyyy-MM-dd")) > new Date(this.setInfo.up_date).getTime()) {
                    var up_dict = {
                        in_data: formatTimeToStr(new Date(res.data.list[i]['in_data']), "yyyy-MM-dd"),
                        type_code: res.data.list[i]['type_code'],
                        in_value: Number(res.data.list[i]['in_value']),
                        comment: res.data.list[i]['comment'],
                        currency: "RMB",
                        department_code: res.data.list[i]['department_code'],
                        create_user: res.data.list[i]['create_user'],
                        create_role: res.data.list[i]['create_role'],
                        up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd")
                    }
                    up_cycle_data.push(up_dict)
                }
            }
            let up_res = await createBudgetDatas(up_cycle_data);
            if (up_res.code == 0) {
                this.$message({
                    type: "success",
                    message: "复制成功"
                })
                this.getCycle()
                this.copy_visible = false
            }
            this.save_remaining_sum()
        },
        //日历颜色样式
        addClass(days) {
            if (this.date_list.includes(days)) {
                return "will-date"
            } else {
                return ""
            }
        },
        addClassText(days, type) {
            if (this.date_list.includes(days)) {
                if (type == '收入') { return "text-red" } else {
                    return "text-green"
                }
            } else {
                return ""
            }
        },
        //获取日历周期
        async getCalenderDate() {
            var start_dict = {
                'start_date': formatTimeToStr(this.setInfo.up_date, "yyyy-MM-dd")
            }
            let res = await getCycleConfirmDate(start_dict)
            var i = 1
            const datelist = []
            var end_dateStr
            var start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000, "yyyy-MM-dd")
            // console.log(res)
            if (res.data.list) {
                end_dateStr = formatTimeToStr(new Date(res.data.list[0].end_date).getTime() + 86400000, "yyyy-MM-dd")
            } else {
                end_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * 22, "yyyy-MM-dd")

            }
            while (start_dateStr != end_dateStr) {
                i = i + 1
                datelist.push(start_dateStr)
                start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * i, "yyyy-MM-dd")
            }
            this.date_list = datelist
        },

        //获取周期数据
        async getCycle() {
            let confirm_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            this.in_out = {}
            let res = await getCycleData(confirm_value)
            // console.log(res)
            //获取有录入数据的日期
            this.reslist = res.data.list
            var list_date = []
            for (var i in res.data.list) {
                list_date.push(formatTimeToStr(new Date(res.data.list[i]['in_data']), "yyyy-MM-dd"))
            }
            var one_list = Array.from(new Set(list_date))
            for (var k in one_list) {
                //申明变量  
                var key = one_list[k]
                // console.log(one_list)
                var out_data = 0
                var in_data = 0
                var day_lists = []
                for (var j in res.data.list) {
                    if (formatTimeToStr(new Date(res.data.list[j]['in_data']), "yyyy-MM-dd") == key) {
                        day_lists.push(res.data.list[j])
                    }
                }
                for (var l in day_lists) {
                    if (day_lists[l]['type_code'] == "sxzy") {
                        //授信占用
                        out_data = out_data + day_lists[l]['in_value']
                    } else {
                        in_data = in_data + day_lists[l]['in_value']
                    }
                }
                this.$set(this.in_out, key, [out_data, in_data])
            }
            this.get_remaining_sum()

        },

        //打开添加数据的抽屉
        async openDraw(data) {

            this.choseDate = data.day
            let search_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                in_data: this.choseDate,
                up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }

            let res = await clickBudgetData(search_value)

            if (res.data.list) {
                this.arrangeReturnData(res.data.list)
                this.bank_show = true;
            } else {
                this.bank_show = true;
            }

            // console.log(data)
        },
        //关闭抽屉
        cancelClick() {
            this.bank_show = false;
            this.form = this.$options.data().form
        },

        async confirmClick() {
            let up_value = this.arrangeData(this.form);
            if (up_value.length != 0) {
                let res = await createBudgetDatas(up_value);
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "填报成功"
                    })
                    this.getCycle()
                    this.cancelClick()
                }
            } else {
                this.$message({
                    type: "error",
                    message: "未填写数据"
                })

            }
        },

        //转换字典为list
        confirmEnding(string, target) {
            return string.substr(-target.length) === target ? true : false;
        },
        //将表单数据变成一条一条得数据类型
        arrangeData(data) {
            let lists = [];
            for (var k in data) {
                if (!this.confirmEnding(k, "bz")) {
                    if (Number(data[k]) != 0) {
                        const one = {
                            in_data: this.choseDate,
                            type_code: k,
                            in_value: Number(data[k]),
                            comment: data[k + "bz"],
                            currency: "RMB",
                            department_code: this.setInfo.department_value,
                            create_user: this.setInfo.role_value,
                            create_role: this.setInfo.role_value,
                            up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd")
                        };
                        lists.push(one);
                    }
                }
            }
            return lists;
        },
        arrangeReturnData(data) {

            data.forEach((item) => {
                this.form[item.type_code] = item.in_value
            })
        },
    },

    created() {
        this.getCalenderDate()
        this.getCycle()
    },

}

</script>

<style>
.my-label {
    background: #E1F3D8;
}

.my-content {
    background: #FDE2E2;
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

.confirm {
    height: 100%;
    width: 100%;
    text-align: center;
}

.will-date {
    background-color: #eee;
}

.el-calendar-table .el-calendar-day {
    padding: 2px;
}

.text-green {
    color: green;
}

.text-red {
    color: red;
}

.backgroud-green {
    background-color: green;
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
