<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="日期">
          <div class="block">
            <el-date-picker
              v-model="searchInfo.time"
              type="datetime"
              placeholder="选择日期时间"
              default-time="12:00:00">
            </el-date-picker>
          </div>
        </el-form-item>  
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.department"></el-input>
        </el-form-item>    
        <el-form-item label="类型">
          <el-input placeholder="搜索条件" v-model="searchInfo.type"></el-input>
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
    
    <el-table-column label="日期" prop="time" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="department" width="120"></el-table-column> 
    
    <el-table-column label="类型" prop="type" width="120"></el-table-column> 
    
    <el-table-column label="盈亏" prop="profit" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateFutureDeparmentRights(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="日期:" prop="time">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
          </div>
       </el-form-item>
       
         <el-form-item label="部门:">
            <el-input v-model="formData.department" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="类型:">
            <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="盈亏:">
            <el-input-number v-model="formData.profit" clearable placeholder="请输入" ></el-input-number>
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
    createFutureDeparmentRights,
    deleteFutureDeparmentRights,
    deleteFutureDeparmentRightsByIds,
    updateFutureDeparmentRights,
    findFutureDeparmentRights,
    getFutureDeparmentRightsList
} from "@/api/internalSystem/scaleRights/futureDeparmentRights";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "FutureDeparmentRights",
  mixins: [infoList],
  data() {
    return {
      listApi: getFutureDeparmentRightsList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:"",
            department:"",
            type:"",
            
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
           this.deleteFutureDeparmentRights(row);
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
        const res = await deleteFutureDeparmentRightsByIds({ ids })
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
    async updateFutureDeparmentRights(row) {
      const res = await findFutureDeparmentRights({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.refutureDeparmentRights;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:"",
          department:"",
          type:"",
          
      };
    },
    async deleteFutureDeparmentRights(row) {
      const res = await deleteFutureDeparmentRights({ ID: row.id });
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
          res = await createFutureDeparmentRights(this.formData);
          break;
        case "update":
          res = await updateFutureDeparmentRights(this.formData);
          break;
        default:
          res = await createFutureDeparmentRights(this.formData);
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
