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
          <el-input placeholder="搜索条件" v-model="searchInfo.productName"></el-input>
        </el-form-item>    
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.departmentName"></el-input>
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
                <el-button @click="onDelete" size="mini" type="primary" :disabled="isDisable">确定</el-button>
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
    
    <el-table-column label="时间" prop="time" width="220"></el-table-column> 
    
    <el-table-column label="品种" prop="productName" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="departmentName" width="120"></el-table-column> 
    
    <el-table-column label="三大费用" prop="threeCharges" width="120"></el-table-column> 
    
    <el-table-column label="杂项费用" prop="sundryExpense" width="120"></el-table-column> 
    
    <el-table-column label="期货费用" prop="futureCharges" width="120"></el-table-column> 
    <el-table-column label="费用调整" prop="adjustExpense" width="120"></el-table-column> 
    <el-table-column label="客户利润" prop="customerProfit" width="120"></el-table-column> 
    
    <el-table-column label="品种累计权益" prop="cumulativeRights" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateRightsDetail(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作"  v-dialogDrag>
      <el-form :model="formData" label-position="right" label-width="100px" :rules="rules" ref="prForm">
         <el-form-item label="时间:" prop="time">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
       </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="部门:" prop="departmentName">
            <el-input v-model="formData.departmentName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="三大费用:">
           <el-input-number v-model="formData.threeCharges" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="杂项费用:">
           <el-input-number v-model="formData.sundryExpense" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="期货费用:">
           <el-input-number v-model="formData.futureCharges" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
         <el-form-item label="费用调整:">
           <el-input-number v-model="formData.adjustExpense" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
         <el-form-item label="客户利润:">
           <el-input-number v-model="formData.customerProfit" clearable placeholder="请输入" ></el-input-number>
      </el-form-item>
       
         <el-form-item label="品种累计权益:">
           <el-input-number v-model="formData.cumulativeRights" clearable placeholder="请输入" ></el-input-number>
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
    createRightsDetail,
    deleteRightsDetail,
    deleteRightsDetailByIds,
    updateRightsDetail,
    findRightsDetail,
    getRightsDetailList
} from "@/api/internalSystem/futureData/rightsDetail";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "RightsDetail",
  mixins: [infoList],
  data() {
    return {
      listApi: getRightsDetailList,
      dialogFormVisible: false,
      type: "",
      isDisable:false,
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:"",
            productName:"",
            departmentName:"",
            cumulativeRights:"",
            futureCharges:"",
            sundryExpense:"",
            threeCharges:"",
            customerProfit:"",
            adjustExpense:0,
            
      },
      rules: {
        time:[ { required: true, message: '请输入', trigger: 'blur' }],
        accountId:[ { required: true, message: '请输入', trigger: 'blur' }],
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
           this.deleteRightsDetail(row);
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
        const res = await deleteRightsDetailByIds({ ids })
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
    async updateRightsDetail(row) {
      const res = await findRightsDetail({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rerightsDetail;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:"",
          productName:"",
          departmentName:"",
          cumulativeRights:"",
          futureCharges:"",
          sundryExpense:"",
          threeCharges:"",
          adjustExpense:"",
            customerProfit:"",
          
      };
    },
    async deleteRightsDetail(row) {
      const res = await deleteRightsDetail({ ID: row.id });
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
              res = await createRightsDetail(this.formData);
              break;
            case "update":
              res = await updateRightsDetail(this.formData);
              break;
            default:
              res = await createRightsDetail(this.formData);
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
