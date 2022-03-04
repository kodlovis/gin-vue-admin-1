<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="marketDay字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.marketDay" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="futuresVariety字段:">
                <el-input v-model="formData.futuresVariety" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="warehouseReceipt字段:">
                  <el-input-number v-model="formData.warehouseReceipt" :precision="2" clearable></el-input-number>
           </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createUs004FutureStockLme,
    updateUs004FutureStockLme,
    findUs004FutureStockLme
} from "@/api/us004FutureStockLme";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Us004FutureStockLme",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            marketDay:new Date(),
            futuresVariety:"",
            warehouseReceipt:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createUs004FutureStockLme(this.formData);
          break;
        case "update":
          res = await updateUs004FutureStockLme(this.formData);
          break;
        default:
          res = await createUs004FutureStockLme(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findUs004FutureStockLme({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reus004FutureStockLme
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>