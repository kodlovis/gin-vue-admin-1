<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="时间">
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
        <el-form-item>
              <el-upload
                :action="`${path}/excel/importExcel`"
                :headers="{'x-token':token}"
                :on-success="loadUs001DceBrokerPositionExcel"
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
    
    <el-table-column label="时间" prop="time" width="120"></el-table-column> 
    
    <el-table-column label="期货公司" prop="brokerId" width="120"></el-table-column> 
    
    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="一般法人客户做市商合约日均持仓" prop="corporateMarketPosition" width="120"></el-table-column> 
    
    <el-table-column label="一般法人客户非做市商合约日均持仓" prop="corporateNonMarketPosition" width="120"></el-table-column> 
    
    <el-table-column label="其他类型客户做市商日均持仓" prop="otherMarketPosition" width="120"></el-table-column> 
    
    <el-table-column label="其他类型客户非做市商日均持仓" prop="otherNonMarketPosition" width="120"></el-table-column> 
    
    <el-table-column label="上缴手续费" prop="payFee" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUs001DceBrokerPositionDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增、编辑">
      <el-form :model="formData" label-position="right" label-width="180px">
         <el-form-item label="时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="期货公司:">
            <el-input v-model="formData.brokerId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="一般法人客户做市商合约日均持仓:">
            <el-input v-model="formData.corporateMarketPosition" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="一般法人客户非做市商合约日均持仓:">
            <el-input v-model="formData.corporateNonMarketPosition" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="其他类型客户做市商日均持仓:">
            <el-input v-model="formData.otherMarketPosition" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="其他类型客户非做市商日均持仓:">
            <el-input v-model="formData.otherNonMarketPosition" clearable placeholder="请输入" ></el-input></el-form-item>
       
         <el-form-item label="上缴手续费:">
            <el-input v-model="formData.payFee" clearable placeholder="请输入" ></el-input></el-form-item>
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
    createUs001DceBrokerPositionDaily,
    deleteUs001DceBrokerPositionDaily,
    deleteUs001DceBrokerPositionDailyByIds,
    updateUs001DceBrokerPositionDaily,
    findUs001DceBrokerPositionDaily,
    getUs001DceBrokerPositionDailyList,
    loadUs001DceBrokerPositionExcelData
} from "@/api/internalSystem/positionReward/dce/us001DceBrokerPositionDaily";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Us001DceBrokerPositionDaily",
  mixins: [infoList],
  data() {
    return {
      path: path,
      listApi: getUs001DceBrokerPositionDailyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:new Date(),
            brokerId:"",
            productCode:"",
            
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
      //条件搜索前端看此方法
      async loadUs001DceBrokerPositionExcel() {
        const res = await loadUs001DceBrokerPositionExcelData();
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '上传成功'
          })
          this.getTableData();
        }
      },
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
           this.deleteUs001DceBrokerPositionDaily(row);
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
        const res = await deleteUs001DceBrokerPositionDailyByIds({ ids })
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
    async updateUs001DceBrokerPositionDaily(row) {
      const res = await findUs001DceBrokerPositionDaily({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reus001DceBrokerPositionDaily;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:new Date(),
          brokerId:"",
          productCode:"",
          
      };
    },
    async deleteUs001DceBrokerPositionDaily(row) {
      const res = await deleteUs001DceBrokerPositionDaily({ ID: row.ID });
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
          res = await createUs001DceBrokerPositionDaily(this.formData);
          break;
        case "update":
          res = await updateUs001DceBrokerPositionDaily(this.formData);
          break;
        default:
          res = await createUs001DceBrokerPositionDaily(this.formData);
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
