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
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('create')" type="primary">新增</el-button>
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
                :on-success="loadAccountPositionExcel"
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
    
    <el-table-column label="账户" prop="accountId" width="120"></el-table-column> 
    
    <el-table-column label="合约" prop="instrument" width="120"></el-table-column> 
    
    <el-table-column label="方向" prop="direction" width="120">
      <template slot-scope="scope">
          <span>{{directionfilterDict(scope.row.direction)}}</span>
      </template></el-table-column> 
    
    <el-table-column label="投机/套保" prop="hedgeFlag" width="120">
      <template slot-scope="scope">
          <span>{{hedgeFlagfilterDict(scope.row.hedgeFlag)}}</span>
      </template></el-table-column> 
    
    <el-table-column label="数量" prop="amount" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAccountPositionDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="80px" :rules="rules" ref="prForm">
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
       
         <el-form-item label="账户:">
            <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="合约:">
            <el-input v-model="formData.instrument" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="方向:">
          <el-select v-model="formData.direction" placeholder="请选择">
            <el-option
              v-for="item in directionDictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>  </el-form-item>
       
         <el-form-item label="套保:">
          <el-select v-model="formData.hedgeFlag" placeholder="请选择">
            <el-option
              v-for="item in hedgeFlagDictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>  </el-form-item>
       
         <el-form-item label="数量:">
            <el-input-number v-model="formData.amount" clearable placeholder="请输入" ></el-input-number></el-form-item>
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
    createAccountPositionDaily,
    deleteAccountPositionDaily,
    deleteAccountPositionDailyByIds,
    updateAccountPositionDaily,
    findAccountPositionDaily,
    getAccountPositionDailyList,
    loadAccountPositionExcelData
} from "@/api/internalSystem/positionReward/accountPositionDaily";  //  此处请自行替换地址
import { getDict } from "@/utils/dictionary";
import {getAccountInfoList} from "@/api/internalSystem/accountInfo"; 
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AccountPositionDaily",
  mixins: [infoList],
  data() {
    return {
      path: path,
      listApi: getAccountPositionDailyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      isDisable:false,
      brokerDictList:[],
      hedgeFlagDictList:[],
      directionDictList:[],
      multipleSelection: [],formData: {
            tradingDate:"",
            brokerId:"",
            accountId:"",
            instrument:"",
            direction:"",
            hedgeFlag:"",
            amount:"",
      },
      accountInfoData:{
           accountId:"",
           comment:"",
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
      async loadAccountPositionExcel() {
        const res = await  loadAccountPositionExcelData();
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '上传成功'
          })
          this.getTableData();
        }
      },
    openDialog(type) {
      switch (type) {
        case "create":
          this.dialogTitle = "新增持仓信息";
          break;
        case "update":
          this.dialogTitle = "编辑持仓信息";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        this.getTableData()
      },
      //字典
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
      //字典
      directionfilterDict(direction){
        const re = this.directionDictList.filter(item=>{
          return item.value == direction
        })[0]
        if(re){
          return re.label
          }
        else{
          return""
          }
      },
      //字典
      hedgeFlagfilterDict(hedgeFlag){
        const re = this.hedgeFlagDictList.filter(item=>{
          return item.value == hedgeFlag
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
           this.deleteAccountPositionDaily(row);
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
        const res = await deleteAccountPositionDailyByIds({ ids })
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
    async updateAccountPositionDaily(row) {
      const res = await findAccountPositionDaily({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reAccountPositionDaily;
        this.openDialog("update");
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          tradingDate:"",
          brokerId:"",
          accountId:"",
          instrument:"",
          direction:"",
          hedgeFlag:"",
          amount:"",
          
      };
    },
    async deleteAccountPositionDaily(row) {
      const res = await deleteAccountPositionDaily({ ID: row.id });
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
          res = await createAccountPositionDaily(this.formData);
          break;
        case "update":
          res = await updateAccountPositionDaily(this.formData);
          break;
        default:
          res = await createAccountPositionDaily(this.formData);
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
    //账户选项
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
    //获取期货公司字典
    const broker = await getDict("broker");
    broker.map(item=>item.value)
    this.brokerDictList = broker
    await this.getTableData();
    //获取方向字典
    const direct = await getDict("direct");
    direct.map(item=>item.value)
    this.directionDictList = direct
    //获取投机\套保字典
    const hedgeFlag = await getDict("hedgeFlag");
    hedgeFlag.map(item=>item.value)
    this.hedgeFlagDictList = hedgeFlag
    //加载期货账户信息
    const accountInfo = await getAccountInfoList({ page: 1, pageSize: 999 });
    this.setAccountInfoOptions(accountInfo.data.list);
  
}
};
</script>

<style>
</style>
