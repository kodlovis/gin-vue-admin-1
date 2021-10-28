<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">            
        <el-form-item label="生效日期">
          <el-input placeholder="搜索条件" v-model="searchInfo.effectiveDate"></el-input>
        </el-form-item>    
        <el-form-item label="失效日期">
          <el-input placeholder="搜索条件" v-model="searchInfo.expirationDate"></el-input>
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
    
    <el-table-column label="交易所" prop="exchangeId" width="120"></el-table-column> 
    
    <el-table-column label="期货公司" prop="brokerId" width="120">
      <template slot-scope="scope">
          <span>{{brokerfilterDict(scope.row.brokerId)}}</span>
      </template>
    </el-table-column> 
    
    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="基准日均持仓" prop="basePosition" width="120"></el-table-column> 
    
    <el-table-column label="返还类型" prop="rewardType" width="120">
      <template slot-scope="scope">
          <span>{{rewardTypefilterDict(scope.row.rewardType)}}</span>
      </template>
    </el-table-column> 
    
    <el-table-column label="返还数量" prop="rewardAmount" width="120"></el-table-column> 
    
    <el-table-column label="生效日期" prop="effectiveDate" width="120"></el-table-column> 
    
    <el-table-column label="失效日期" prop="expirationDate" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePositionRewardRule(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="120px">
         <el-form-item label="交易所:">
            <el-input v-model="formData.exchangeId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="期货公司:">
          <el-select v-model="formData.brokerId" placeholder="请选择">
            <el-option
              v-for="item in brokerDictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>  
      </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="基准日均持仓:">
              <el-input-number v-model="formData.basePosition" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="返还类型:">
          <el-select v-model="formData.rewardType" placeholder="请选择">
            <el-option
              v-for="item in rewardTypeDictList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
      </el-form-item>
       
         <el-form-item label="返还数量:">
              <el-input-number v-model="formData.rewardAmount" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="生效日期:">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.effectiveDate" clearable default-time="12:00:00"></el-date-picker>
          </div>
      </el-form-item>
       
         <el-form-item label="失效日期:">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.expirationDate" clearable default-time="12:00:00"></el-date-picker>
          </div>
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
    createPositionRewardRule,
    deletePositionRewardRule,
    deletePositionRewardRuleByIds,
    updatePositionRewardRule,
    findPositionRewardRule,
    getPositionRewardRuleList
} from "@/api/internalSystem/positionReward/positionRewardRule";  //  此处请自行替换地址
import { getDict } from "@/utils/dictionary";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PositionRewardRule",
  mixins: [infoList],
  data() {
    return {
      listApi: getPositionRewardRuleList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      rewardTypeDictList:[],
      brokerDictList:[],
      multipleSelection: [],formData: {
            exchangeId:"",
            brokerId:"",
            productCode:"",
            basePosition:0,
            rewardType:"",
            rewardAmount:0,
            effectiveDate:"",
            expirationDate:"",
            
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
      rewardTypefilterDict(rewardType){
        const re = this.rewardTypeDictList.filter(item=>{
          return item.value == rewardType
        })[0]
        if(re){
          return re.label
          }
        else{
          return""
          }
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
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deletePositionRewardRule(row);
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
        const res = await deletePositionRewardRuleByIds({ ids })
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
    async updatePositionRewardRule(row) {
      const res = await findPositionRewardRule({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.repositionRewardRule;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          exchangeId:"",
          brokerId:"",
          productCode:"",
          basePosition:0,
          rewardType:"",
          rewardAmount:0,
          effectiveDate:"",
          expirationDate:"",
          
      };
    },
    async deletePositionRewardRule(row) {
      const res = await deletePositionRewardRule({ ID: row.id });
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
          res = await createPositionRewardRule(this.formData);
          break;
        case "update":
          res = await updatePositionRewardRule(this.formData);
          break;
        default:
          res = await createPositionRewardRule(this.formData);
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
    //获取返还类型字典
    const rewardType = await getDict("rewardType");
    rewardType.map(item=>item.value)
    this.rewardTypeDictList = rewardType
    //获取期货公司字典
    const broker = await getDict("broker");
    broker.map(item=>item.value)
    this.brokerDictList = broker
  
}
};
</script>

<style>
</style>
