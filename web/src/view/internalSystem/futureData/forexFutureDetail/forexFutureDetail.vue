<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="时间">
          <div class="block">
            <el-date-picker
              v-model="searchInfo.time"
              type="datetime"
              placeholder="选择日期时间"
              default-time="12:00:00">
            </el-date-picker>
          </div>
        </el-form-item>  
        <el-form-item label="品种">
          <el-input placeholder="搜索条件" v-model="searchInfo.productName"></el-input>
        </el-form-item>    
        <el-form-item label="账户">
          <el-input placeholder="搜索条件" v-model="searchInfo.accountId"></el-input>
        </el-form-item>    
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.departmentName"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增数据</el-button>
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
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    
    <el-table-column label="时间" prop="time" width="220"></el-table-column> 
    
    <el-table-column label="品种" prop="productName" width="120"></el-table-column> 
    
    <el-table-column label="账户" prop="accountId" width="120"></el-table-column> 
    
    <el-table-column label="浮动盈亏" prop="profitByFloat" width="120"></el-table-column> 
    
    <el-table-column label="平仓盈亏" prop="profitByTrade" width="120"></el-table-column> 
    
    <el-table-column label="手续费" prop="tradeFee" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="departmentName" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateForexFutureDetail(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作" v-dialogDrag>
      <el-form :model="formData" label-position="right" label-width="100px">
         <el-form-item label="时间:">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
          </div>
       </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="账号:">
            <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="浮动盈亏:">
            <el-input-number v-model="formData.profitByFloat" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="平仓盈亏:">
            <el-input-number v-model="formData.profitByTrade" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="手续费:">
            <el-input-number v-model="formData.tradeFee" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="部门:">
            <el-input v-model="formData.departmentName" clearable placeholder="请输入" ></el-input>
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
    createForexFutureDetail,
    deleteForexFutureDetail,
    deleteForexFutureDetailByIds,
    updateForexFutureDetail,
    findForexFutureDetail,
    getForexFutureDetailList
} from "@/api/internalSystem/futureData/forexFutureDetail";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "ForexFutureDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getForexFutureDetailList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:"",
            productName:"",
            accountId:"",
            profitByFloat:"",
            profitByTrade:"",
            tradeFee:"",
            departmentName:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
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
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteForexFutureDetail(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
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
        const res = await deleteForexFutureDetailByIds({ ids })
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
    async updateForexFutureDetail(row) {
      const res = await findForexFutureDetail({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reforexFutureDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:"",
          productName:"",
          accountId:"",
          profitByFloat:"",
          profitByTrade:"",
          tradeFee:"",
          departmentName:"",
          
      };
    },
    async deleteForexFutureDetail(row) {
      const res = await deleteForexFutureDetail({ ID: row.id });
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
          res = await createForexFutureDetail(this.formData);
          break;
        case "update":
          res = await updateForexFutureDetail(this.formData);
          break;
        default:
          res = await createForexFutureDetail(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
