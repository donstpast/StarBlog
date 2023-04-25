<template>
  <div>
    <el-card class="cardBox">
      <!-- 搜索栏 -->
      <el-row :gutter="20">
        <el-col :span="7">
          <el-input
            v-model="searchData"
            @input="inputEmpty"
            @keydown.enter="clickSearch"
            :clearable="true"
            type="text"
            placeholder="请输入要查找的评论内容(支持模糊查询)"
          >
            <template #append>
              <el-button icon="Search" @click="clickSearch" />
            </template>
          </el-input>
        </el-col>
      </el-row>
      <!-- 评论列表 -->
      <el-table stripe style="width: 100%" class="commentTable" :data="commentList">
        <el-table-column prop="ID" label="ID" align="center" width="80" />
        <el-table-column prop="articleID" label="文章ID" align="center" width="80" />
        <el-table-column prop="uid" label="用户id" align="center" width="80"/>
        <el-table-column prop="nickName" label="用户昵称" align="center" width="80"/>
        <el-table-column prop="userEmail" label="评论者邮箱" align="center" width="220"/>
        <el-table-column prop="comment_content" label="评论内容" align="center" width="280"/>
        <el-table-column prop="UpdatedAt" label="评论时间" align="center" width="200">
          <template v-slot="scope">{{ dayjs(scope.row.UpdatedAt).format('YYYY-MM-DD HH:mm:ss') }}</template>
        </el-table-column>


        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button size="small" @click="editCommentBtn(scope.row)">Edit </el-button>
            <!-- 删除按钮 -->
            <el-button size="small" type="danger" @click="deleteComment(scope.row)">Delete </el-button>
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
    <!-- 编辑评论表单 -->
    <!-- Form -->
    <el-dialog
      v-model="editCommentFormVisible"
      title="编辑评论"
      width="30%"
      @close="clearForm(editCommentRef)"
    >
      <el-form
        :rules="editCommentRules"
        :model="editCommentForm"
        ref="editCommentRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="昵称" :label-width="formLabelWidth" prop="nickName">
          <el-input v-model="editCommentForm.nickName" :clearable="true" autocomplete="off" />
        </el-form-item>
        <el-form-item label="用户邮箱" :label-width="formLabelWidth" prop="userEmail">
          <el-input v-model="editCommentForm.userEmail" :clearable="true" autocomplete="off" />
        </el-form-item>
        <el-form-item label="评论内容" :label-width="formLabelWidth" prop="comment_content">
          <el-input
            v-model="editCommentForm.comment_content"
            :clearable="true"
            type="textarea"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editCommentFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="editCommentFormVisible;editComment(editCommentRef)"
          >
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage, ElPagination, ElMessageBox } from 'element-plus'
import axios from 'axios'
import { onMounted, ref, reactive } from 'vue'
import CommentList from "@/components/comment/comment-list.vue";
import dayjs from "dayjs";

const commentList = ref([]) //列表数据
const total = ref(0) //记录总数
const pageSize = ref(10) //分页大小
const currentPage = ref(1) //当前页
const searchData = ref('') //搜索数据
const editCommentFormVisible = ref(false) //编辑用户表单对话框

const formLabelWidth = '80px' //表单宽度
const editCommentRef = ref<FormInstance>()
const editCommentForm = reactive({
  id: '',
  comment_content : "",
  articleID: 0,
  uid: 0,
  nickName:"",
  userEmail:"",
})
//编辑评论采用的规则
const editCommentRules = reactive<FormRules>({
  comment_content: [
    {
      required: true,
      message: '评论内容不能为空',
      trigger: 'blur'
    },
  ],
  nickName: [
    {
      required: true,
      message: '昵称不能为空',
      trigger: 'blur'
    }
    ],
  userEmail: [
    {
      required: true,
      message: '邮箱不能为空',
      trigger: 'blur'
    },
    {
      //这里使用了一个正则表达式来判断邮箱的格式是否正确，
      // 正则表达式的含义是：首先匹配一个或多个字母、数字或下划线，
      // 然后可能跟着一个或多个由中横线、加号或句点连接的字母、数字或下划线，
      // 接着是一个@符号，然后匹配一个或多个字母、数字或中横线，
      // 接着是一个点号，最后是一个或多个字母、数字或中横线。
      pattern: /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
      message: '请输入正确的邮箱格式',
      trigger: 'blur'
    }
  ]


})

// 显示评论列表
const showComments = async () => {
  try {
    const response = await axios.get(`comments`, {
      params: { pagesize: pageSize.value, pagenum: currentPage.value }
    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    commentList.value = response.data.data
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

// 定义一个函数，用于检测搜索框是否为空
const inputEmpty = () => {
  if (searchData.value == '') {
    // 若搜索框为空，则调用 showComments() 函数展示所有用户信息
    showComments()
  }
}

// 定义一个函数，用于响应用户点击搜索按钮的事件
const clickSearch = () => {
  // 将当前页码设置为 1
  currentPage.value = 1
  // 调用 showComments() 函数展示符合搜索条件的用户信息
  showComments()
}

// 定义一个函数，用于处理评论编辑事件
const editCommentBtn = async (commentList: CommentList) => {
  // // 打印出被编辑的用户的索引和信息
  editCommentForm.id = commentList.ID
  editCommentFormVisible.value = true
  editCommentForm.nickName = commentList.nickName
  editCommentForm.comment_content = commentList.comment_content
  editCommentForm.userEmail = commentList.userEmail
}
const editComment = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        console.log(editCommentForm)
        const comment_id = parseInt(editCommentForm.id)
        const response = await axios.put(`comment/${comment_id}`, {
          comment_content: editCommentForm.comment_content,
          userEmail: editCommentForm.userEmail,
          uid: editCommentForm.uid || 0,
          nickName: editCommentForm.nickName
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          editCommentFormVisible.value = false
          await showComments()
          return ElMessage({
            message: '编辑用户成功',
            type: 'success'
          })
        }
      } catch (error) {
        console.log(error) // 处理错误响应
      }
    } else {
      ElMessage({
        message: '请正确填写表单.',
        type: 'error',
        duration: 2000
      })
    }
  })
}

// 定义一个异步函数，用于处理删除评论事件
const deleteComment = async (commentList: CommentList) => {
  try {
    // 使用 Element UI 的消息框组件，提示用户是否确定删除该用户信息
    const confirmResult = await ElMessageBox.confirm('确认要删除该评论么？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    // 若用户选择取消，则弹出提示消息并结束函数
    if (confirmResult != 'confirm') {
      ElMessage({
        message: '已取消',
        type: 'error'
      })
      return
    }
    // 发送删除该用户信息的请求
    const response = await axios.delete(`comment/${commentList.ID}`, {})
    // 调用 showUsers() 函数展示剩余用户信息
    await showComments()
    // 若删除操作未成功，则抛出错误
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
  } catch (error: unknown) {
    // 捕获错误，并根据错误类型进行不同的处理
    const err = error as Error
    if (typeof err === 'string' && err === 'cancel') {
      // 若用户取消删除操作，则弹出提示消息
      ElMessage({
        message: '已取消操作',
        type: 'warning'
      })
    } else {
      // 若删除操作出错，则弹出错误消息
      ElMessage({
        message: err.message,
        type: 'error'
      })
    }
  }
}

const clearForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl?.resetFields()
}

// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
  // 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
  // 调用 showComments() 函数展示符合当前搜索条件的用户信息
  showComments()
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
  // 将当前页码设置为用户选择的页码
  currentPage.value = val
  // 调用 showComments() 函数展示符合当前搜索条件的用户信息
  showComments()
}

// 在组件挂载完成后，调用 showComments() 函数展示所有用户信息
onMounted(() => {
  showComments()
})
</script>

<style scoped>
.commentTable {
    margin-top: 15px;
}
</style>
