<template>
  <div>
    <el-card class="cardBox">
      <!-- 搜索栏 -->
      <el-row :gutter="20">

        <el-col :span="4">
          <!-- 新增按钮 -->
          <el-button type="primary" @click="addCategoryFormVisible = true">
            新增分类
          </el-button>
        </el-col>
      </el-row>
      <!-- 分类列表 -->
      <el-table
        stripe
        style="width: 100%"
        class="categoryTable"
        :data="categoryList"
      >
        <el-table-column prop="id" label="ID" align="center" width="180" />
        <el-table-column prop="name" label="分类名" align="center" width="180" />
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button
              size="small"
              @click="editCategoryBtn(scope.row)"
            >Edit
            </el-button>
            <!-- 删除按钮 -->
            <el-button
              size="small"
              type="danger"
              @click="deleteCategory(scope.row)"
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
    <!-- 新增分类表单 -->
    <!-- Form -->
    <el-dialog v-model="addCategoryFormVisible" title="新增分类" width="30%" @close="clearForm(addCategoryRef)">
      <el-form
        :rules="addCategoryRules"
        :model="addCategoryForm"
        ref="addCategoryRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="分类名" :label-width="formLabelWidth" prop="name">
          <el-input v-model="addCategoryForm.name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="addCategoryFormVisible = false" >Cancel</el-button>
        <el-button type="primary" @click="addCategoryFormVisible;addCategory(addCategoryRef)">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
    <!-- 编辑分类表单 -->
    <!-- Form -->
    <el-dialog v-model="editCategoryFormVisible" title="编辑分类" width="30%" @close="clearForm(editCategoryRef)">
      <el-form
        :rules="editCategoryRules"
        :model="editCategoryForm"
        ref="editCategoryRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="分类名" :label-width="formLabelWidth" prop="name">
          <el-input
            v-model="editCategoryForm.name"
            :clearable = true
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="editCategoryFormVisible = false" >Cancel</el-button>
        <el-button type="primary" @click="editCategoryFormVisible;editCategory(categoryList,editCategoryRef)">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>

</template>

<script lang="ts" setup>
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage, ElPagination, ElMessageBox } from "element-plus";
import axios from "axios";
import { onMounted, ref ,reactive} from "vue";
import CategoryList from "@/components/category/category-list.vue";

const categoryList = ref([]); //列表数据
const total = ref(0); //记录总数
const pageSize = ref(10); //分页大小
const currentPage = ref(1); //当前页
const addCategoryFormVisible = ref(false) //添加用户表单对话框
const editCategoryFormVisible = ref(false) //编辑用户表单对话框

const formLabelWidth = '80px'  //表单宽度
const addCategoryRef = ref<FormInstance>()
const editCategoryRef = ref<FormInstance>()
const addCategoryForm = reactive({
  id:0,
  name: ''
})
const editCategoryForm = reactive({
  id : 0,
  name: ''
})
//新增用户采用的规则
const addCategoryRules = reactive<FormRules>({
  name:[
    {
      required: true,
      message: '分类名不能为空',
      trigger: 'blur'
    }
    ]
})
//编辑用户采用的规则
const editCategoryRules = reactive<FormRules>({
  username:[
    {
      required: true,
      message: '分类不能为空',
      trigger: 'blur'
    }
  ]
})

// 显示分类列表
const showCategories = async () => {
  try {
    const response = await axios.get( `categories`, {
      params: { pagesize: pageSize.value, pagenum: currentPage.value },
    });
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    categoryList.value = response.data.data;
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


// 定义一个函数，用于处理用户编辑事件
const editCategoryBtn = async (userList: CategoryList) => {
// // 打印出被编辑的用户的索引和信息
  editCategoryFormVisible.value = true;
  editCategoryForm.id = userList.id
  editCategoryForm.name = userList.name;
}
const editCategory = async (userList: CategoryList,formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const category_id = editCategoryForm.id
        const response  = await axios.put(`category/${category_id}`,{
          name:editCategoryForm.name,
        });
        if (response.data.status != 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          editCategoryFormVisible.value = false
          await showCategories();
          return ElMessage({
            message:"编辑分类成功",
            type:"success",
          })
        }
      } catch (error) {
        console.log(error); // 处理错误响应
      }
    } else {
      ElMessage({
        message: '请正确填写表单.',
        type: 'error',
        duration: 2000,
      });
    }
  });
};

// 定义一个异步函数，用于处理删除分类事件
const deleteCategory = async (categoryList: CategoryList) => {
  try {
// 使用 Element UI 的消息框组件，提示用户是否确定删除该分类信息
    const confirmResult = await ElMessageBox.confirm(
      "确认要删除该分类信息么？",
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
// 发送删除该分类信息的请求
    const response = await axios.delete(`category/${categoryList.id}` , {
    });
// 调用 showUsers() 函数展示剩余分类信息
    await showCategories();
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

const clearForm = (formEl: FormInstance | undefined)=> {
  if (!formEl) return;
  formEl?.resetFields()
}
//新增分类表单提交
const addCategory = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const response = await axios.post('category/add', {
          name:addCategoryForm.name
        });
        if (response.data.status != 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          addCategoryFormVisible.value = false
          await showCategories();
          return ElMessage({
            message:"新增分类成功",
            type:"success",
          })
        }
      } catch (error) {
        console.log(error); // 处理错误响应
      }
    } else {
      ElMessage({
        message: '请正确填写表单.',
        type: 'error',
        duration: 2000,
      });
    }
  });
};

// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
// 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showCategories();
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
// 将当前页码设置为用户选择的页码
  currentPage.value = val;
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showCategories();
};

// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息
onMounted(() => {
  showCategories();
});


</script>


<style scoped>
.categoryTable {
    margin-top: 15px;
}

</style>