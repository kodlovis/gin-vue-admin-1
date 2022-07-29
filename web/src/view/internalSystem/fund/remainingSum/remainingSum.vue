<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增余额/可质押仓单</el-button>
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
      <el-table-column label="填写日期" width="180">
        <template slot-scope="scope">{{ scope.row.upDate | formatDate }}</template>
      </el-table-column>

      <el-table-column label="创建人" prop="createUser" width="120"></el-table-column>

      <el-table-column label="类型" prop="type" width="120"></el-table-column>

      <el-table-column label="余额/可质押" prop="inValue" width="120"></el-table-column>

      <el-table-column label="平台" prop="company" width="120"></el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateRemainingSum(scope.row)" size="small" type="primary"
            icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }" :total="total" @current-change="handleCurrentChange"
      @size-change="handleSizeChange" layout="total, sizes, prev, pager, next, jumper"></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" width="450px" title="资金余额/可质押录入">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="填写日期:">
          <el-date-picker type="date" placeholder="选择填报日期" v-model="formData.upDate" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="formData.type"   placeholder="请选择类型">
            <el-option v-for="item in style_options" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="余额:">
          <el-input style="width:220px" v-model.number="formData.inValue" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="平台:">
          <el-input style="width:220px" v-model="formData.company" :disabled="true">
          </el-input>

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
  createRemainingSum,
  deleteRemainingSum,
  deleteRemainingSumByIds,
  updateRemainingSum,
  findRemainingSum,
  getRemainingSumList
} from "@/api/internalSystem/fund/remainingSum";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "RemainingSum",
  mixins: [infoList],
  data() {
    return {
      style_options: [{
        value: '资金余额',
        label: '资金余额'
      }, {
        value: '可质押余额',
        label: '可质押余额'
      }],
      listApi: getRemainingSumList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        inValue: 0,
        company: "杉贸",
        createUser: JSON.parse( localStorage.getItem("vuex"))['user']['userInfo']['nickName']
      }
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
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteRemainingSum(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteRemainingSumByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateRemainingSum(row) {
      const res = await findRemainingSum({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reremainingSum;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        inValue: 0,

      };
    },
    async deleteRemainingSum(row) {
      const res = await deleteRemainingSum({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
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
          res = await createRemainingSum(this.formData);
          break;
        case "update":
          res = await updateRemainingSum(this.formData);
          break;
        default:
          res = await createRemainingSum(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
            this.formData={
        inValue: 0,
        company: "杉贸",
        createUser: JSON.parse( localStorage.getItem("vuex"))['user']['userInfo']['nickName']

      }
      this.dialogFormVisible = true;
    },
    loadComments() {
      var list = JSON.parse( localStorage.getItem("vuex"));
      console.log(list)
      this.formData.company = list['user']['userInfo']['nickName'];
    },
  },
  async created() {
    await this.getTableData();
  }
};
</script>

<style>
.el-input {
  width: 220px
}
</style>
