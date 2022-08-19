<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="日期">
          <el-input placeholder="搜索条件" v-model="searchInfo.time"></el-input>
        </el-form-item>      
        <el-form-item label="期货公司">
          <el-input placeholder="搜索条件" v-model="searchInfo.brokerId"></el-input>
        </el-form-item>    
        <el-form-item label="品种">
          <el-input placeholder="搜索条件" v-model="searchInfo.productCode"></el-input>
        </el-form-item>                      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
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
        <el-form-item>
              <el-upload
                :action="`${path}/excel/importExcel`"
                :headers="{'x-token':token}"
                :on-success="loadUs001CzceBrokerPositionExcel"
                :show-file-list="false"
                :limit="2"
                multiple
              >
                <el-button size="small" type="primary" icon="el-icon-upload2">Excel导入</el-button>
              </el-upload>
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
    
    <el-table-column label="日期" prop="time" width="200"><template slot-scope="scope">{{scope.row.time|formatDate}}</template></el-table-column> 
    
    
    <el-table-column label="期货公司" prop="brokerId" width="120"></el-table-column> 
    
    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="个人客户日均持仓" prop="customerDailyPosition" width="140"></el-table-column> 
    
    <el-table-column label="机构客户日均持仓" prop="institutionalDailyPosition" width="140"></el-table-column> 
    
    <el-table-column label="特法客户日均持仓" prop="specialLawCustomerDailyPosition" width="140"></el-table-column> 
    
    <el-table-column label="法人客户日均持仓" prop="corporateCustomerDailyPosition" width="220"></el-table-column> 
    
    <el-table-column label="今日日均持仓" prop="dailyAveragePosition" width="120"></el-table-column> 
    
    
    <el-table-column label="重点合约当月日均持仓" prop="keyContractMonthlyAveragePosition" width="180"></el-table-column> 
    
    <el-table-column label="本月截至今日上缴手续费" prop="theHandlingFeeDueTodayThisMonth" width="200"></el-table-column> 
    
    <el-table-column label="法人持仓占比乘数" prop="corporatePositionMultiplier" width="140"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUs001CzceBrokerPositionDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
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
      <el-form :model="formData" label-position="right" label-width="250px">
         <el-form-item label="日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="期货公司:">
            <el-input v-model="formData.brokerId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="个人客户日均持仓:">
              <el-input-number v-model="formData.customerDailyPosition" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="机构客户日均持仓:">
              <el-input-number v-model="formData.institutionalDailyPosition" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="特法客户日均持仓:">
              <el-input-number v-model="formData.specialLawCustomerDailyPosition" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="法人客户日均持仓:">
              <el-input-number v-model="formData.corporateCustomerDailyPosition" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="今日日均持仓:">
              <el-input-number v-model="formData.dailyAveragePosition" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="重点合约当月日均持仓:">
              <el-input-number v-model="formData.keyContractMonthlyAveragePosition" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="本月截至今日上缴手续费:">
              <el-input-number v-model="formData.theHandlingFeeDueTodayThisMonth" :precision="2" clearable></el-input-number></el-form-item>
       
         <el-form-item label="法人持仓占比乘数:">
              <el-input-number v-model="formData.corporatePositionMultiplier" :precision="2" clearable></el-input-number></el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import { mapGetters } from 'vuex';
import {
    createUs001CzceBrokerPositionDaily,
    deleteUs001CzceBrokerPositionDaily,
    deleteUs001CzceBrokerPositionDailyByIds,
    updateUs001CzceBrokerPositionDaily,
    findUs001CzceBrokerPositionDaily,
    getUs001CzceBrokerPositionDailyList,
    loadUs001CzceBrokerPositionExcelData
} from "@/api/internalSystem/positionReward/czce/brokerPositionDaily";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Us001CzceBrokerPositionDaily",
  mixins: [infoList],
  data() {
    return {
      path: path,
      listApi: getUs001CzceBrokerPositionDailyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:new Date(),
            exchangeId:"",
            brokerId:"",
            productCode:"",
            customerDailyPosition:0,
            
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
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
      async loadUs001CzceBrokerPositionExcel() {
        const res = await loadUs001CzceBrokerPositionExcelData();
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '上传成功'
          })
          this.getTableData();
        }
      },
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
           this.deleteUs001CzceBrokerPositionDaily(row);
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
        const res = await deleteUs001CzceBrokerPositionDailyByIds({ ids })
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
    async updateUs001CzceBrokerPositionDaily(row) {
      const res = await findUs001CzceBrokerPositionDaily({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reus001CzceBrokerPositionDaily;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:new Date(),
          exchangeId:"",
          brokerId:"",
          productCode:"",
          customerDailyPosition:0,
          
      };
    },
    async deleteUs001CzceBrokerPositionDaily(row) {
      const res = await deleteUs001CzceBrokerPositionDaily({ ID: row.ID });
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
          res = await createUs001CzceBrokerPositionDaily(this.formData);
          break;
        case "update":
          res = await updateUs001CzceBrokerPositionDaily(this.formData);
          break;
        default:
          res = await createUs001CzceBrokerPositionDaily(this.formData);
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
