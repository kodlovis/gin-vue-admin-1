<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增风控设置</el-button>
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
    
    <el-table-column label="名称" prop="name" width="120"></el-table-column> 
    <el-table-column label="代码" prop="code" width="220"></el-table-column> 
    <el-table-column label="券商阈值" prop="brokersValue" width="120"></el-table-column> 
    <el-table-column label="风控阈值" prop="adminValue" width="120"></el-table-column> 
    <el-table-column label="描述" prop="description" width="220"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateRiskSetting(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteRiskSetting(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      background
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="编辑风控设置">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="代码:">
            <el-input v-model="formData.code" clearable placeholder="请输入" 
            :disabled="isDisable"></el-input>
      </el-form-item>
         <el-form-item label="券商阈值:">
            <el-input v-model="formData.brokersValue" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="风控阈值:">
            <el-input v-model="formData.adminValue" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="描述:">
            <el-input v-model="formData.description" type="textarea" :autosize="{ minRows: 5, maxRows: 10}"  clearable placeholder="请输入" ></el-input>
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
    createRiskSetting,
    deleteRiskSetting,
    deleteRiskSettingByIds,
    updateRiskSetting,
    findRiskSetting,
    getRiskSettingList
} from "@/api/rms/riskSetting";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "RiskSetting",
  mixins: [infoList],
  data() {
    return {
      listApi: getRiskSettingList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      isDisable:false,
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            code:"",
            description:"",
            brokersValue:0,
            adminValue:0,
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
        const res = await deleteRiskSettingByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateRiskSetting(row) {
      const res = await findRiskSetting({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.isDisable = true;
        this.formData = res.data.reRiskSetting;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.isDisable = false
      this.formData = {
            name:"",
            code:"",
            description:"",
            brokersValue:0,
            adminValue:0,
          
      };
    },
    async deleteRiskSetting(row) {
      this.visible = false;
      const res = await deleteRiskSetting({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createRiskSetting({...this.formData,adminValue:Number(this.formData.adminValue),brokersValue:Number(this.formData.brokersValue)});
          break;
        case "update":
          res = await updateRiskSetting({...this.formData,adminValue:Number(this.formData.adminValue),brokersValue:Number(this.formData.brokersValue)});
          break;
        default:
          res = await createRiskSetting({...this.formData,adminValue:Number(this.formData.adminValue),brokersValue:Number(this.formData.brokersValue)});
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