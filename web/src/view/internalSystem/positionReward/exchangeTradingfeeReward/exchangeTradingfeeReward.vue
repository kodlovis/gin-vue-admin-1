<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="生效日期">
          <div class="block">
            <el-date-picker
              v-model="searchInfo.effectiveDate"
              type="datetime"
              placeholder="选择日期时间"
              default-time="8:00:00">
            </el-date-picker>
          </div>
        </el-form-item>
        <!-- <el-form-item label="交易所名称">
           <el-select v-model="searchInfo.exchangeId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.exchangeId"
              :label="`${item.exchangeId}`"
              :value="item.exchangeId"
              v-for="item in exchangeInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>                  -->
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('create')" type="primary">新增规则</el-button>
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
        <!-- <el-form-item>
              <el-upload
                :action="`${path}/excel/importExcel`"
                :headers="{'x-token':token}"
                :on-success="loadExchangeTradingfeeRewardExcel"
                :show-file-list="false"
              >
                <el-button size="small" type="primary" icon="el-icon-upload2">Excel导入</el-button>
              </el-upload>
        </el-form-item> -->
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
    
    <el-table-column label="生效日期" prop="effectiveDate" width="220"></el-table-column> 
    
    <el-table-column label="交易所名称 " prop="exchangeId" width="120"></el-table-column> 
    
    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="手续费返还率" prop="rewardRate" width="120"></el-table-column> 
    
    <el-table-column label="描述" prop="description" width="420"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateExchangeTradingfeeReward(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      <el-row :gutter="9">
        <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="93px"
          label-position="left">
          <el-col :span="15">
            <el-form-item label="生效日期">
              <el-date-picker v-model="formData.effectiveDate" clearable default-time="12:00:00"
                :style="{width: '100%'}" placeholder="请选择生效日期" ></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="18">
            <el-form-item label="交易所代号">
              <el-input v-model="formData.exchangeId" placeholder="请输入交易所代号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="18">
            <el-form-item label="品种代码" prop="productCode">
              <el-input v-model="formData.productCode" placeholder="请输入品种代码" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="18">
            <el-form-item label="返还费率" prop="rewardRate">
              <el-input v-model="formData.rewardRate" precision="4" placeholder="请输入返还费率" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="22">
            <el-form-item label="描述" prop="description">
              <el-input v-model="formData.description" type="textarea" placeholder="请输入描述"
                :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
        </el-form>
      </el-row>
      <div slot="footer">
        <el-button @click="closeDialog">取消</el-button>
        <el-button type="primary" :disabled="isDisable" @click="enterDialog">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createExchangeTradingfeeReward,
    deleteExchangeTradingfeeReward,
    deleteExchangeTradingfeeRewardByIds,
    updateExchangeTradingfeeReward,
    findExchangeTradingfeeReward,
    getExchangeTradingfeeRewardList,
    loadExchangeTradingfeeRewardExcelData
} from "@/api/internalSystem/positionReward/exchangeTradingfeeReward";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
import infoList from "@/mixins/infoList";
export default {
  name: "ExchangeTradingfeeReward",
  mixins: [infoList],
  data() {
    return {
      listApi: getExchangeTradingfeeRewardList,
      dialogFormVisible: false,
      type: "",
      isDisable:false,
      deleteVisible: false,
      multipleSelection: [],formData: {
          effectiveDate:"",
          exchangeId:"",
          productCode:"",
          rewardRate:0,
          description:"",
          exchangeInfoOptions:"",

      },
      exchangeInfo:{
        echangegeId:"",
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
    async loadExchangeTradingfeeRewardExcel() {
        const res = await  loadExchangeTradingfeeRewardExcelData();
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
           this.deleteExchangeTradingfeeReward(row);
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
        const res = await deleteExchangeTradingfeeRewardByIds({ ids })
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
    async updateExchangeTradingfeeReward(row) {
      const res = await findExchangeTradingfeeReward({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reexchangeTradingfeeReward;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          effectiveDate:"",
          exchangeId:"",
          productCode:"",
          rewardRate:0,
          description:"",
          
      };
    },
    async deleteExchangeTradingfeeReward(row) {
      const res = await deleteExchangeTradingfeeReward({ ID: row.id });
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
          res = await createExchangeTradingfeeReward(this.formData);
          break;
        case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateExchangeTradingfeeReward(this.formData);
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
          res = await createExchangeTradingfeeReward(this.formData);
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
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
    setExchangeInfoOptions(exchangeInfoData) {
        this.exchangeInfoOptions = [];
        this.ids = [];
        this.setexchangeInfoOptionsData(exchangeInfoData, this.exchangeInfoOptions ,this.ids);
      },
      setexchangeInfoOptionsData(exchangeInfoData, optionsData ,ids) {
        exchangeInfoData &&
          exchangeInfoData.map(item => {
              const option = {
                productCode: item.productCode,
                exchangeId: item.exchangeName
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
    const exchangeInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setExchangeInfoOptions(exchangeInfo.data.list);
    await this.getTableData();
    //加载品种信息
    const productInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setProductInfoOptions(productInfo.data.list);
  
}
};
</script>

<style>
</style>
