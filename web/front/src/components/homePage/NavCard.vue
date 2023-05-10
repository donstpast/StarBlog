<template>
  <el-card
    class="box-card blogInfo"
    :body-style="cardBodyStyle"
  >
      <div class="avatar-circle">
        <div class="avatar-container">
          <a href="/"> <el-avatar
            class="blogAvatar"
            :size="70"
            src="https://www.donstpast.cn/about/img/touxiang.jpg"
          /></a>
          <div class="name" ><a href="/"> Donstpast</a></div>
        </div>
      </div>
    <el-row :gutter="20" class="totalData">
      <el-col :span="6" :offset="3"><div class="totalNum" >{{totalArticle}}</div>文章</el-col>
      <el-col :span="6"><div class="totalNum">{{totalComment}}</div> 评论</el-col>
      <el-col :span="6"><div class="totalNum">{{totalCategory}}</div> 分类</el-col>
    </el-row>
  </el-card>

  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>页面导航</span>
      </div>
    </template>
    <div
      v-for="{id,name} in categoryList"
      :key="id"
      class="category-list"
    >
      <a :href="'category/' + name">
        {{ name }}
      </a>
    </div>
  </el-card>
  <el-card class="friendLinkCard box-card">
    <template #header>
      <div class="card-header">
        <span>友情链接</span>
      </div>
    </template>
    <div
      v-for="{avatar, id, link, name} in friendList"
      :key="id"
      class="friend-link"
    >
      <a :href="link">
        <img :src="avatar" alt="" class="friend-link-avatar">
        {{ name }}

      </a>
    </div>
  </el-card>


</template>

<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { ElMessage } from "element-plus";
const categoryList = ref([])
const friendList = ref([])

const totalArticle = ref(0) //记录总数
const totalCategory = ref(0) //记录总数
const totalComment = ref(0) //记录总数
const cardBodyStyle:any = {
  padding: '0px',
};

// 显示分类列表
const showCategories = async () => {
  try {
    const response = await axios.get(`categories`)
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    totalCategory.value = response.data.TotalNum
    categoryList.value = response.data.data
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}
// 显示友链列表
const showFriends = async () => {
  try {
    const response = await axios.get(`friends`)
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    friendList.value = response.data.data
    console.log(friendList)
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}
const GetTotalNum = async () => {
  try {
    const articleResponse = await axios.get(`articles`)
    const commentResponse = await axios.get(`comments`)
    if (articleResponse.data.status !== 200) {
      throw new Error(articleResponse.data.message)
    } else if (commentResponse.data.status !== 200){
      throw new Error(commentResponse.data.message)
    }
    totalArticle.value = articleResponse.data.TotalNum
    totalComment.value = commentResponse.data.TotalNum
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}

// 在组件挂载完成后，调用 showCategories()) 函数展示所有分类
onMounted(() => {
  showCategories()
  showFriends()
  GetTotalNum()
})
</script>

<style scoped>
.box-card {
}

.avatar-circle {
    background-image: linear-gradient(to bottom, rgba(255, 0, 149, 0.2), rgba(0, 247, 255, 0.2) 55%, rgba(250,251,254,0.9) 45%);
    width: 100%;
    height: 120px;
    overflow: hidden;
    margin-bottom: 20px;
    position: relative;
}

.avatar-container {
    position: absolute;
    top: 55%;
    left: 50%;
    transform: translate(-50%, -50%);
}
.blogAvatar {
    transition: transform 0.3s ease-in-out; /* 添加过渡效果 */
    margin-top: 10px;
    margin-bottom: 10px;
}

.blogAvatar:hover {
    transform: rotate(360deg) scale(1.2); /* 旋转一圈并且变大 */
}
.name a{
    text-decoration: none;
    color: black;
}
.totalData{
    margin-bottom: 20px;
}
.totalNum{
    padding-left: 10px;
    color: rgb(63,137,194);
}

.category-list{
    border-bottom: 1px dashed #ccc;
    text-align: center;
    margin: 10px;
    padding-bottom: 10px;

}
.category-list a {
    text-decoration: none;
    color: black;
}
.category-list a:hover{
    color: rgb(238,121,124);
}

.friend-link-avatar {
    width: 30px;
    height: 30px;
    margin-right: 10px;
    border-radius: 50%;
    object-fit: cover;
}

.friend-link {
    border-bottom: 1px dashed #ccc;
    text-align: center;
    margin: 10px;
}

.friend-link a {
    display: flex;
    align-items: center;
    text-decoration: none;
    color: black;
}
.friend-link a:hover{
    color: rgb(238,121,124);
}




</style>