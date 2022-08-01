<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="查询日期">
                    <el-date-picker v-model="setInfo.chose_date" type="date" placeholder="Pick a day" @change="getConfirm"
                        :size="'default'" />
                </el-form-item>
<!--     
                <el-form-item>
                    <el-button type="primary" icon="copyDocument" @click="getConfirm">查询</el-button>
                </el-form-item> -->
                <el-form-item style="float:right">
                    <el-switch  v-model="setInfo.is_confirm" active-color="#13ce66" inactive-color="#ff4949" @change="upConfirmSettings"
                        active-text="允许填报" inactive-text="拒绝填报">
                    </el-switch>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <el-table :data="tableData" style="width: 100%" :cell-style="cellStyle" :header-cell-style="rowClass">
                <el-table-column prop="date" label="日期"> </el-table-column>
                <el-table-column label="杉贸">
                    <el-table-column label="业务部门">
                        <el-table-column label="有色事业部">
                            <el-table-column prop="有色事业部-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="有色事业部-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                        <el-table-column label="黑色事业部">
                            <el-table-column prop="黑色事业部-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="黑色事业部-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                                     <el-table-column label="钢材事业部">
                            <el-table-column prop="钢材事业部-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="钢材事业部-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                        <el-table-column label="农产事业部">
                            <el-table-column prop="农产事业部-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="农产事业部-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                        <el-table-column label="原料事业部">
                            <el-table-column prop="原料事业部-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="原料事业部-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                        <el-table-column label="总经理二室">
                            <el-table-column prop="总经理二室-交易员" label="交易员"></el-table-column>
                            <el-table-column prop="总经理二室-商务员" label="商务员"></el-table-column>
                        </el-table-column>
                    </el-table-column>
                    <el-table-column label="总经理一室">
                        <el-table-column prop="总经理一室-资金助理" label="资金助理"></el-table-column>
                    </el-table-column>
                    <el-table-column label="杉贸本部">
                        <el-table-column prop="杉贸本部-资金助理" label="资金助理"></el-table-column>
                    </el-table-column>
                    <el-table-column label="交易部">
                        <el-table-column prop="交易部-交易员" label="交易员"></el-table-column>
                    </el-table-column>
                </el-table-column>
                <el-table-column label="智维">
                    <el-table-column label="钢材事业部">
                        <el-table-column prop="智维钢材-交易员" label="交易员"></el-table-column>
                        <el-table-column prop="智维钢材-商务员" label="商务员"></el-table-column>
                    </el-table-column>
                    <el-table-column label="智维本部">
                        <el-table-column prop="智维本部-资金助理" label="资金助理"></el-table-column>
                    </el-table-column>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script>
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { getOneConfirmData } from "@/api/internalSystem/fund/fund_budget_data"; //  此处请自行替换地址
import {
    getConfirmSetting,
    upConfirmSetting
} from "@/api/internalSystem/fund/fund_confirm_date";  //  此处请自行替换地址

export default {
    name: "BudgetIndex",
    mixins: [infoList],
    data() {
        return {
            setInfo: {
                chose_date: new Date(),
                is_confirm:true,
            },
            tableData: [
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
        //数据确认提交
        //根据确认数据创建时间取值
        async getConfirm() {
            let confirm_value = {
                chose_date: formatTimeToStr(this.setInfo.chose_date, "yyyy-MM-dd"),
            }
                let setting_value = {
                confirm_date: formatTimeToStr(this.setInfo.chose_date, "yyyy-MM-dd"),
            }
            let res = await getOneConfirmData(confirm_value)
            let res_setting = await getConfirmSetting(setting_value) 
          
            if (res.code == 0 && res_setting.code==0) {
                this.$message({
                    type: "success",
                    message: "获取成功"

                })  
                console.log(res)
                this.tableData = this.listConfirm(res.data.list)
                if(res_setting.data.list[0]['is_ok']==0){
                    this.setInfo.is_confirm = true
                }else{
                     this.setInfo.is_confirm = false
                }
                
            } else {
                this.$message({
                    type: "error",
                    message: "获取失败"
                })
            }
        },
        //修改确认状态
        async upConfirmSettings(){
            var is_oks
            if(this.setInfo.is_confirm){
                is_oks = 0
            }else{
                is_oks=1
            }
            let setting_value = {
                is_ok:is_oks,
                confirm_date: formatTimeToStr(this.setInfo.chose_date, "yyyy-MM-dd"),
            }
            let res = await upConfirmSetting(setting_value)
            console.log(res)
        },

        listConfirm(data) {
            const listConfirmData = [{}]
            listConfirmData[0]['date'] = formatTimeToStr(this.setInfo.chose_date, "yyyy-MM-dd")
            for (var i in data) {
                listConfirmData[0][data[i]['department_code'] + "-" + data[i]['create_role']] = data[i]['is_confirm']
            }
            return listConfirmData
        },
        cellStyle() {
            return "text-align:center";
        },
        rowClass() {
            return "text-align:center";
        },
    },
    created() {
        this.getConfirm()
    },
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
</style>
