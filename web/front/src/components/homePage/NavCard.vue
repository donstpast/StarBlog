<template>
  <el-card
    class="box-card blogInfo"
    :body-style="cardBodyStyle"
  >
      <div class="avatar-circle">
        <div class="avatar-container">
          <el-avatar
            class="blogAvatar"
            :size="70"
            src="https://www.donstpast.cn/about/img/touxiang.jpg"
          />

        </div>
      </div>
    <div class="name"></div>
    <div class="articleData"></div>


  </el-card>

  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>页面导航</span>

      </div>
    </template>

    <div
      v-for="item in categoryList"
      :key="item.id"
    >{{ item.name }}</div>
  </el-card>
    <el-card class="friendLinkCard box-card">
      <template #header>
        <div class="card-header">
          <span>友情链接</span>
        </div>
      </template>
      <div
        v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
    </el-card>

</template>

<script setup lang="ts">
// 显示分类列表
import axios from "axios";
import { onMounted, ref } from "vue";
import { ElMessage } from "element-plus";
const categoryList = ref([])
const cardBodyStyle:any = {
  padding: '0px',
};

const showCategories = async () => {
  try {
    const response = await axios.get(`categories`)
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    categoryList.value = response.data.data
    console.log(categoryList)
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
})
</script>

<style scoped>
.box-card {
}

.avatar-circle {
    background-image: linear-gradient(to bottom, rgba(255, 0, 149, 0.2), rgba(0, 247, 255, 0.2) 65%, rgba(250,251,254,0.9) 45%);
    width: 100%;
    height: 100px;
    overflow: hidden;
    margin-bottom: 30px;
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
}

.blogAvatar:hover {
    transform: rotate(360deg) scale(1.2); /* 旋转一圈并且变大 */
}






</style>