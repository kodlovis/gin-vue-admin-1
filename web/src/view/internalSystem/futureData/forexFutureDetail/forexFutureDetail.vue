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
           <el-select v-model="searchInfo.productName" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productName"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productName"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>    
        <el-form-item label="账户">
           <el-select v-model="searchInfo.accountId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.accountId"
              :label="`${item.comment}(${item.accountId})`"
              :value="item.accountId"
              v-for="item in accountInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>    
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.departmentName"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('create')" type="primary">新增数据</el-button>
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
    
    <el-table-column label="时间" prop="time" width="220">
        <template slot-scope="scope">{{scope.row.time|formatDate}}</template></el-table-column> 
    
    <el-table-column label="品种" prop="productName" width="120"></el-table-column> 
    
    <el-table-column label="账户" prop="accountInfo.comment" width="150"></el-table-column> 
    
    <el-table-column label="浮动盈亏" prop="profitByFloat" width="120"></el-table-column> 
    
    <el-table-column label="平仓盈亏" prop="profitByTrade" width="120"></el-table-column> 
    
    <el-table-column label="手续费" prop="tradeFee" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="departmentName" width="130"></el-table-column> 
    
      <el-table-column label="操作">
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogTitle" v-dialogDrag>
      <el-form :model="formData" label-position="right" label-width="100px" :rules="rules" ref="prForm">
         <el-form-item label="时间:" prop="time">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
          </div>
       </el-form-item>
       
         <el-form-item label="品种:" prop="productName">
           <el-select v-model="formData.productName" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productName"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productName"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
       
         <el-form-item label="账号:" prop="accountId">
           <el-select v-model="formData.accountId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.accountId"
              :label="`${item.comment}(${item.accountId})`"
              :value="item.accountId"
              v-for="item in accountInfoOptions">
            </el-option>
          </el-select>
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
       
         <el-form-item label="部门:" prop="departmentName">
           <el-select v-model="formData.departmentName" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.name"
              v-for="item in departmentInfoOptions">
            </el-option>
          </el-select>
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
import {
    createForexFutureDetail,
    deleteForexFutureDetail,
    deleteForexFutureDetailByIds,
    updateForexFutureDetail,
    findForexFutureDetail,
    getForexFutureDetailList
} from "@/api/internalSystem/futureData/forexFutureDetail";  //  此处请自行替换地址
import{getAccountInfoList}from "@/api/internalSystem/accountInfo"; 
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
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
      isDisable:false,
      deleteVisible: false,
      productInfoOptions:"",
      accountInfoOptions:[],
      dialogTitle:"",
      multipleSelection: [],formData: {
            time:"",
            productName:"",
            accountId:"",
            profitByFloat:"",
            profitByTrade:"",
            tradeFee:"",
            departmentName:"",
            accountInfo:{accountId:"",comment:""},
            
      },departmentInfoOptions:[
          {name:"有色金属部"},
          {name:"农产品部"},
          {name:"黑色建材部"},
          {name:"总经理二室"},
          {name:"钢材事业部_杉贸"},
          {name:"原料事业部"},
          {name:"钢材事业部"},
          {name:"能化橡胶部"},
   ],
      accountInfoData:{
           accountId:"",
           comment:"",
      }, 
      productInfo:{
           productInfoCode:"",
           productInfoName:"",
      }, 
      rules: {
        time:[ { required: true, message: '请输入', trigger: 'blur' }],
        accountId:[ { required: true, message: '请输入', trigger: 'blur' }],
        productName:[ { required: true, message: '请输入', trigger: 'blur' }],
        departmentName:[ { required: true, message: '请输入', trigger: 'blur' }],
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
            ids.push(item.id)
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
        this.openDialog("update");
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
      this.$refs.prForm.validate(async valid => {
        if (valid) {
          this.isDisable=true;
          let res;
          switch (this.type) {
            case "create":
              res = await createForexFutureDetail(this.formData);
              break;
            case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateForexFutureDetail(this.formData);
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
              res = await createForexFutureDetail(this.formData);
              break;
          }
          this.isDisable=false
          if (res.code == 0) {
            this.$message({
              type:"success",
              message:"创建/更改成功"
            })
            this.closeDialog();
            this.getTableData();
          }
        }
      });
    },
    openDialog(type) {
      switch (type) {
        case "create":
          this.dialogTitle = "新增数据";
          break;
        case "update":
          this.dialogTitle = "编辑数据";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    setAccountInfoOptions(accountInfoData) {
        this.accountInfoOptions = [];
        this.ids = [];
        this.setAccountInfoOptionsData(accountInfoData, this.accountInfoOptions ,this.ids);
      },
      setAccountInfoOptionsData(AccountInfoData, optionsData ,ids) {
        AccountInfoData &&
          AccountInfoData.map(item => {
            if(item.type=='3'){
              const option = {
                accountId: item.accountId,
                comment: item.comment
              };
              optionsData.push(option);
              const idOption = {
                accountId: item.accountId,
              };
              ids.push(idOption)}
          });
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
  },
      
  async created() {
    //加载期货账户信息
    const accountInfo = await getAccountInfoList({ page: 1, pageSize: 999 });
    this.setAccountInfoOptions(accountInfo.data.list);
    //加载品种信息
    const productInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setProductInfoOptions(productInfo.data.list);
    await this.getTableData();
  
}
};
</script>

<style>
</style>
