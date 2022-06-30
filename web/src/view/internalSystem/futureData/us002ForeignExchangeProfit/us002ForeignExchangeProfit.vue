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
        <el-form-item label="币种">
           <el-select v-model="searchInfo.currency" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.currencyCode"
              :label="`${item.currencyName}`"
              :value="item.currencyCode"
              v-for="item in currencyInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>      
        <el-form-item label="部门">
           <el-select v-model="searchInfo.departmentCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.departmentCode"
              :label="`${item.departmentName}`"
              :value="item.departmentCode"
              v-for="item in deptInfoOptions">
            </el-option>
          </el-select>
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
    <el-table-column label="时间" prop="time" width="220">
        <template slot-scope="scope">{{scope.row.time|formatDate}}</template></el-table-column> 
    
    <el-table-column label="币种" prop="currencys.currencyName" width="120"></el-table-column> 
    
    <el-table-column label="金额" prop="amountOfProfit" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="department.departmentName" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUs002ForeignExchangeProfit(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增/编辑">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="时间:" prop="time">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
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
       
         <el-form-item label="金额:">
              <el-input-number v-model="formData.amountOfProfit" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="部门:">
           <el-select v-model="formData.departmentCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.departmentCode"
              :label="`${item.departmentName}`"
              :value="item.departmentCode"
              v-for="item in deptInfoOptions">
            </el-option>
          </el-select>
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
    createUs002ForeignExchangeProfit,
    deleteUs002ForeignExchangeProfit,
    deleteUs002ForeignExchangeProfitByIds,
    updateUs002ForeignExchangeProfit,
    findUs002ForeignExchangeProfit,
    getUs002ForeignExchangeProfitList
} from "@/api/internalSystem/futureData/us002ForeignExchangeProfit";  //  此处请自行替换地址

import {getDepartmentList} from "@/api/internalSystem/futureData/department"; 
import {getUs100CurrencyList} from "@/api/internalSystem/public/us100Currency"; 
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Us002ForeignExchangeProfit",
  mixins: [infoList],
  data() {
    return {
      listApi: getUs002ForeignExchangeProfitList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      deptInfoOptions:[],
      currencyInfoOptions:[],
      multipleSelection: [],formData: {
            time:"",
            currency:"",
            amountOfProfit:0,
            departmentCode:"",
            
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
           this.deleteUs002ForeignExchangeProfit(row);
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
        const res = await deleteUs002ForeignExchangeProfitByIds({ ids })
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
    async updateUs002ForeignExchangeProfit(row) {
      const res = await findUs002ForeignExchangeProfit({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reus002ForeignExchangeProfit;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:"",
          currency:"",
          amountOfProfit:0,
          departmentCode:"",
          
      };
    },
    async deleteUs002ForeignExchangeProfit(row) {
      const res = await deleteUs002ForeignExchangeProfit({ ID: row.ID });
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
          res = await createUs002ForeignExchangeProfit(this.formData);
          break;
        case "update":
          res = await updateUs002ForeignExchangeProfit(this.formData);
          break;
        default:
          res = await createUs002ForeignExchangeProfit(this.formData);
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
    setDeptInfoOptions(deptInfoData) {
        this.deptInfoOptions = [];
        this.ids = [];
        this.setDeptInfoOptionsData(deptInfoData, this.deptInfoOptions ,this.ids);
      },
      setDeptInfoOptionsData(DeptInfoData, optionsData ,ids) {
        DeptInfoData &&
          DeptInfoData.map(item => {
              const option = {
                departmentName: item.departmentName,
                departmentCode: item.departmentCode
              };
              optionsData.push(option);
              const idOption = {
                ID: item.ID,
              };
              ids.push(idOption)
          });
      },
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
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    //加载部门信息
    const deptInfo = await getDepartmentList({ page: 1, pageSize: 999 });
    this.setDeptInfoOptions(deptInfo.data.list);
  
    //币种
    const currencyInfo = await getUs100CurrencyList({ page: 1, pageSize: 999 });
    this.setCurrencyInfoOptions(currencyInfo.data.list);
}
};
</script>

<style>
</style>
