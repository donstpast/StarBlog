<template>
  <el-menu class="header" background-color="#00000000" mode="horizontal" :ellipsis="false">
    <el-sub-menu index="">
      <template #title>账户</template>
      <el-menu-item index="0-1">个人资料</el-menu-item>
      <el-menu-item index="0-2" @click="writeArticle">撰写文章</el-menu-item>
      <el-menu-item index="0-3">系统配置</el-menu-item>
      <el-menu-item index="0-3" @click="logout">退出登录</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" setup>
import router from '@/router'
const logout = () => {
  sessionStorage.removeItem('token')
  router.push('/login')
}
const writeArticle = () => {
  router.push('/admin/write-article')
}
</script>

<style scoped>
.header {
  border-bottom: none;
}
</style>
