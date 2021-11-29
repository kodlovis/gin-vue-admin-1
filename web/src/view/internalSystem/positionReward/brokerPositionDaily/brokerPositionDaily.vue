<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="交易日期">
          <div class="block">
            <el-date-picker
              v-model="searchInfo.tradingDate"
              type="datetime"
              placeholder="选择日期时间"
              default-time="12:00:00">
            </el-date-picker>
          </div>
        </el-form-item>
        <el-form-item label="期货公司">
           <el-select v-model="searchInfo.brokerId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.brokerId"
              :label="`${item.comment}(${item.brokerId})`"
              :value="item.brokerId"
              v-for="item in accountInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>                 
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('create')" type="primary">新增持仓</el-button>
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
                :on-success="loadBrokerPositionExcel"
                :show-file-list="false"
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
    
    <el-table-column label="交易日期" prop="tradingDate" width="220"></el-table-column> 
    
    <el-table-column label="期货公司" prop="accountInfo.comment" width="120"></el-table-column> 
    
    <el-table-column label="品种" prop="productInfo.productName" width="120"></el-table-column> 
    
    <el-table-column label="总上缴手续费" prop="totalTradingFee" width="120"></el-table-column> 
    
    <el-table-column label="量化客户手续费" prop="specialTradingFee" width="120"></el-table-column> 
    
    <el-table-column label="日均持仓" prop="averageDailyPosition" width="120"></el-table-column> 
    
    <el-table-column label="最大可操作量" prop="maximumToOpen" width="120"></el-table-column> 
    
    <el-table-column label="最新持仓" prop="currentPosition" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateBrokerPositionDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogTitle">
      <el-form :model="formData" label-position="right" label-width="100px"  ref="prForm">
         <el-form-item label="交易日期:">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.tradingDate" clearable default-time="12:00:00"></el-date-picker>
          </div>
      </el-form-item>
       
         <el-form-item label="期货公司:">
           <el-select v-model="formData.brokerId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.brokerId"
              :label="`${item.comment}(${item.brokerId})`"
              :value="item.brokerId"
              v-for="item in accountInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
       
         <el-form-item label="品种:">
           <el-select v-model="formData.productCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productCode"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productCode"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
       
         <el-form-item label="交易手续费:">
              <el-input-number v-model="formData.totalTradingFee" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="特殊交易费用:">
              <el-input-number v-model="formData.specialTradingFee" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="日均持仓:">
              <el-input-number v-model="formData.averageDailyPosition" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="最大可操作量:">
              <el-input-number v-model="formData.maximumToOpen" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="最新持仓:">
              <el-input-number v-model="formData.currentPosition" :precision="2" clearable></el-input-number>
       </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary" :disabled="isDisable">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import { mapGetters } from 'vuex';
import {
    createBrokerPositionDaily,
    deleteBrokerPositionDaily,
    deleteBrokerPositionDailyByIds,
    updateBrokerPositionDaily,
    findBrokerPositionDaily,
    getBrokerPositionDailyList,
    loadBrokerPositionExcelData
} from "@/api/internalSystem/positionReward/brokerPositionDaily";  //  此处请自行替换地址
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
import {getAccountInfoList} from "@/api/internalSystem/accountInfo"; 
import { formatTimeToStr } from "@/utils/date";
import { getDict } from "@/utils/dictionary";
import infoList from "@/mixins/infoList";
export default {
  name: "BrokerPositionDaily",
  mixins: [infoList],
  data() {
    return {
      path: path,
      listApi: getBrokerPositionDailyList,
      dialogFormVisible: false,
      type: "",
      isDisable:false,
      deleteVisible: false,
      accountInfoOptions:[],
      brokerDictList:[],
      productInfoOptions:"",
      dialogTitle:"",
      multipleSelection: [],formData: {
            tradingDate:"",
            brokerId:"",
            productCode:"",
            totalTradingFee:0,
            specialTradingFee:0,
            averageDailyPosition:0,
            maximumToOpen:0,
            currentPosition:0,
            
      },
      accountInfoData:{
           accountId:"",
           comment:"",
      }, 
      productInfo:{
           productInfoCode:"",
           productInfoName:"",
      }, 
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
      async loadBrokerPositionExcel() {
        const res = await  loadBrokerPositionExcelData();
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
      brokerfilterDict(brokerId){
        const re = this.brokerDictList.filter(item=>{
          return item.value == brokerId
        })[0]
        if(re){
          return re.label
          }
        else{
          return""
          }
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
           this.deleteBrokerPositionDaily(row);
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
        const res = await deleteBrokerPositionDailyByIds({ ids })
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
    async updateBrokerPositionDaily(row) {
      const res = await findBrokerPositionDaily({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reBrokerPositionDaily;
        this.openDialog("update");
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          tradingDate:"",
          brokerId:"",
          productCode:"",
          totalTradingFee:0,
          specialTradingFee:0,
          averageDailyPosition:0,
          maximumToOpen:0,
          currentPosition:0,
          
      };
    },
    async deleteBrokerPositionDaily(row) {
      const res = await deleteBrokerPositionDaily({ ID: row.id });
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
      this.isDisable=true;
      let res;
      switch (this.type) {
        case "create":
          res = await createBrokerPositionDaily(this.formData);
          break;
        case "update":
          res = await updateBrokerPositionDaily(this.formData);
          break;
        default:
          res = await createBrokerPositionDaily(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.isDisable=false
        this.getTableData();
      }
    },
    openDialog(type) {
      switch (type) {
        case "create":
          this.dialogTitle = "新增持仓";
          break;
        case "update":
          this.dialogTitle = "编辑持仓";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    setProductInfoOptions(productInfoData) {
        this.productInfoOptions = [];
        this.ids = [];
        this.setProductInfoOptionsData(productInfoData, this.productInfoOptions ,this.ids);
      },
      setProductInfoOptionsData(ProductInfoData, optionsData ,ids) {
        ProductInfoData &&
          ProductInfoData.map(item => {
              const option = {
                productCode: item.productCode,
                productName: item.productName
              };
              optionsData.push(option);
              const idOption = {
                productId: item.productId,
              };
              ids.push(idOption)
          });
      },
    setAccountInfoOptions(accountInfoData) {
        this.accountInfoOptions = [];
        this.ids = [];
        this.setAccountInfoOptionsData(accountInfoData, this.accountInfoOptions ,this.ids);
      },
      setAccountInfoOptionsData(AccountInfoData, optionsData ,ids) {
        AccountInfoData &&
          AccountInfoData.map(item => {
            if(item.type=='4'){
              const option = {
                brokerId: item.accountId,
                comment: item.comment
              };
              optionsData.push(option);
              const idOption = {
                brokerId: item.accountId,
              };
              ids.push(idOption)}
          });
      },
  },
  async created() {
    //加载品种信息
    const productInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setProductInfoOptions(productInfo.data.list);
    await this.getTableData();
    //获取期货公司字典
    const broker = await getDict("broker");
    broker.map(item=>item.value)
    this.brokerDictList = broker
    //加载期货账户信息
    const accountInfo = await getAccountInfoList({ page: 1, pageSize: 999 });
    this.setAccountInfoOptions(accountInfo.data.list);
  
  
}
};
</script>

<style>
</style>
