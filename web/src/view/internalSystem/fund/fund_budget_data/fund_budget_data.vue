<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增录入数据</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @selection-change="handleSelectionChange" border ref="multipleTable" stripe
      style="width: 100%" tooltip-effect="dark">
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="创建日期" width="180">
        <template slot-scope="scope">{{
            scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>
      <el-table-column label="预测日期" width="180">
        <template slot-scope="scope">{{
            scope.row.in_data | formatDate
        }}</template>
      </el-table-column>


      <el-table-column  label="录入类型" width="230">
          <template slot-scope="scope">{{
             all_type[scope.row.type_code] 
        }}</template>
      </el-table-column>

      <el-table-column label="录入数据" prop="in_value" width="120"></el-table-column>

      <el-table-column label="数据备注" prop="comment" width="120"></el-table-column>

      <el-table-column label="币种" prop="currency" width="120"></el-table-column>

      <el-table-column label="部门" prop="department_code" width="120"></el-table-column>

      <el-table-column label="录入角色" prop="create_role" width="120"></el-table-column>

      <el-table-column label="录入人员" prop="create_user" width="120"></el-table-column>


      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateBudgetData(scope.row)" size="small" type="primary"
            icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }" :total="total" @current-change="handleCurrentChange"
      @size-change="handleSizeChange" layout="total, sizes, prev, pager, next, jumper"></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="录入日期:">
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.in_data" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="录入类型:">
          <el-input v-model="formData.type_code" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="录入数据:">
          <el-input v-model.number="formData.in_value" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="数据备注:">
          <el-input v-model="formData.comment" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="币种:">
          <el-input v-model="formData.currency" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="部门编号:">
          <el-input v-model="formData.department_code" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="录入人员:">
          <el-input v-model="formData.create_user" clearable placeholder="请输入"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createBudgetData,
  deleteBudgetData,
  deleteBudgetDataByIds,
  updateBudgetData,
  findBudgetData,
  getBudgetDataList,
} from "@/api/internalSystem/fund/fund_budget_data"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "BudgetData",
  mixins: [infoList],
  data() {
    return {
      listApi: getBudgetDataList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        in_data: new Date(),
        type_code: "",
        in_value: 0,
        comment: "",
        currency: "",
        department_code: "",
        create_user: "",
        create_role:"",
      },
      all_type: {
        cghk: "支出-现货采购-采购货款",
        cghkbz: "",
        mhjzy: "支出-现货采购-卖货解质押",
        mhjzybz: "",
        outllyw: "支出-现货采购-流量业务",
        outllywbz: "",
        xzjtrzjk: "支出-现货采购-协助集团融资借款",
        xzjtrzjkbz: "",
        dqfh: "支出-现货采购-到期付汇",
        dqfhbz: "",
        kzbzj: "支出-现货采购-开证保证金",
        kzbzjbz: "",
        sj: "支出-现货采购-税金",
        sjbz: "",
        zcxs: "收入-现货销售-正常销售",
        zcxsbz: "",
        mhzy: "收入-现货销售-买货质押",
        mhzybz: "",
        inllyw: "收入-现货销售-流量业务",
        inllywbz: "",
        xzjtjkhk: "收入-现货销售-协助集团融资借款回款",
        xzjtjkhkbz: "",
        kzbzjth: "收入-现货销售-开保证金及收益退回",
        kzbzjthbz: "",
        kc: "支出-期货入金-开仓",
        kcbz: "",
        yhjzy: "支出-期货入金-银行解质押",
        yhjzybz: "",
        mjgjzy: "支出-期货入金-卖交割解质押",
        mjgjzybz: "",
        outmjg: "支出-期货入金-买交割",
        outmjgbz: "",
        tbrj: "支出-期货入金-提保入金",
        tbrjbz: "",
        pc: "收入-期货出金-平仓",
        pcbz: "",
        yhzy: "收入-期货出金-银行质押",
        yhzybz: "",
        mjgzy: "收入-期货出金-买交割质押",
        mjgzybz: "",
        inmjg: "收入-期货出金-卖交割",
        inmjgbz: "",
        sfbzj: "收入-期货出金-释放保证金",
        sfbzjbz: "",
        jjfz: "支出-工资奖金和房租",
        jjfzbz: "",
        ywck: "支出-业务存款",
        ywckbz: "",
        ywhk: "支出-业务还款",
        ywhkbz: "",
        cwsxf: "支出-财务手续费",
        cwsxfbz: "",
        rzhk: "支出-融资还款",
        rzhkbz: "",
        jkcc: "支出-借款拆出",
        jkccbz: "",
        outcjhk: "支出-拆借还款",
        outcjhkbz: "",
        rzck: "支出-融资存款",
        rzckbz: "",
        hbck: "支出-回报存款",
        hbckbz: "",
        lxsr: "收入-利息收入",
        lxsrbz: "",
        zxf: "收入-咨询费",
        zxfbz: "",
        bzjjsyth: "收入-保证金及收益退回",
        bzjjsythbz: "",
        jkcr: "收入-借款拆入",
        jkcrbz: "",
        injkhk: "收入-借款还款",
        injkhkbz: "",
        rztx: "收入-融资贴现",
        rztxbz: "",
        hbtx: "收入-回报贴现",
        hbtxbz: "",
      },
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
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    deleteRow(row) {
      this.$confirm("确定要删除吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.deleteBudgetData(row);
      });
    },
    async onDelete() {
      const ids = [];
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: "warning",
          message: "请选择要删除的数据",
        });
        return;
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID);
        });
      const res = await deleteBudgetDataByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async updateBudgetData(row) {
      const res = await findBudgetData({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rebudget_data;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        in_data: new Date(),
        type_code: "",
        in_value: 0,
        comment: "",
        currency: "",
        department_code: "",
        create_user: "",
      };
    },
    async deleteBudgetData(row) {
      const res = await deleteBudgetData({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createBudgetData(this.formData);
          break;
        case "update":
          res = await updateBudgetData(this.formData);
          break;
        default:
          res = await createBudgetData(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
  async created() {
    await this.getTableData();
  },
};
</script>

<style>
</style>
