<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <!-- <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增用户</el-button>
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
    
    <el-table-column label="用户名称" prop="name" width="120"></el-table-column> 
    <el-table-column label="角色" prop="role.name" width="120"></el-table-column> 
    <el-table-column label="mac地址" width="120">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Macs"
            :key="index">{{item.name}}<br/></span>
      </template>
    </el-table-column> 
    <el-table-column label="产品" width="120">
      <template slot-scope="scope">
        <span v-for="(item,index) in scope.row.Products"
            :key="index">{{item.name}}<br/></span>
      </template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateUsers(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <!-- <el-button class="table-button" @click="updateMacs(scope.row)" size="small" type="primary" icon="el-icon-edit">mac管理</el-button>
          <el-button class="table-button" @click="updateProducts(scope.row)" size="small" type="primary" icon="el-icon-edit">产品管理</el-button> -->
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUsers(scope.row)">确定</el-button>
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
      :page-sizes="[3,5,10,20,30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :rules="rules" title="编辑用户">
      <el-form :model="formData" label-position="right" label-width="120px"  ref="userForm">
         <el-form-item label="用户:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
          <el-form-item label="添加角色" prop="roleid"> 
              <el-cascader
                @change="(val)=>{handleOptionChange(val)}"
                v-model="formData.roleid"
                :options="roleOptions"
                clearable
                :props="{ checkStrictly: true,label:'name',value:'id'}"
                filterable
              ></el-cascader>
          </el-form-item>
          <el-form-item label="添加产品"> 
              <el-cascader
                @change="(val)=>{handleOptionChange(val)}"
                v-model="formData.Products"
                :options="productOptions"
                clearable
                :props="{ checkStrictly: true,label:'name',value:'id',multiple: true}"
                filterable
              ></el-cascader>
          </el-form-item>
        <el-form-item label="mac地址一:">
            <el-input v-model="formData.Macs.name" clearable placeholder="请输入" ></el-input>
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
    createUsers,
    deleteUsers,
    deleteUsersByIds,
    updateUsers,
    findUsers,
    getUsersList,

    getLastUser,
    createUserMac,
    createUserProduct,
    removeUserMacProduct,
} from "@/api/rms/users";  //  此处请自行替换地址
import {getProductList} from "@/api/rms/product";
import {getRoleList} from "@/api/rms/role";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Users",
  mixins: [infoList],
  data() {
    return {
      listApi: getUsersList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      productOptions:[],
      roleOptions:[],
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            roleid:0,
            Macs:[{name:""}],
            Products:[{id:"",name:""}],
            role:[{name:""}],
      },
      productData:{name:""},
      roleData:{name:""},
      rules: {
        name:[ { required: true, message: '请输入', trigger: 'blur' }],
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
      
      setProductOptions(productData) {
        this.productOptions = [];
        this.ids = [];
        this.setProductOptionsData(productData, this.productOptions ,this.ids);
      },
      setProductOptionsData(ProductData, optionsData ,ids) {
        ProductData &&
          ProductData.map(item => {
              const option = {
                id: item.ID,
                name: item.name
              };
              optionsData.push(option);
              const idOption = {
                id: item.ID,
              };
              ids.push(idOption)
          });
      },
      setRoleOptions(roleData) {
        this.roleOptions = [];
        this.ids = [];
        this.setRoleOptionsData(roleData, this.roleOptions ,this.ids);
      },
      setRoleOptionsData(RoleData, optionsData ,ids) {
        RoleData &&
          RoleData.map(item => {
              const option = {
                id: item.ID,
                name: item.name
              };
              optionsData.push(option);
              const idOption = {
                id: item.ID,
              };
              ids.push(idOption)
          });
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
        const res = await deleteUsersByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateUsers(row) {
      const res = await findUsers({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reUsers;
        const arr = []
        for (let index = 0; index < this.formData.Products.length; index++) {
          arr.push(this.formData.Products[index].ID)
        }
        this.formData.Products = arr
        this.dialogFormVisible = true;
      }
    },
    // async updateMacs(row) {
    //   const res = await findUsers({ ID: row.ID });
    //   if (res.code == 0) {
    //     this.formData = res.data.reUsers;
    //     this.dialogMacsFormVisible = true;
    //   }
    // },
    // async updateProducts(row) {
    //   const res = await findUsers({ ID: row.ID });
    //   if (res.code == 0) {
    //     this.formData = res.data.reUsers;
    //     this.dialogProductsFormVisible = true;
    //   }
    // },
    closeDialog() {
      this.dialogFormVisible = false;
      // this.formData = {
      //     name:"",
          
      // };
    },
    async deleteUsers(row) {
      this.visible = false;
      const res = await deleteUsers({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    // 初始化表单
    initForm() {
      if (this.$refs.userForm) {
        this.$refs.userForm.resetFields();
      }
      this.formData = {
            name:"",
            Macs:[{name:""}],
            Products:[{name:""}],
            role:[{name:""}],
      };
    },
    async enterDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          let res;
          var item = []
          var list = []
          switch (this.type) {
            case "create":
              var re = await createUsers({...this.formData,roleid:Number(this.formData.roleid),Macs:[],Products:[]});
              if(re.code == 0){
                var User = await getLastUser()
                var lastUser = User.data.reUser
                for (let i = 0; i < this.formData.Macs.length; i++) {
                  item.push({
                    macId:Number(this.formData.Macs[i]),
                    userId:Number(lastUser.ID),
                    })
                }
                for (let i = 0; i < this.formData.Products.length; i++) {
                  list.push({
                    productId:Number(this.formData.Products[i]),
                    userId:Number(lastUser.ID),
                    })
                } 
                res = await createUserMac({item})
                res = await createUserProduct({list})
              }
              break;
            case "update":
              var um= await removeUserMacProduct({ID: this.formData.ID})
              if (um.code == 0 && this.formData.Macs != null) {
                for (let i = 0; i < this.formData.Macs.length; i++) {
                  item.push({
                    macId:Number(this.formData.Macs[i]),
                    userId:Number(this.formData.ID),
                    })
                  }
              createUserMac({item})
              }
              if (um.code == 0 && this.formData.Products != null) {
                for (let i = 0; i < this.formData.Products.length; i++) {
                  list.push({
                    productId:Number(this.formData.Products[i]),
                    userId:Number(this.formData.ID),
                    })
                  }
              createUserProduct({list})
              }
              res = await updateUsers({...this.formData,roleid:Number(this.formData.roleid),Macs:[],Products:[]});
              break;
            default:
              res = await createUsers({...this.formData,roleid:Number(this.formData.roleid),Macs:[],Products:[]});
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
          this.initForm();
        }
      });
    },
    async openDialog() {
      const product = await getProductList();
      this.productData = product.data.list;
      const role = await getRoleList();
      this.roleData = role.data.list;
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    const role = await getRoleList({ page: 1, pageSize: 999 });
    this.setRoleOptions(role.data.list);
    const product = await getProductList({ page: 1, pageSize: 999 });
    this.setProductOptions(product.data.list);
  
}
};
</script>

<style>
</style>