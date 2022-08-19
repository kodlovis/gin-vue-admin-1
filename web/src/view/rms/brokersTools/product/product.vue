<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增产品</el-button>
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
    
    <el-table-column label="产品名称" prop="name" width="120"></el-table-column> 
    <el-table-column label="账号" prop="accountid" width="120"></el-table-column> 
    <el-table-column label="柜台信息" prop="counterInfo" width="120"></el-table-column> 
    <el-table-column label="节点信息" prop="nodeInfo" width="120"></el-table-column> 
    <el-table-column label="备注" prop="comment" width="120"></el-table-column> 
    <el-table-column label="描述" prop="description" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateProduct(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteProduct(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="编辑产品">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="产品名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="账号:">
            <el-input v-model="formData.accountid" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="柜台信息:">
            <el-input v-model="formData.counterInfo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="节点信息:">
            <el-input v-model="formData.nodeInfo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="备注:">
            <el-input v-model="formData.comment" clearable placeholder="请输入" ></el-input>
      </el-form-item>
         <el-form-item label="描述:">
            <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
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
    createProduct,
    deleteProduct,
    deleteProductByIds,
    updateProduct,
    findProduct,
    getProductList
} from "@/api/rms/product";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Product",
  mixins: [infoList],
  data() {
    return {
      listApi: getProductList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            accountid: "",
            counterInfo: "",
            password: "",
            nodeInfo: "",
            comment: "",
            description: "",
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
        const res = await deleteProductByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateProduct(row) {
      const res = await findProduct({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reProduct;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          Name:"",
          
      };
    },
    async deleteProduct(row) {
      this.visible = false;
      const res = await deleteProduct({ ID: row.ID });
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
          res = await createProduct(this.formData);
          break;
        case "update":
          res = await updateProduct(this.formData);
          break;
        default:
          res = await createProduct(this.formData);
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