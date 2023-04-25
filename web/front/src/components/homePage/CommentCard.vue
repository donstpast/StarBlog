<template>
  <el-card class="commentCard box-card">
    <template #header>
      <div class="card-header">
        <span>最新评论</span>
      </div>
    </template>
    <div
      v-for="{ID,comment_content,nickName,UpdatedAt} in commentList"
      :key="ID"
      class="comment"
    ><div class="comment-info">
          {{ nickName }} • {{ formatDate(UpdatedAt) }}
        </div>
        <div class="comment-content">
          {{ comment_content }}
        </div>
        <el-divider border-style="dashed" />
      </div>
  </el-card>
</template>

<script setup lang="ts">

import { onMounted, ref } from "vue";
import axios from "axios";
import { ElMessage } from "element-plus";
const pageSize = ref(10) //分页大小
const currentPage = ref(1) //当前页
const commentList = ref([])
// 显示最新的十条评论
const showComment = async () => {
  try {
    const response = await axios.get(`comments`,{
      params: { pagesize: pageSize.value, pagenum: currentPage.value }

    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    commentList.value = response.data.data
    console.log(commentList)
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}
const formatDate = (date:any) => {
  const now = new Date();
  const commentDate = new Date(date);
  const diff = Math.floor((now.getTime() - commentDate.getTime()) / (1000 * 60 * 60 * 24));
  if (diff === 0) {
    return '今天';
  } else {
    return `${diff}天前`;
  }
}

// 在组件挂载完成后，调用 showComment()) 函数展示所有分类
onMounted(() => {
  showComment()
})
</script>

<style scoped>
.commentCard {
    height: auto;
    max-height: 600px;
}
.comment {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px; /* 两个评论之间的间隔 */
}

.comment-info {
    margin-bottom: 10px; /* 名字和时间戳之间的间隔 */
    color: rgb(238,121,124);

}

.comment-content {
    word-wrap: break-word; /* 换行 */
    word-break: break-all; /* 省略号 */
}
</style>