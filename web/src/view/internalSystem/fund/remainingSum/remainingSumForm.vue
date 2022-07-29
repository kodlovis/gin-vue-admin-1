<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
          <el-form-item label="填写日期:">
          <el-date-picker v-model="formData.upDate" type="date" placeholder="选择日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="formData.type" placeholder="请选择类型">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="余额:">
          <el-input v-model.number="formData.inValue" clearable placeholder="请输入余额"></el-input>
        </el-form-item>
        <el-form-item label="平台:">
          <el-input v-model.number="formData.company" disabled></el-input>
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model.number="formData.createUser" disabled></el-input>
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
    createRemainingSum,
    updateRemainingSum,
    findRemainingSum
} from "@/api/internalSystem/fund/remainingSum";   
import infoList from "@/mixins/infoList";
export default {
  name: "RemainingSum",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
        upDate: new Date(),
        type: '资金余额',
        inValue: 0,

      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createRemainingSum(this.formData);
          break;
        case "update":
          res = await updateRemainingSum(this.formData);
          break;
        default:
          res = await createRemainingSum(this.formData);
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
    const res = await findRemainingSum({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reremainingSum
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