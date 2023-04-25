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
            placeholder="请输入要查找的友链标题(支持模糊查询)"
          >
            <template #append>
              <el-button icon="Search" @click="clickSearch" />
            </template>
          </el-input>
        </el-col>
        <el-col :span="4">
          <!-- 新增按钮 -->
          <el-button type="primary" @click="addFriendFormVisible = true"> 新增友链 </el-button>
        </el-col>
      </el-row>
      <!-- 友链列表 -->
      <el-table stripe style="width: 100%" class="commentTable" :data="friendsList">
        <el-table-column prop="id" label="ID" align="center" width="80" />
        <el-table-column prop="name" label="友链名称" align="center" width="280" />
        <el-table-column prop="link" label="友链网址" align="center" width="280"/>
        <el-table-column prop="avatar" label="头像" align="center" width="250">
          <template v-slot:="scope">
            <img :src="scope.row.avatar" class="avatarImg" >
          </template>
        </el-table-column>
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button size="small" @click="editFriendBtn(scope.row)">Edit </el-button>
            <!-- 删除按钮 -->
            <el-button size="small" type="danger" @click="deleteFriend(scope.row)">Delete </el-button>
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
    <!-- 新增用户表单 -->
    <!-- Form -->
    <el-dialog
      v-model="addFriendFormVisible"
      title="新增友链"
      width="30%"
      @close="clearForm(addFriendRef)"
    >
      <el-form
        :rules="addFriendRules"
        :model="addFriendForm"
        ref="addFriendRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="友链名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="addFriendForm.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="友链网址" :label-width="formLabelWidth" prop="link">
          <el-input v-model="addFriendForm.link" autocomplete="off" />
        </el-form-item>
        <el-form-item label="友链头像" :label-width="formLabelWidth" prop="avatar">
          <el-input v-model="addFriendForm.avatar" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="addFriendFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="addFriendFormVisible;addFriend(addFriendRef)"
          >
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 编辑友链表单 -->
    <!-- Form -->
    <el-dialog
      v-model="editFriendFormVisible"
      title="编辑用户"
      width="30%"
      @close="clearForm(editFriendRef)"
    >
      <el-form
        :rules="editFriendRules"
        :model="editFriendForm"
        ref="editFriendRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="友链名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="editFriendForm.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="友链网址" :label-width="formLabelWidth" prop="link">
          <el-input v-model="editFriendForm.link" autocomplete="off" />
        </el-form-item>
        <el-form-item label="友链头像" :label-width="formLabelWidth" prop="avatar">
          <el-input v-model="editFriendForm.avatar" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editFriendFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="editFriendFormVisible;editFriend(editFriendRef)"
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
import FriendsLinks from "@/components/friends/friends-links.vue";

const friendsList = ref([]) //列表数据
const total = ref(0) //记录总数
const pageSize = ref(10) //分页大小
const currentPage = ref(1) //当前页
const searchData = ref('') //搜索数据
const editFriendFormVisible = ref(false) //编辑用户表单对话框
const addFriendFormVisible = ref(false) //添加用户表单对话框
const formLabelWidth = '80px' //表单宽度
const addFriendRef = ref<FormInstance>()
const editFriendRef = ref<FormInstance>()

const addFriendForm = reactive({
  name: '',
  link: '',
  avatar: '',
})
const editFriendForm = reactive({
  id: '',
  name:'',
  link:'',
  avatar:''
})

//新增友链采用的规则
const addFriendRules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '标题不能为空',
      trigger: 'blur'
    },
  ],
  link: [
    {
      required: true,
      message:'请输入友链地址',
      trigger:'blur'
    },
    {
      pattern: /^https?:\/\/\S+$/,
      message: '请输入正确的网址格式',
      trigger: 'blur'
    }
  ],
  avatar:[
    {
      pattern: /^https?:\/\/\S+$/,
      message: '请输入正确的网址格式',
      trigger: 'blur'
    }
  ]
})
//编辑友链采用的规则
const editFriendRules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '标题不能为空',
      trigger: 'blur'
    },
  ],
  link: [
    {
      required: true,
      message:'请输入友链地址',
      trigger:'blur'
    },
    {
      pattern: /^https?:\/\/\S+$/,
      message: '请输入正确的网址格式',
      trigger: 'blur'
    }
  ],
  avatar:[
    {
      pattern: /^https?:\/\/\S+$/,
      message: '请输入正确的网址格式',
      trigger: 'blur'
    }
  ]
})

// 显示友链列表
const showFriends = async () => {
  try {
    const response = await axios.get(`friends`, {
      params: { pagesize: pageSize.value, pagenum: currentPage.value }
    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    friendsList.value = response.data.data
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
    showFriends()
  }
}

// 定义一个函数，用于响应用户点击搜索按钮的事件
const clickSearch = () => {
  // 将当前页码设置为 1
  currentPage.value = 1
  // 调用 showComments() 函数展示符合搜索条件的用户信息
  showFriends()
}


//新增友链表单提交
const addFriend = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const response = await axios.post('friend/add', {
          name: addFriendForm.name,
          link: addFriendForm.link,
          avatar: addFriendForm.avatar
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          addFriendFormVisible.value = false
          await showFriends()
          return ElMessage({
            message: '新增友链成功',
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

// 定义一个函数，用于处理评论编辑事件
const editFriendBtn = async (friendsList: FriendsLinks) => {
  // // 打印出被编辑的用户的索引和信息
  editFriendFormVisible.value = true
  editFriendForm.id = friendsList.id
  editFriendForm.name = friendsList.name
  editFriendForm.link = friendsList.link
  editFriendForm.avatar = friendsList.avatar
}
const editFriend = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const friendLink_id = editFriendForm.id
        const response = await axios.put(`friend/${friendLink_id}`, {
          name: editFriendForm.name,
          link: editFriendForm.link,
          avatar: editFriendForm.avatar,
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          editFriendFormVisible.value = false
          await showFriends()
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
const deleteFriend = async (friendsLinks: FriendsLinks) => {
  try {
    // 使用 Element UI 的消息框组件，提示用户是否确定删除该用户信息
    const confirmResult = await ElMessageBox.confirm('确认要删除该友链么？', '提示', {
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
    const response = await axios.delete(`friend/${friendsLinks.ID}`, {})
    // 调用 showUsers() 函数展示剩余用户信息
    await showFriends()
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
  showFriends()
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
  // 将当前页码设置为用户选择的页码
  currentPage.value = val
  // 调用 showComments() 函数展示符合当前搜索条件的用户信息
  showFriends()
}

// 在组件挂载完成后，调用 showComments() 函数展示所有用户信息
onMounted(() => {
  showFriends()
})
</script>

<style scoped>
.commentTable {
    margin-top: 15px;
}
.avatarImg{
    width:60px;
    height:60px;
    border-radius: 10px;
}

</style>
