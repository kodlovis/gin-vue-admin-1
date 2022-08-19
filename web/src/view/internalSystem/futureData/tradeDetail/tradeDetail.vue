<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="时间">
          <el-input placeholder="搜索条件" v-model="searchInfo.time"></el-input>
        </el-form-item>    
        <el-form-item label="账号">
          <el-input placeholder="搜索条件" v-model="searchInfo.accountId"></el-input>
        </el-form-item>    
        <el-form-item label="合约">
          <el-input placeholder="搜索条件" v-model="searchInfo.instrument"></el-input>
        </el-form-item>                        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增平仓明细</el-button>
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
    
    <el-table-column label="时间" prop="time" width="170"></el-table-column> 
    
    <el-table-column label="账号" prop="accountId" width="100"></el-table-column> 
    
    <el-table-column label="合约" prop="instrument" width="100"></el-table-column> 
    
    <el-table-column label="交易时间" prop="tradeTime" width="120"></el-table-column> 
    
    <el-table-column label="交易ID" prop="tradeId" width="100"></el-table-column> 
    
    <el-table-column label="方向" prop="direction" width="80"></el-table-column> 
    
    <el-table-column label="开/平" prop="offsetFlag" width="80"></el-table-column> 
    
    <el-table-column label="投机/套保" prop="hedgeFlag" width="100"></el-table-column> 
    
    <el-table-column label="成交价" prop="tradePrice" width="100"></el-table-column> 
    
    <el-table-column label="手数" prop="volume" width="100"></el-table-column> 
    
    <el-table-column label="成交额" prop="amount" width="100"></el-table-column> 
    
    <!-- <el-table-column label="tradeFee字段" prop="tradeFee" width="120"></el-table-column>  -->
    
    <el-table-column label="平仓盈亏" prop="profitByTrade" width="100"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateTradeDetail(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="编辑">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="time字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="账号:">
            <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="合约:">
            <el-input v-model="formData.instrument" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="交易时间:">
            <el-input v-model="formData.tradeTime" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="交易ID:">
            <el-input v-model="formData.tradeId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="方向:">
            <el-input v-model="formData.direction" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="开/平:">
            <el-input v-model="formData.offsetFlag" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="投机/套保:">
            <el-input v-model="formData.hedgeFlag" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="成交价:">
              <el-input-number v-model="formData.tradePrice" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="手数:"><el-input v-model.number="formData.volume" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="成交额:">
              <el-input-number v-model="formData.amount" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <!-- <el-form-item label="tradeFee字段:">
              <el-input-number v-model="formData.tradeFee" :precision="2" clearable></el-input-number>
       </el-form-item> -->
       
         <el-form-item label="平仓盈亏:">
              <el-input-number v-model="formData.profitByTrade" :precision="2" clearable></el-input-number>
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
    createTradeDetail,
    deleteTradeDetail,
    deleteTradeDetailByIds,
    updateTradeDetail,
    findTradeDetail,
    getTradeDetailList
} from "@/api/internalSystem/futureData/tradeDetail";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "TradeDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getTradeDetailList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:new Date(),
            accountId:"",
            instrument:"",
            tradeTime:"",
            tradeId:"",
            direction:"",
            offsetFlag:"",
            hedgeFlag:"",
            tradePrice:0,
            volume:0,
            amount:0,
            tradeFee:0,
            profitByTrade:0,
            
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
           this.deleteTradeDetail(row);
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
            ids.push(item.id)
          })
        const res = await deleteTradeDetailByIds({ ids })
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
    async updateTradeDetail(row) {
      const res = await findTradeDetail({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.retradeDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:new Date(),
          accountId:"",
          instrument:"",
          tradeTime:"",
          tradeId:"",
          direction:"",
          offsetFlag:"",
          hedgeFlag:"",
          tradePrice:0,
          volume:0,
          amount:0,
          tradeFee:0,
          profitByTrade:0,
          
      };
    },
    async deleteTradeDetail(row) {
      const res = await deleteTradeDetail({ ID: row.id });
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
          res = await createTradeDetail(this.formData);
          break;
        case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateTradeDetail(this.formData);
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
          res = await createTradeDetail(this.formData);
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
