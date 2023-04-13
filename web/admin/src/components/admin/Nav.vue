<template>
  <div class="logo" >
      <span>{{ 'StarBolg Admin' }}</span>
  </div>

  <el-menu
    class="slideMenu"
    active-text-color="#ffd04b"
    background-color="#343A44"
    :unique-opened="true"
    text-color=rgb(134,145,160)
    @open="handleOpen"
    @close="handleClose"
    :default-active="activeIndex"
    :router="true"
  >
    <el-sub-menu index="1" >
      <template #title>
        <el-icon><SetUp /></el-icon>
        <span>控制台</span>
      </template>
      <el-menu-item-group title="概要">
        <el-menu-item index="1-1">博客数据</el-menu-item>
        <el-menu-item index="1-2">评论数据</el-menu-item>
        <el-menu-item index="1-3">访客信息</el-menu-item>
      </el-menu-item-group>
      <el-menu-item-group title="个人设置">
        <el-menu-item index="1-4">个人资料</el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>
    <el-sub-menu index="2" >
      <template #title>
        <el-icon><Document /></el-icon>
        <span>撰写</span>
      </template>
      <el-menu-item index="write-article" >编写文章</el-menu-item>
      <el-menu-item index="create-page">创建页面</el-menu-item>
    </el-sub-menu>
    <el-sub-menu index="3" >
      <template #title>
        <el-icon><Star /></el-icon>
        <span>管理</span>
      </template>
      <el-menu-item index="Article" >文章</el-menu-item>
      <el-menu-item index="3-2" >独立页面</el-menu-item>
      <el-menu-item index="3-3">评论</el-menu-item>
      <el-menu-item index="Category">分类</el-menu-item>
      <el-menu-item index="3-5">标签</el-menu-item>
      <el-menu-item index="User">用户</el-menu-item>
      <el-menu-item index="3-7">友链</el-menu-item>
      <el-menu-item index="3-8">文件</el-menu-item>
    </el-sub-menu>
    <el-sub-menu index="4" >
      <template #title>
        <el-icon><setting /></el-icon>
        <span>设置</span>
      </template>
      <el-menu-item index="4-1">站点设置</el-menu-item>
      <el-menu-item index="4-2">评论</el-menu-item>
      <el-menu-item index="4-3">阅读</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import router from "@/router";

const activeIndex = ref( '1-1')

const handleOpen = (key: string, keyPath: string[]) => {
  console.log('当前展开的菜单项：', keyPath[0]);
};

const handleClose = (key: string, keyPath: string[]) => {
  console.log("当前折叠的菜单项：", keyPath[0]);

};

onMounted(() => {
  const targetPath = router.currentRoute.value.path.slice(1);
  if (targetPath.length === 0) {
    activeIndex.value = "1-1"
  }else {
    activeIndex.value = String(targetPath)
  }
})


</script>


<style scoped>
.logo {
  color: rgb(250, 251, 254);
  height: 32px;
  margin: 16px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
.slideMenu {
    border-right: none;
}
</style>
