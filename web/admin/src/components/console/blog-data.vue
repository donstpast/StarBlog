<template>
  <el-row :gutter="24">
    <el-col :span="6">
      <el-card class="visitorCountCard topCard">
        <span>友链总数</span>
        <el-icon class="topCardIcon"><Avatar /></el-icon>
        <div class="topCardContent">
          <a href="friends-links">{{ totalFriends }} 个友链</a>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="articleCountCard topCard">
        <span>文章总计</span>
        <el-icon class="topCardIcon"><DocumentCopy /></el-icon>
        <div class="topCardContent">
          <a href="Article">{{ totalArticle }} 篇文章</a>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="commentCountCard topCard">
        <span>评论总计</span>
        <el-icon class="topCardIcon"><Comment /></el-icon>
        <div class="topCardContent">
          <a href="Comment">{{ totalComment }} 条评论</a>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="categoryCountCard topCard">
        <span>分类总计</span>
        <el-icon class="topCardIcon"><Orange /></el-icon>
        <div class="topCardContent">
          <a href="Category">{{ totalCategory }} 个分类</a>
        </div>
      </el-card>
    </el-col>
  </el-row>
  <el-row :gutter="24">
    <el-col :span="12">
      <el-card class="bottomCard">
        <a class="bottomCardTitle">最近发布的文章</a>
        <el-table :data="articleList" v-for="{UpdatedAt,title,ID} in articleList" style="width: 100%" class="bottomCardContent">
          <el-table-column  label="日期" width="180" align="center">
            {{ UpdatedAt.slice(0,10) }}
          </el-table-column>
          <el-table-column  label="文章标题" width="380" align="center">
            <a :href="'http://localhost:5174/article/'+ ID" class="linkContent">{{ title }}</a>
          </el-table-column>
        </el-table>

      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="bottomCard">
        <a class="bottomCardTitle">最近的收到的回复</a>
        <el-table :data="commentList" v-for="{UpdatedAt,nickName,comment_content,userEmail} in commentList" style="width: 100%" class="bottomCardContent">
          <el-table-column label="日期" width="180" align="center">
            {{ UpdatedAt.slice(0,10) }}
          </el-table-column>
          <el-table-column label="昵称及评论内容" width="380" align="center">
            <a :href="'mailto:'+userEmail" class="linkContent"> {{nickName}}</a>: {{comment_content}}
          </el-table-column>
        </el-table>
      </el-card>

    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { ElMessage } from "element-plus";
import { update } from "lodash-es";

const totalArticle = ref(0) //记录总数
const totalCategory = ref(0) //记录总数
const totalComment = ref(0) //记录总数
const totalFriends = ref(0) //友链总数

const articleList = ref([]) //列表数据
const commentList = ref([])
const pageSize = ref(10) //分页大小
const currentPage = ref(1) //当前页
const GetTotalNum = async () => {
  try {
    const commentResponse = await axios.get(`comments`)
    const categoryResponse = await axios.get(`categories`)
    const friendResponse = await axios.get(`friends`)
    if (commentResponse.data.status !== 200){
      throw new Error(commentResponse.data.message)
    } else if (categoryResponse.data.status !== 200){
      throw new Error(categoryResponse.data.message)
    }else if (friendResponse.data.status !== 200){
      throw new Error(friendResponse.data.message)
    }
    totalComment.value = commentResponse.data.TotalNum
    totalCategory.value = categoryResponse.data.TotalNum
    totalFriends.value = friendResponse.data.TotalNum
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}

// 显示文章列表
const showArticles = async () => {
  try {
    const response = await axios.get(`articles`, {
      params: {  pagesize: pageSize.value, pagenum: currentPage.value }
    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    articleList.value = response.data.data
    totalArticle.value = response.data.TotalNum
  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}

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


onMounted(() => {
  GetTotalNum()
  showArticles()
  showComment()
})

</script>

<style scoped>
.topCard{
    height: 140px;
    border-radius: 9px;
    color: white;
}
.visitorCountCard {
    background-image: linear-gradient(to right, rgb(242,191,150),rgb(254,112,150));
}
.articleCountCard {
    background-image: linear-gradient(to right, rgb(144,202,249),rgb(4,126,223));
}
.commentCountCard {
    background-image: linear-gradient(to right, rgb(132,217,210),rgb(7,205,174));
}
.categoryCountCard {
    background-image: linear-gradient(to right, rgb(238, 124, 208), rgb(217, 48, 183));
}
.topCardContent{
    padding-top: 10px;
    font-size: 26px;
}
.topCardContent a {
    color: white;
    text-decoration: none;
}
.topCardContent a:hover{
    text-decoration:underline;

}
.topCardIcon {
    font-size: 25px;
    padding-left: 160px;
}
.bottomCard {
    margin-top: 50px;
    height: 600px;

}
.bottomCardTitle {
    font-size: 18px;
    font-weight: bold;
}
.bottomCardContent {
    margin-top: 15px;
}
.linkContent {
    text-decoration: none;
    color: rgb(81,139,247);
}
.linkContent:hover {
    text-decoration:underline;
    color: rgb(95,132,196);
}

</style>
