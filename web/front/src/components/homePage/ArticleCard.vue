<template>
  <el-card class="top-card box-card">
    <div>最近文章 • 第 {{currentPage}} 页</div>
  </el-card>
<el-card
  class="box-card"
  v-for="article in articleList"
  :key="article.id"
>
<template #header>
  <div class="card-header">
    <h3>{{ article.title }}</h3>
    <el-button
      class="button"
      text
      @click="showArticleDetailBtn(article)"
    >阅读全文</el-button>
  </div>
</template>
<div class="text item">{{ article.desc }}</div>
</el-card>
  <!-- 分页栏 -->
  <el-pagination
    background
    layout="total, prev, pager, next, jumper"
    v-model:current-page="currentPage"
    v-model:page-size="pageSize"
    :total="total"
    @size-change="handleSizeChange(pageSize)"
    @current-change="handleCurrentChange(currentPage)"
  />
</template>

<script setup lang="ts">
import { ElMessage, ElPagination, ElMessageBox } from 'element-plus'
import axios from 'axios'
import { onMounted, provide, ref } from "vue";
import router from "@/router";
const articleList = ref([]) //列表数据
const total = ref(0) //记录总数
const pageSize = ref(5) //分页大小
const currentPage = ref(1) //当前页
provide('currentPage', currentPage)


// 显示文章列表
const showArticles = async () => {
  try {
    const response = await axios.get(`articles`, {
      params: { pagesize: pageSize.value, pagenum: currentPage.value }
    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    articleList.value = response.data.data
    total.value = response.data.TotalNum
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}
 const showArticleDetailBtn = async (article :any) => {
  await router.push(`/article/${article.ID}/`)
 }


// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
  // 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
  // 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showArticles()
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
  // 将当前页码设置为用户选择的页码
  currentPage.value = val
  // 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showArticles()
}
// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息
onMounted(() => {
  showArticles()
})
</script>

<style scoped>
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    height: 250px;
}
.top-card {
    height: 60px;
}




</style>