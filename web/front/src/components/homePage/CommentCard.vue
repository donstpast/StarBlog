<template>
  <el-card class="commentCard box-card">
    <template #header>
      <div class="card-header">
        <span>最新评论</span>
      </div>
    </template>
    <div v-for="o in 10" :key="o" class="text item">{{ 'List item ' + o }}</div>
  </el-card>
</template>

<script setup lang="ts">
</script>

<style scoped>
.commentCard {
    height: auto;
    max-height: 600px;
}
</style>