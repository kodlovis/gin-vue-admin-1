<template>
  <div>
      <h1 style="font-size:20px">合约基差录入</h1>
    <div class="search-term">
      <el-form :inline="true" :model="form" class="demo-form-inline">
        <el-form-item label="时间">
          <div class="block">
            <el-date-picker
              v-model="form.time"
              type="datetime"
              placeholder="选择日期时间"
              default-time="12:00:00">
            </el-date-picker>
          </div>
        </el-form-item>         
        <el-form-item >
          <el-button @click="searchInstrumentByDate()" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="deliveryInstrumentOptions"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :key="bomCheckKey"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    
    
    <el-table-column label="时间" prop="time" width="220">
        <template slot-scope="scope">{{scope.row.time|formatDate}}</template></el-table-column> 
    <el-table-column label="品种" prop="product" width="120"></el-table-column> 
    <el-table-column label="近月合约" prop="instrument" width="120"></el-table-column> 
    <!-- <el-table-column label="是否可转抛" prop="is_throw" width="120"></el-table-column>  -->
    
    <el-table-column label="基准合约" width="140">
        <template slot-scope="scope">
           <!-- <el-input v-model.number="formData.accountId" clearable placeholder="请输入"></el-input> -->
           <el-select v-model="scope.row.farMonthInstrument.instrument" placeholder="请选择" clearable filterable @change="calNoRiskValue(scope.row)">
            <el-option
              :key="item.instrument"
              :value="item.instrument"
              v-for="item in scope.row.farMonthInstrument">
            </el-option>
          </el-select>
        </template>
      </el-table-column> 
      
    <el-table-column label="无风险价差" prop="basisWithoutRisk" width="120"></el-table-column> 
    
    <el-table-column label="基差" prop="adjustBasis" width="120">
          <template slot-scope="scope">
              <el-input v-model="scope.row.adjustBasis" clearable placeholder="请输入"  @change="InsertAdjustBasis(scope.row)"></el-input>
          </template>
        </el-table-column> 
    </el-table>
    
          <el-button @click="createBasisOfInstrumentCloseList()" type="primary" style="margin-left: 760px;" >提交录入的基差</el-button>
<el-row>
  <el-col :span="24"><div class="grid-content bg-purple-dark"></div></el-col>
</el-row>
      <h1 style="font-size:20px">历史基差详情</h1>
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

    <el-table-column label="品种" prop="productCode" width="120"></el-table-column> 
    
    <el-table-column label="近月合约" prop="deliveryMonthInstrument" width="120"></el-table-column> 
    
    <el-table-column label="基准合约" prop="benchmarkInstrument" width="140"></el-table-column> 
    
    <el-table-column label="无风险价差" prop="basisWithoutRisk" width="120"></el-table-column> 
    
    <el-table-column label="基差" prop="adjustBasis" width="120"></el-table-column> 
    
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUs002FutureBasisOfInstrumentClose(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="品种:">
            <el-input v-model="formData.productCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="近月合约:">
            <el-input v-model="formData.deliveryMonthInstrument" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="基准合约:">
            <el-input v-model="formData.benchmarkInstrument" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="基差:">
              <el-input-number v-model="formData.adjustBasis" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="无风险价差:">
              <el-input-number v-model="formData.basisWithoutRisk" :precision="2" clearable></el-input-number>
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
    createUs002FutureBasisOfInstrumentClose,
    deleteUs002FutureBasisOfInstrumentClose,
    deleteUs002FutureBasisOfInstrumentCloseByIds,
    updateUs002FutureBasisOfInstrumentClose,
    findUs002FutureBasisOfInstrumentClose,
    getUs002FutureBasisOfInstrumentCloseList,
    getPositionDeliveryMonthInstrumentList,
    //getBenchmarkInstrumentList,
    getNoRiskValue,
    createBasisOfInstrumentCloseList
} from "@/api/internalSystem/futureData/us002FutureBasisOfInstrumentClose";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Us002FutureBasisOfInstrumentClose",
  mixins: [infoList],
  data() {
    return {
      listApi: getUs002FutureBasisOfInstrumentCloseList,
      dialogFormVisible: false,
      type: "",
      bomCheckKey:0,
      deliveryInstrumentOptions:[],
      deleteVisible: false,
      form: {},
      multipleSelection: [],formData: {
            time:new Date(),
            deliveryMonthInstrument:"",
            benchmarkInstrument:"",
            adjustBasis:0,
            basisWithoutRisk:0,
            productCode:"",
            is_throw:"",
      },
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
      //计算无风险价差
      async calNoRiskValue(row){
        const res = await getNoRiskValue({ 
          farMonthInstrument:row.farMonthInstrument.instrument,
          instrument:row.instrument,
          time:row.time
          })
        //匹配合约进行赋值
        if (res.code == 0) {
          for (let i = 0; i < this.deliveryInstrumentOptions.length; i++) {
            for (let j = 0; j < res.data.regetNoRiskValue.length; j++) {
              if (this.deliveryInstrumentOptions[i].instrument==res.data.regetNoRiskValue[j].instrument){
                this.deliveryInstrumentOptions[i].basisWithoutRisk=res.data.regetNoRiskValue[j].noRiskValue
                this.deliveryInstrumentOptions[i].selectedInstrument=res.data.regetNoRiskValue[j].farMonthInstrument
              }
            }
          }
          //及时刷新table数据this.deliveryInstrumentOptions.splice
          this.$nextTick(()=>{
              this.bomCheckKey++;
            })
        }
        
      },
      //匹配合约赋值基差
      InsertAdjustBasis (row){
          for (let i = 0; i < this.deliveryInstrumentOptions.length; i++) {
              if (this.deliveryInstrumentOptions[i].instrument==row.instrument){
                this.deliveryInstrumentOptions[i].adjustBasis=row.adjustBasis
              }
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
           this.deleteUs002FutureBasisOfInstrumentClose(row);
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
        const res = await deleteUs002FutureBasisOfInstrumentCloseByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功,报表同步会有稍许延迟'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateUs002FutureBasisOfInstrumentClose(row) {
      const res = await findUs002FutureBasisOfInstrumentClose({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reus002FutureBasisOfInstrumentClose;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          time:new Date(),
          deliveryMonthInstrument:"",
          benchmarkInstrument:"",
          adjustBasis:0,
          basisWithoutRisk:0,
          productCode:"",
      };
    },
    async deleteUs002FutureBasisOfInstrumentClose(row) {
      const res = await deleteUs002FutureBasisOfInstrumentClose({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功,报表同步会有稍许延迟"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    //批量提交基差信息
    async createBasisOfInstrumentCloseList(){
      var item=[]
      for (let index = 0; index < this.deliveryInstrumentOptions.length; index++) {
        if (Number(this.deliveryInstrumentOptions[index].adjustBasis)!=0){
                item.push({
                  adjustBasis:Number(this.deliveryInstrumentOptions[index].adjustBasis),
                  benchmarkInstrument:this.deliveryInstrumentOptions[index].selectedInstrument,
                  deliveryMonthInstrument:this.deliveryInstrumentOptions[index].instrument,
                  productCode:this.deliveryInstrumentOptions[index].product,
                  basisWithoutRisk:this.deliveryInstrumentOptions[index].basisWithoutRisk,
                  time:this.deliveryInstrumentOptions[index].time,
                  userid:this.userInfo.ID,
                })
            }
      }
      let res;
      res = await createBasisOfInstrumentCloseList({item})
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"提交成功,报表同步会有稍许延迟"
        })
        this.deliveryInstrumentOptions=[]
          //及时刷新table数据this.deliveryInstrumentOptions.splice
          this.$nextTick(()=>{
              this.bomCheckKey++;
            })
        this.getTableData();
      }
    },
    async searchInstrumentByDate(){
        this.page = 1
        this.pageSize = 10
        //加载近月合约信息
        const deliveryInstrument = await getPositionDeliveryMonthInstrumentList({ page: 1, pageSize: 999, ...this.form,userid:this.userInfo.ID  });
        this.setdeliveryInstrumentOptions(deliveryInstrument.data.list);
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createUs002FutureBasisOfInstrumentClose(this.formData);
          break;
        case "update":
          res = await updateUs002FutureBasisOfInstrumentClose(this.formData);
          break;
        default:
          res = await createUs002FutureBasisOfInstrumentClose(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功,报表同步会有稍许延迟"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
    setdeliveryInstrumentOptions(deliveryInstrumentData) {
        this.deliveryInstrumentOptions = [];
        this.setdeliveryInstrumentOptionsData(deliveryInstrumentData, this.deliveryInstrumentOptions );
      },
      setdeliveryInstrumentOptionsData(DeliveryInstrumentData, optionsData ) {
        DeliveryInstrumentData &&
          DeliveryInstrumentData.map(item => {
              const option = {
                time: item.time,
                product: item.product,
                instrument: item.instrument,
                farMonthInstrument:item.farMonthInstrument,
                noRiskValue:item.noRiskValue,
                adjustBasis:"",
                selectedInstrument:"",
                is_throw:item.is_throw
              };
              optionsData.push(option);
          });
      },
  },
  async created() {
    await this.getTableData();
    //加载近月合约信息
    const deliveryInstrument = await getPositionDeliveryMonthInstrumentList({ page: 1, pageSize: 999 ,userid:this.userInfo.ID});
    this.setdeliveryInstrumentOptions(deliveryInstrument.data.list);
    // const benchmarkInstrument = await getBenchmarkInstrumentList({ page: 1, pageSize: 999 });
    // this.setInstrumentInfoOptions(benchmarkInstrument.data.list);
}
};
</script>

<style>
  .el-row {
    margin-bottom: 20px;
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 5px;
    margin-top: 20px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
</style>
