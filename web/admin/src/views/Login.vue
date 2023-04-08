import type { ElForm } from 'element-plus';

<template>
  <div class="contain">
    <div class="loginbox">
      <el-form class="loginform">
        <el-form-item style="margin-bottom: 30px">
          <el-input v-model="username" type="text" placeholder="请输入用户名" >
            <template #prefix>
              <el-icon class="el-input__icon"><user /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="password" type="password" placeholder="请输入密码" show-password>
            <template #prefix>
              <el-icon class="el-input__icon"><key /></el-icon>
            </template>
          </el-input>



        </el-form-item>
        <el-form-item class="loginbtn">
          <el-button type="primary" style="margin-right: 30px">Login</el-button>
          <el-button style="margin-left: 30px">Cancel</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<style scoped>
.contain {
  height: 100%;
  width: auto;
}
.loginbox {
  width: 450px;
  height: 300px;
  background-color: rgb(255, 255, 255, 0.5);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.loginform {
  width: 100%;
  position: absolute;
  bottom: 20%;
  padding: 0 30px;
  box-sizing: border-box;
}
.loginbtn {
  display: flex;
  justify-content: flex-end;
  margin-top: 30px;
  padding: 0 90px;
}
</style>
<script lang="ts" setup>
import { ref } from 'vue'
const username = ref('')
const password = ref('')

</script>
