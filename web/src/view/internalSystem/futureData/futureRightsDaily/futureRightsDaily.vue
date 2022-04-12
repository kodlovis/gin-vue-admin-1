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
        <el-form-item label="部门">
          <el-input placeholder="搜索条件" v-model="searchInfo.department"></el-input>
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
    
    <el-table-column label="时间" prop="time" width="120">
        <template slot-scope="scope">{{scope.row.time|formatDate}}</template></el-table-column> 
    
    <el-table-column label="累计权益" prop="cumulativeRights" width="120"></el-table-column> 
    
    <el-table-column label="部门" prop="department" width="130"></el-table-column> 
    
    <el-table-column label="品种" prop="product" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateFutureRightsDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="时间:" prop="time">
          <div class="block">
              <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
          </div>
      </el-form-item>
       
         <el-form-item label="累计权益:">
              <el-input-number v-model="formData.cumulativeRights" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="部门:" prop="department">
           <el-select v-model="formData.department" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.name"
              v-for="item in departmentInfoOptions">
            </el-option>
          </el-select>
      </el-form-item>
       
         <el-form-item label="品种:">
           <el-select v-model="formData.productName" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productName"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productName"
              v-for="item in productInfoOptions">
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
    createFutureRightsDaily,
    deleteFutureRightsDaily,
    deleteFutureRightsDailyByIds,
    updateFutureRightsDaily,
    findFutureRightsDaily,
    getFutureRightsDailyList
} from "@/api/internalSystem/futureData/futureRightsDaily";  //  此处请自行替换地址
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "FutureRightsDaily",
  mixins: [infoList],
  data() {
    return {
      listApi: getFutureRightsDailyList,
      dialogFormVisible: false,
      type: "",
      isDisable:false,
      deleteVisible: false,
      productInfoOptions:"",
      dialogTitle:"",
      multipleSelection: [],formData: {
            time:"",
            cumulativeRights:"",
            department:"",
            product:"",
            
      },departmentInfoOptions:[
          {name:"杉贸本部"},
          {name:"智维本部"},
          {name:"总经理一室"},
   ],
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
           this.deleteFutureRightsDaily(row);
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
        const res = await deleteFutureRightsDailyByIds({ ids })
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
    async updateFutureRightsDaily(row) {
      const res = await findFutureRightsDaily({ ID: row.id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.refutureRightsDaily;
        this.openDialog("update");
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:"",
          cumulativeRights:"",
          department:"",
          product:"",
          
      };
    },
    async deleteFutureRightsDaily(row) {
      const res = await deleteFutureRightsDaily({ ID: row.id });
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
              res = await createFutureRightsDaily(this.formData);
              break;
            case "update":
              this.$confirm('确定要修改当前数据吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(async () =>{
              res = await updateFutureRightsDaily(this.formData);
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
              res = await createFutureRightsDaily(this.formData);
              break;
          }
          this.isDisable=false
          if (res.code == 0) {
            this.$message({
              type:"success",
              message:"创建/更改成功"
            })
            this.closeDialog();
            this.isDisable=false
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
    //加载品种信息
    const productInfo = await getExchangeProductInfoList({ page: 1, pageSize: 999 });
    this.setProductInfoOptions(productInfo.data.list);
    await this.getTableData();
  
}
};
</script>

<style>
</style>
