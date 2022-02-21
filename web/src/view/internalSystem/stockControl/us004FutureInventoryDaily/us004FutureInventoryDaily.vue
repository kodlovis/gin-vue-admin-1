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
        <el-form-item label="品种">
          <el-select v-model="searchInfo.productCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productCode"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productCode"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>      
        <el-form-item label="交易所">
            <el-select v-model="searchInfo.exchangeId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.exchangeId"
              :label="`${item.comment}(${item.exchangeId})`"
              :value="item.exchangeId"
              v-for="item in exchangeInfoData">
            </el-option>
          </el-select>
        </el-form-item>    
        <el-form-item label="简称">
              <el-select v-model="searchInfo.comment" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.name"
              v-for="item in commentChoice">
            </el-option>
          </el-select>
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

    <el-table-column label="时间" width="180">
      <template slot-scope="scope">{{scope.row.time|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="数量" prop="volume" width="120"></el-table-column> 
    <el-table-column label="单位" prop="unit" width="120"></el-table-column> 
    
    <el-table-column label="交易所" prop="exchangeId" width="120"></el-table-column> 
    
    <el-table-column label="简称" prop="comment" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUs004FutureInventoryDaily(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
        <el-form-item label="时间:">
          <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable default-time="12:00:00"></el-date-picker>
        </el-form-item>
       
        <el-form-item label="品种:">
            <el-select v-model="formData.productCode" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.productCode"
              :label="`${item.productName}(${item.productCode})`"
              :value="item.productCode"
              v-for="item in productInfoOptions">
            </el-option>
          </el-select>
        </el-form-item>
    
        <el-form-item label="简称:">
           <el-select v-model="formData.comment" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.name"
              v-for="item in commentChoice">
            </el-option>
          </el-select>
        </el-form-item>
       
        <el-form-item label="交易所:">
           <el-select v-model="formData.exchangeId" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.exchangeId"
              :label="`${item.comment}(${item.exchangeId})`"
              :value="item.exchangeId"
              v-for="item in exchangeInfoData">
            </el-option>
          </el-select>
        </el-form-item>
       
        <el-form-item label="数量:">
          <el-input style="width:220px" v-model="formData.volume" clearable placeholder="请输入" ></el-input>
        </el-form-item>
        <el-form-item label="单位:">
           <el-select v-model="formData.unit" placeholder="请选择" clearable filterable >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.name"
              v-for="item in unitData">
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
    createUs004FutureInventoryDaily,
    deleteUs004FutureInventoryDaily,
    deleteUs004FutureInventoryDailyByIds,
    updateUs004FutureInventoryDaily,
    findUs004FutureInventoryDaily,
    getUs004FutureInventoryDailyList
} from "@/api/internalSystem/stockControl/us004FutureInventoryDaily";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import {getExchangeProductInfoList} from "@/api/internalSystem/exchangeProductInfo"; 
import infoList from "@/mixins/infoList";
export default {
  name: "Us004FutureInventoryDaily",
  mixins: [infoList],
  data() {
    return {
      listApi: getUs004FutureInventoryDailyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            time:new Date(),
            productCode:"CU",
            exchangeId:"",
            comment:"",
            unit:"",
            
            
      },
      productInfo:{
        productInfoCode:"",
        productInfoName:"",
      }, 
      exchangeInfoData:[
        {exchangeId:"SHF",comment:"上海期货交易所"},
        // {exchangeId:"ZCE",comment:"郑州商品交易所"},
        // {exchangeId:"DCE",comment:"大连商品交易所"},
        {exchangeId:"INE",comment:"上海能源交易所"},
        {exchangeId:"LME",comment:"LME"},
      ], 
      commentChoice:[
        {name:"现货"},
        {name:"保税"},
        {name:"LME"},
      ], 
      unitData:[
        {name:"吨"},
        {name:"万吨"},
        {name:"磅"},
      ]
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd");
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
           this.deleteUs004FutureInventoryDaily(row);
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
        const res = await deleteUs004FutureInventoryDailyByIds({ ids })
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
    async updateUs004FutureInventoryDaily(row) {
      const res = await findUs004FutureInventoryDaily({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reus004FutureInventoryDaily;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:new Date(),
          productCode:"CU",
          exchangeId:"",
          comment:"",
          
      };
    },
    async deleteUs004FutureInventoryDaily(row) {
      const res = await deleteUs004FutureInventoryDaily({ ID: row.ID });
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
      console.log(this.formData)
      switch (this.type) {
        case "create":
          res = await createUs004FutureInventoryDaily(this.formData);
          break;
        case "update":
          res = await updateUs004FutureInventoryDaily(this.formData);
          break;
        default:
          res = await createUs004FutureInventoryDaily(this.formData);
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
