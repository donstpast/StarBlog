<template>
  <div class="contain">
    <div class="loginBox">
      <el-form
        ref="loginFormRef"
        :model="formData"
        :rules="rules"
        class="loginForm"
      >
        <el-form-item prop="username" style="margin-bottom: 30px">
          <el-input
            v-model="formData.username"
            type="text"
            placeholder="请输入用户名"
            @keydown.enter="login(loginFormRef)"
          >
            <template #prefix>
              <el-icon class="el-input__icon"><user /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formData.password"
            type="password"
            placeholder="请输入密码"
            show-password
            @keydown.enter="login(loginFormRef)"
          >
            <template #prefix>
              <el-icon class="el-input__icon"><key /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item class="formBtn">
          <el-button class="loginBtn" type="primary" @click="login(loginFormRef)" style="margin-right: 30px">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script lang="ts" setup>
import axios from 'axios'
import { reactive,ref } from "vue";
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from 'element-plus'
import router from "@/router";


const loginFormRef = ref<FormInstance>()
const formData = reactive({
  username: '',
  password: ''
})

const rules = reactive<FormRules>({
  username:[
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur'
    },
    {
      min: 1,
      max:40,
      message: '用户名长度为1-40个字符',
      trigger: 'blur'
    }
  ],
  password:[
    {
      required: true,
      message: '密码不能为空',
      trigger: 'blur'
    },
    {
      min: 6,
      max:40,
      message: '密码长度为6-40个字符',
      trigger: 'blur'
    }
  ]
})


const login = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const response = await axios.post('login', formData);
        if (response.data.status != 200)
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        sessionStorage.setItem('token',response.data.token)
        await router.push('/')

      } catch (error) {
        console.log(error); // 处理错误响应
      }
    } else {
      ElMessage({
        message: '请正确填写表单.',
        type: 'error',
        duration: 2000,
      });
    }
  });
};
</script>

<style scoped>
.contain {
  height: 100%;
  width: auto;
}
.loginBox {
  width: 450px;
  height: 300px;
  background-color: rgb(255, 255, 255, 0.5);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 20%;
  padding: 0 30px;
  box-sizing: border-box;
}
.formBtn {
  display: flex;
  justify-content: flex-end;
  margin-top: 30px;
  padding: 0 90px;
}
.loginBtn {
    display: flex;
    justify-content: flex-end;
    padding: 20px 90px;
}
</style>
