<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="日期">
          <el-input placeholder="搜索条件" v-model="searchInfo.time"></el-input>
        </el-form-item>    
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.department"></el-input>
        </el-form-item>    
        <el-form-item label="交易类别">
          <el-input placeholder="搜索条件" v-model="searchInfo.tradeType"></el-input>
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
    <el-table-column type="selection" width="40"></el-table-column>
    
    <el-table-column label="日期" prop="time" width="170"></el-table-column> 
    
    <el-table-column label="部门" prop="department" width="120"></el-table-column> 
    
    <el-table-column label="交易类别" prop="tradeType" width="100"></el-table-column> 
    <el-table-column label="风险类别" prop="riskType" width="100"></el-table-column> 
    
    <el-table-column label="多头合约" prop="longInstrument" width="100"></el-table-column> 
    
    <el-table-column label="多头开仓均价（基差）" prop="longBasis" width="110"></el-table-column> 
    
    <el-table-column label="多头市价（基差）" prop="longMarket" width="80"></el-table-column> 
    <el-table-column label="多头既定价" prop="longPrice" width="90"></el-table-column> 
    
    <el-table-column label="多头账户" prop="longAccount" width="120"></el-table-column> 
    
    <el-table-column label="多头持仓吨数" prop="longPosition" width="110"></el-table-column> 
    
    <el-table-column label="多头持仓备注" prop="longComment" width="120"></el-table-column> 
    
    <el-table-column label="空头合约" prop="shortInstrument" width="100"></el-table-column> 
    
    <el-table-column label="空头开仓均价（基差）" prop="shortBasis" width="110"></el-table-column> 
    
    <el-table-column label="空头市价（基差）" prop="shortMarket" width="80"></el-table-column> 
    
    <el-table-column label="空头既定价" prop="shortPrice" width="90"></el-table-column> 
    <el-table-column label="空头账户" prop="shortAccount" width="120"></el-table-column> 
    
    <el-table-column label="空头持仓吨数" prop="shortPosition" width="110"></el-table-column> 
    
    <el-table-column label="空头持仓备注" prop="shortComment" width="120"></el-table-column> 
    
      <el-table-column label="操作" width="180">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateFuturePositionSizeDetail(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="160px">
         <el-form-item label="日期:">
            <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="部门:">
            <el-input v-model="formData.department" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="交易类别:">
            <el-input v-model="formData.tradeType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="风险类别:">
            <el-input v-model="formData.riskType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="多头合约:">
            <el-input v-model="formData.longInstrument" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="多头开仓均价（基差）:">
            <el-input-number v-model="formData.longBasis" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
         <el-form-item label="多头市价（基差）:">
            <el-input-number v-model="formData.longMarket" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="多头既定价:">
            <el-input-number v-model="formData.longPrice" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="多头账户:">
            <el-input v-model="formData.longAccount" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="多头持仓(吨):">
            <el-input-number v-model="formData.longPosition" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="多头持仓备注:">
            <el-input v-model="formData.longComment" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="空头合约:">
            <el-input v-model="formData.shortInstrument" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="空头开仓均价（基差）:">
            <el-input-number v-model="formData.shortBasis" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="空头市价（基差）:">
            <el-input-number v-model="formData.shortMarket" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
         <el-form-item label="空头既定价:">
            <el-input-number v-model="formData.shortPrice" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="空头账户:">
            <el-input v-model="formData.shortAccount" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="空头持仓(吨):">
            <el-input-number v-model="formData.shortPosition" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="空头持仓备注:">
            <el-input v-model="formData.shortComment" clearable placeholder="请输入" ></el-input>
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
    createFuturePositionSizeDetail,
    deleteFuturePositionSizeDetail,
    deleteFuturePositionSizeDetailByIds,
    updateFuturePositionSizeDetail,
    findFuturePositionSizeDetail,
    getFuturePositionSizeDetailList
} from "@/api/internalSystem/scaleRights/futurePositionSizeDetail";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "FuturePositionSizeDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getFuturePositionSizeDetailList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:"",
            department:"",
            tradeType:"",
            riskType:"",
            longInstrument:"",
            longBasis:"",
            longMarket:"",
            shortPrice:"",
            longPrice:"",
            longAccount:"",
            longPosition:"",
            longComment:"",
            shortInstrument:"",
            shortBasis:"",
            shortMarket:"",
            shortAccount:"",
            shortPosition:"",
            shortComment:"",
            
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
           this.deleteFuturePositionSizeDetail(row);
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
        const res = await deleteFuturePositionSizeDetailByIds({ ids })
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
    async updateFuturePositionSizeDetail(row) {
      const res = await findFuturePositionSizeDetail({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.refuturePositionSizeDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
            time:"",
            department:"",
            tradeType:"",
            riskType:"",
            longInstrument:"",
            longBasis:"",
            longMarket:"",
            shortPrice:"",
            longPrice:"",
            longAccount:"",
            longPosition:"",
            longComment:"",
            shortInstrument:"",
            shortBasis:"",
            shortMarket:"",
            shortAccount:"",
            shortPosition:"",
            shortComment:"",
          
      };
    },
    async deleteFuturePositionSizeDetail(row) {
      const res = await deleteFuturePositionSizeDetail({ ID: row.ID });
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
          res = await createFuturePositionSizeDetail(this.formData);
          break;
        case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateFuturePositionSizeDetail(this.formData);
              if (res.code == 0) {
                this.closeDialog();
                this.$message({
                  type: "success",
                  message: "修改成功!"
                });
                this.getTableData();
              }
              }).catch(() => {
              this.$message({
                type: "info",
                message: "已取消修改"
              });
            });
          break;
        default:
          res = await createFuturePositionSizeDetail(this.formData);
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
