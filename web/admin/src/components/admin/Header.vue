<template>
  <div class="header">
    <el-button type="danger" @click="logout" round>退出</el-button>
  </div>
</template>

<script lang="ts" setup>
import router from "@/router";
const logout = () => {
  sessionStorage.clear()
  router.push("/login")
}
</script>

<style scoped>
</style>
