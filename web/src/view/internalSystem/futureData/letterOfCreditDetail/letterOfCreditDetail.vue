<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建人">
          <el-input placeholder="搜索条件" v-model="searchInfo.createdUser"></el-input>
        </el-form-item>              
        <el-form-item label="币种">
           <el-select v-model="formData.currency" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.currencyCode"
              :label="`${item.currencyName}`"
              :value="item.currencyCode"
              v-for="item in currencyInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('create')" type="primary">新增信用证</el-button>
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
    
    <el-table-column label="信用证号" prop="creditId" width="120"></el-table-column> 
    
    <el-table-column label="到期付汇金额" prop="maturityAmount" width="120"></el-table-column> 
    
    <el-table-column label="到期付汇日" prop="maturityDate" width="160">
        <template slot-scope="scope">{{scope.row.maturityDate|formatDate}}</template></el-table-column> 
    
    <el-table-column label="初始汇率" prop="initialRate" width="120"></el-table-column> 
    
    <el-table-column label="购汇汇率" prop="purchaseRate" width="120"></el-table-column> 
    
    <el-table-column label="币种" prop="currencys.currencyName" width="120"></el-table-column> 

    <el-table-column label="品种" prop="productInfo.productName" width="120"></el-table-column> 

    <el-table-column label="创建人" prop="userInfo.nickName" width="120"></el-table-column> 
    
    <el-table-column label="所属部门" prop="departmentInfo.departmentName" width="120"></el-table-column> 
    
    <el-table-column label="创建日期" prop="CreatedAt" width="160">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateLetterOfCreditDetail(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
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
      <el-form :model="formData" label-position="right" label-width="100px">
       
         <el-form-item label="信用证号:">
            <el-input v-model="formData.creditId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="到期付汇金额:">
            <el-input-number v-model="formData.maturityAmount" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="到期付汇日:">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.maturityDate" clearable default-time="12:00:00"></el-date-picker>
       </el-form-item>
       
         <el-form-item label="初始汇率:">
            <el-input-number v-model="formData.initialRate" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="购汇汇率:">
            <el-input-number v-model="formData.purchaseRate" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="币种:">
           <el-select v-model="formData.currency" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.currencyCode"
              :label="`${item.currencyName}`"
              :value="item.currencyCode"
              v-for="item in currencyInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>

         <el-form-item label="品种:"  prop="productName">
           <el-select v-model="formData.productCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productName"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productCode"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
         <el-form-item label="所属部门:" prop="departmentName">
           <el-select v-model="formData.departmentCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.departmentName"
              :label="`${item.departmentName}(${item.departmentCode})`"
              :value="item.departmentCode"
              v-for="item in departmentInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
         <el-form-item label="日期:">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.CreatedAt" clearable default-time="12:00:00"></el-date-picker>
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
    createLetterOfCreditDetail,
    deleteLetterOfCreditDetail,
    deleteLetterOfCreditDetailByIds,
    updateLetterOfCreditDetail,
    findLetterOfCreditDetail,
    getLetterOfCreditDetailList
} from "@/api/internalSystem/futureData/letterOfCreditDetail";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import {getUs100CurrencyList} from "@/api/internalSystem/public/us100Currency"; 
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
import {
    getUserDepartmentListByID,
} from "@/api/internalSystem/public/userDepartment";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "LetterOfCreditDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getLetterOfCreditDetailList,
      dialogFormVisible: false,
      isDisable:false,
      type: "",
      deleteVisible: false,
      currencyInfoOptions:[],
      multipleSelection: [],formData: {
            CreatedAt:"",
            createdUser:"",
            creditId:"",
            initialRate:"",
            purchaseRate:"",
            productCode:"",
            maturityAmount:"",
            maturityDate:new Date(),
            currency:"",
      }
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"])
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
           this.deleteLetterOfCreditDetail(row);
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
        const res = await deleteLetterOfCreditDetailByIds({ ids })
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
    async updateLetterOfCreditDetail(row) {
      const res = await findLetterOfCreditDetail({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reletterOfCreditDetail;
        this.openDialog("update");
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
            createdUser:"",
            creditId:"",
            initialRate:"",
            purchaseRate:"",
            productCode:"",
            maturityRate:"",
            maturityDate:new Date(),
            currency:"",
          
      };
    },
    async deleteLetterOfCreditDetail(row) {
      const res = await deleteLetterOfCreditDetail({ ID: row.ID });
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
          res = await createLetterOfCreditDetail({...this.formData,createdUser:this.userInfo.ID});
          break;
        case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateLetterOfCreditDetail(this.formData);
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
          res = await createLetterOfCreditDetail(this.formData);
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
    },
    openDialog(type) {
      switch (type) {
        case "create":
          this.dialogTitle = "新增信用证";
          break;
        case "update":
          this.dialogTitle = "编辑信用证";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    //币种list
    setCurrencyInfoOptions(currencyInfoData) {
        this.currencyInfoOptions = [];
        this.ids = [];
        this.setCurrencyInfoOptionsData(currencyInfoData, this.currencyInfoOptions ,this.ids);
      },
      setCurrencyInfoOptionsData(CurrencyInfoData, optionsData ,ids) {
        CurrencyInfoData &&
          CurrencyInfoData.map(item => {
              const option = {
                currencyName: item.currencyName,
                currencyCode: item.currencyCode
              };
              optionsData.push(option);
              const idOption = {
                ID: item.ID,
              };
              ids.push(idOption)
          });
      },
    //品种list
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
    //部门list
    setDepartmentInfoOptions(departmentInfoData) {
        this.departmentInfoOptions = [];
        this.ids = [];
        this.setDepartmentInfoOptionsData(departmentInfoData, this.departmentInfoOptions ,this.ids);
      },
      setDepartmentInfoOptionsData(DepartmentInfoData, optionsData ,ids) {
        DepartmentInfoData &&
          DepartmentInfoData.map(item => {
              const option = {
                departmentCode: item.departmentCode,
                departmentName: item.departmentInfo.departmentName
              };
              optionsData.push(option);
              const idOption = {
                departmentId: item.departmentId,
              };
              ids.push(idOption)
          });
      },
  },
  async created() {
    await this.getTableData();
    //币种
    const currencyInfo = await getUs100CurrencyList({ page: 1, pageSize: 999 });
    this.setCurrencyInfoOptions(currencyInfo.data.list);
    //加载品种信息
    const productInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setProductInfoOptions(productInfo.data.list);

    //根据用户获取部门信息
    const departmentInfo = await getUserDepartmentListByID({ page: 1, pageSize: 999,userId:this.userInfo.ID });
    this.setDepartmentInfoOptions(departmentInfo.data.list);
    
  
}
};
</script>

<style>
</style>
