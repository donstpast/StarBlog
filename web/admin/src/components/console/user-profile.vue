<template>
  <el-row :gutter="24">
    <el-col :span="24">
      <el-card class="topCard">
        test
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="bottomLeftCard">
        <a class="bottomCardTitle">个人资料</a>
        <el-form
          class="bottomLeftCardForm"
          :label-position="'left'"
          label-width="100px"
          :model="editProfileForm"
          style="max-width: 460px"
          size="large"
        >
          <el-form-item label="昵称" >
            <el-input v-model="editProfileForm.nickName" />
          </el-form-item>
          <el-form-item label="描述">
            <el-input v-model="editProfileForm.desc" />
          </el-form-item>
          <el-form-item label="头像">
            <el-input v-model="editProfileForm.avatar" />
          </el-form-item>
          <el-form-item label="主页地址">
            <el-input v-model="editProfileForm.site" />
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="bottomRightCard">
        账号安全
        <el-form
          class="bottomRightCardForm"
          :label-position="'left'"
          label-width="100px"
          :model="editProfileForm"
          style="max-width: 460px"
          size="large"
        >
          <el-form-item label="用户名" >
            <el-input v-model="editProfileForm.username" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="editProfileForm.email" />
          </el-form-item>
          <el-form-item label="修改密码">
            <el-input
              v-model="editProfileForm.password"
              type="password"
              autocomplete="off"
              show-password
            />
          </el-form-item>
          <el-form-item label="确认密码">
            <el-input
              v-model="editProfileForm.confirmPass"
              type="password"
              autocomplete="off"
              show-password
            />
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">

import { onMounted, reactive, ref } from "vue";
import axios from "axios";
import { ElMessage } from "element-plus";
const userList = ref([]) //列表数据

const editProfileForm = reactive({
  nickName:"",
  desc:"",
  avatar: "",
  site: "",
  username:"",
  email:"",
  password: '',
  confirmPass: '',
});

// 显示个人资料
const showProfile = async () => {
  try {
    const response = await axios.get(`user/profile`)
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    editProfileForm.nickName = response.data.data.nickName
    editProfileForm.avatar = response.data.data.avatar
    editProfileForm.desc = response.data.data.desc
    editProfileForm.site = response.data.data.site
    editProfileForm.username = response.data.data.username
    editProfileForm.email = response.data.data.email

  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}
onMounted(() =>{
  showProfile()
})
</script>

<style scoped>
.topCard {
    height: 150px;
}
.bottomCardTitle {
    font-size: 15px;
    font-weight: bold;
}
.bottomLeftCardForm {
    margin-top: 30px;
}
.bottomRightCardForm {
    margin-top: 30px;
}



</style>