<template>
  <div>
    <el-card class="cardBox">
      <!-- 搜索栏 -->
      <el-row :gutter="20">
        <el-col :span="4">
          <!-- 新增按钮 -->
          <el-button type="primary" @click="addTagFormVisible = true"> 新增标签 </el-button>
        </el-col>
      </el-row>
      <!-- 标签列表 -->
      <el-table stripe style="width: 100%" class="tagTable" :data="tagList">
        <el-table-column prop="id" label="ID" align="center" width="180" />
        <el-table-column prop="name" label="分类名" align="center" width="180" />
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button size="small" @click="editTagBtn(scope.row)">Edit </el-button>
            <!-- 删除按钮 -->
            <el-button size="small" type="danger" @click="deleteTag(scope.row)"
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
    <!-- 新增标签表单 -->
    <!-- Form -->
    <el-dialog
      v-model="addTagFormVisible"
      title="新增标签"
      width="30%"
      @close="clearForm(addTagRef)"
    >
      <el-form
        :rules="addTagRules"
        :model="addTagForm"
        ref="addTagRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="标签名" :label-width="formLabelWidth" prop="name">
          <el-input v-model="addTagForm.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="addTagFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="
              addTagFormVisible;
              addTag(addTagRef)
            "
          >
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 编辑标签表单 -->
    <!-- Form -->
    <el-dialog
      v-model="editTagFormVisible"
      title="编辑标签"
      width="30%"
      @close="clearForm(editTagRef)"
    >
      <el-form
        :rules="editTagRules"
        :model="editTagForm"
        ref="editTagRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="标签名" :label-width="formLabelWidth" prop="name">
          <el-input v-model="editTagForm.name" :clearable="true" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editTagFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="
              editTagFormVisible;
              editTag(tagList, editTagRef)
            "
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
import TagList from "@/components/tag/tag-list.vue";

const tagList = ref([]) //列表数据
const total = ref(0) //记录总数
const pageSize = ref(10) //分页大小
const currentPage = ref(1) //当前页
const addTagFormVisible = ref(false) //添加标签表单对话框
const editTagFormVisible = ref(false) //编辑标签表单对话框

const formLabelWidth = '80px' //表单宽度
const addTagRef = ref<FormInstance>()
const editTagRef = ref<FormInstance>()
const addTagForm = reactive({
  id: 0,
  name: ''
})
const editTagForm = reactive({
  id: 0,
  name: ''
})
//新增标签采用的规则
const addTagRules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '标签名不能为空',
      trigger: 'blur'
    }
  ]
})
//编辑标签采用的规则
const editTagRules = reactive<FormRules>({
  username: [
    {
      required: true,
      message: '标签不能为空',
      trigger: 'blur'
    }
  ]
})

// 显示标签列表
const showTags = async () => {
  try {
    const response = await axios.get(`tags`, {
      params: { pagesize: pageSize.value, pagenum: currentPage.value }
    })
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    tagList.value = response.data.data
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

// 定义一个函数，用于处理用户编辑事件
const editTagBtn = async (tagList: TagList) => {
  // // 打印出被编辑的用户的索引和信息
  editTagFormVisible.value = true
  editTagForm.id = tagList.id
  editTagForm.name = tagList.name
}
const editTag = async (tagList: TagList, formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const tag_id = editTagForm.id
        const response = await axios.put(`category/${tag_id}`, {
          name: editTagForm.name
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          editTagFormVisible.value = false
          await showTags()
          return ElMessage({
            message: '编辑标签成功',
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

// 定义一个异步函数，用于处理删除标签事件
const deleteTag = async (tagList: TagList) => {
  try {
    // 使用 Element UI 的消息框组件，提示用户是否确定删除该分类信息
    const confirmResult = await ElMessageBox.confirm('确认要删除该标签信息么？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    // 若选择取消，则弹出提示消息并结束函数
    if (confirmResult != 'confirm') {
      ElMessage({
        message: '已取消',
        type: 'error'
      })
      return
    }
    // 发送删除该标签信息的请求
    const response = await axios.delete(`category/${tagList.id}`, {})
    // 调用 showTags() 函数展示剩余标签信息
    await showTags()
    // 若删除操作未成功，则抛出错误
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
  } catch (error: unknown) {
    // 捕获错误，并根据错误类型进行不同的处理
    const err = error as Error
    if (typeof err === 'string' && err === 'cancel') {
      // 若取消删除操作，则弹出提示消息
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
//新增标签表单提交
const addTag = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const response = await axios.post('tag/add', {
          name: addTagForm.name
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          addTagFormVisible.value = false
          await showTags()
          return ElMessage({
            message: '新增标签成功',
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

// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
  // 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
  // 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showTags()
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
  // 将当前页码设置为用户选择的页码
  currentPage.value = val
  // 调用 showTags() 函数展示符合当前搜索条件的用户信息
  showTags()
}

// 在组件挂载完成后，调用 showTags() 函数展示所有用户信息
onMounted(() => {
  showTags()
})
</script>

<style scoped>
.tagTable {
    margin-top: 15px;
}
</style>
