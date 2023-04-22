<template>
  <div>
    <el-header class="header">
      <el-menu
        class="navBar"
        :default-active="$route.path"
        mode="horizontal"
        active-text-color="#54a0ff"
        :ellipsis = false
      >
        <el-menu-item index="/" class="item">
          <span>首页</span>
        </el-menu-item>
        <el-menu-item index="/about">
          <span>About Me</span>
        </el-menu-item>
      </el-menu>
      <el-avatar class="rightHeader" src="https://www.donstpast.cn/about/img/touxiang.jpg"></el-avatar>
    </el-header>
  </div>

</template>

<script setup lang="ts">


</script>

<style scoped>
.header{
    background: rgba(255, 255, 255, 0.5);
    display: flex;
    align-items: center;
}
.navBar {
    width: 50%;
    display: flex;
    justify-content: flex-start;
    padding-left: 10px;
    background-color: transparent;
    border-bottom: none;
}
.rightHeader {
    margin-left: 45%;
    display: flex;
    justify-content: flex-end;
}

</style>
