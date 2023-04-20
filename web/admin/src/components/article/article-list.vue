<template>
  <div>
    <el-card class="cardBox">
      <!-- 搜索栏 -->
      <el-row :gutter="20">
        <el-col :span="6">
          <el-input
            v-model="searchData"
            @input="inputEmpty"
            @keydown.enter="clickSearch"
            :clearable = true
            type="text"
            placeholder="请输入要查找的文章标题"
          >
            <template #append>
              <el-button
                icon="Search"
                @click="clickSearch"
              />
            </template>
          </el-input>
        </el-col>
        <el-col :span="4">
          <!-- 新增按钮 -->
          <el-button
            type="primary"
            @click="addArticleBtn()"
          >
            撰写文章
          </el-button>
        </el-col>
        <el-col :span="6" :offset="8" >
          <el-select
            v-model="selectValue"
            clearable
            default-first-option
            placeholder="所有分类"
            :reserve-keyword="false"
            @change="cateChange"
          >
            <el-option
              v-for="item in categoryList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            ></el-option>
          </el-select>
        </el-col>

      </el-row>
      <!-- 文章列表 -->
      <el-table
        stripe
        style="width: 100%"
        class="articleTable"
        :data="articleList"
      >
        <el-table-column prop="ID" label="ID" align="center" width="100" />
        <el-table-column prop="User.username" label="作者" align="center" width="180" />
        <el-table-column prop="title" label="标题" align="center" width="380" />
        <el-table-column prop="Category.name" label="分类" align="center" width="150" />
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button
              size="small"
              @click="editArticleBtn(scope.row)"
            >Edit
            </el-button>
            <!-- 删除按钮 -->
            <el-button
              size="small"
              type="danger"
              @click="deleteArticle(scope.row)"
            >Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页栏 -->
      <el-pagination
        layout="total, prev, pager, next, jumper"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        @size-change="handleSizeChange(pageSize)"
        @current-change="handleCurrentChange(currentPage)"
      />
    </el-card>
  </div>

</template>

<script lang="ts" setup>
import { ElMessage, ElPagination, ElMessageBox } from "element-plus";
import axios from "axios";
import { onMounted, ref ,reactive} from "vue";
import ArticleList from "@/components/article/article-list.vue";
import router from "@/router";
const articleList = ref([]); //列表数据
const categoryList = ref([]);
const total = ref(0); //记录总数
const pageSize = ref(10); //分页大小
const currentPage = ref(1); //当前页
const searchData = ref('');   //搜索数据
const selectValue = ref('')

// 显示文章列表
const showArticles = async () => {
  try {
    const response = await axios.get( `articles`, {
      params: { title : searchData.value, pagesize: pageSize.value, pagenum: currentPage.value },
    });
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    articleList.value = response.data.data;
    total.value = response.data.TotalNum;
  } catch (error: unknown) {
    const err = error as Error;
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: "error",
    });
  }
};

// 显示分类列表
const showCategories = async () => {
  try {
    const response = await axios.get( `categories`);
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    categoryList.value = response.data.data;
  } catch (error: unknown) {
    const err = error as Error;
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: "error",
    });
  }
};
//切换分类选择器
const cateChange = async (category_id:any) => {
  try {
    if (category_id === ''){
      await showArticles()
    }else {
      const response = await axios.get( `category/articles/${category_id}`);
      if (response.data.status !== 200) {
        throw new Error(response.data.message);
      }
      articleList.value = response.data.data;
      total.value = response.data.TotalNum;
    }
    } catch (error: unknown) {
    const err = error as Error;
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: "error",
    });
  }
}
// 定义一个函数，用于检测搜索框是否为空
const inputEmpty = () => {
  if (searchData.value==""){
// 若搜索框为空，则调用 showUsers() 函数展示所有用户信息
    showArticles()
  }
}

// 定义一个函数，用于响应用户点击搜索按钮的事件
const clickSearch = () => {
// 将当前页码设置为 1
  currentPage.value = 1;
// 调用 showUsers() 函数展示符合搜索条件的用户信息
  showArticles()
}
// 定义一个函数，用于处理用户编辑事件
const editArticleBtn = async (articleList: ArticleList) => {
// // 打印出被编辑的用户的索引和信息
  const article_id = articleList.ID
  await router.push(`/admin/write-article/${article_id}/`)
}
const addArticleBtn = async () => {
  await router.push(`/admin/write-article`)
}

// 定义一个异步函数，用于处理删除文章事件
const deleteArticle = async (articleList: ArticleList) => {
  try {
// 使用 Element UI 的消息框组件，提示用户是否确定删除该用户信息
    const confirmResult = await ElMessageBox.confirm(
      "确认要删除该文章信息么？",
      "提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    );
// 若选择取消，则弹出提示消息并结束函数
    if (confirmResult != "confirm") {
      ElMessage({
        message: "已取消",
        type: "error",
      })
      return;
    }
// 发送删除该文章信息的请求
    const response = await axios.delete(`article/${articleList.ID}` , {
    });
// 调用 showUsers() 函数展示剩余文章信息
    await showArticles();
// 若删除操作未成功，则抛出错误
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
  } catch (error: unknown) {
// 捕获错误，并根据错误类型进行不同的处理
    const err = error as Error;
    if (typeof err === "string" && err === "cancel"){
// 若取消删除操作，则弹出提示消息
      ElMessage({
        message: "已取消操作",
        type: "warning",
      });
    }else {
// 若删除操作出错，则弹出错误消息
      ElMessage({
        message: err.message,
        type: "error",
      });
    }
  }
};

// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
// 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showArticles();
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
// 将当前页码设置为用户选择的页码
  currentPage.value = val;
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showArticles();
};

// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息,调用showCategory显示所有分类
onMounted(() => {
  showArticles();
  showCategories();
});


</script>


<style scoped>
.articleTable {
    margin-top: 15px;
}

</style>