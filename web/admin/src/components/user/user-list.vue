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
            placeholder="请输入要查找的用户名"
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
          <el-button type="primary" @click="addUserFormVisible = true">
            新增
          </el-button>
        </el-col>
      </el-row>
      <!-- 用户列表 -->
      <el-table
        stripe
        style="width: 100%"
        class="userTable"
        :data="userList"
      >
        <el-table-column prop="ID" label="ID" align="center" width="180" />
        <el-table-column prop="username" label="用户名" align="center" width="180" />
        <el-table-column prop="role" label="权限" align="center" width="180" :formatter="formatRole" />
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button
              size="small"
              @click="handleEdit(scope.$index, scope.row)"
            >Edit
            </el-button>
            <!-- 删除按钮 -->
            <el-button
              size="small"
              type="danger"
              @click="deleteUser(scope.row)"
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
<!-- 新增用户表单 -->
    <!-- Form -->
    <el-dialog v-model="addUserFormVisible" title="新增用户" width="30%">
      <el-form :model="addUserForm">
        <el-form-item label="用户名" :label-width="formLabelWidth">
          <el-input v-model="addUserForm.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth">
          <el-input
            v-model="addUserForm.password"
            type="password"
            autocomplete="off"
            show-password
          />
        </el-form-item>
        <el-form-item label="确认密码" :label-width="formLabelWidth">
          <el-input
            v-model="addUserForm.confirmPass"
            type="password"
            autocomplete="off"
            show-password
          />
        </el-form-item>
        <el-form-item label="权限" :label-width="formLabelWidth">
          <el-select
            v-model="addUserForm.role"
            default-first-option
            @change="roleChange"
          >
            <el-option label="普通用户" value="2"  />
            <el-option label="管理员" value="1" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="addUserFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="addUserFormVisible = false">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>

</template>

<script lang="ts" setup>
import { ElMessage, ElPagination, ElMessageBox} from "element-plus";
import axios from "axios";
import { onMounted, ref ,reactive} from "vue";
import UserList from "@/components/user/user-list.vue";

const userList = ref([]); //列表数据
const total = ref(0); //记录总数
const pageSize = ref(10); //分页大小
const currentPage = ref(1); //当前页数
const searchData = ref('');   //搜索数据
const addUserFormVisible = ref(false) //添加用户表单对话框
const formLabelWidth = '80px'  //表单宽度
const addUserForm = reactive({
  name: '',
  password: '',
  confirmPass:"",
  role:"",
})

// 显示用户列表
const showUsers = async () => {
  try {
    const response = await axios.get( `user/list`, {
      params: { username : searchData.value, pagesize: pageSize.value, pagenum: currentPage.value },
    });
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    userList.value = response.data.data;
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

// 格式化role列，如果用户role为1则显示为管理员，否则显示为普通用户
const formatRole = (row: any) => {
  if (row.role === 1) {
    return '管理员';
  } else {
    return '普通用户';
  }
};
// 定义一个函数，用于检测搜索框是否为空
const inputEmpty = () => {
  if (searchData.value==""){
// 若搜索框为空，则调用 showUsers() 函数展示所有用户信息
    showUsers()
  }
}

// 定义一个函数，用于响应用户点击搜索按钮的事件
const clickSearch = () => {
// 将当前页码设置为 1
  currentPage.value = 1;
// 调用 showUsers() 函数展示符合搜索条件的用户信息
  showUsers()
}
const roleChange = (value :any) => {
  console.log(value)
}

// 定义一个函数，用于处理用户编辑事件
const handleEdit = (index: number, row: UserList) => {
// 打印出被编辑的用户的索引和信息
  console.log(index, row)
}

// 定义一个异步函数，用于处理删除用户事件
const deleteUser = async (userList: UserList) => {
  try {
// 使用 Element UI 的消息框组件，提示用户是否确定删除该用户信息
    const confirmResult = await ElMessageBox.confirm(
      "确认要删除该用户信息么？",
      "提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    );
// 若用户选择取消，则弹出提示消息并结束函数
    if (confirmResult != "confirm") {
      ElMessage({
        message: "已取消",
        type: "error",
      })
      return;
    }
// 发送删除该用户信息的请求
    const response = await axios.delete(`user/${userList.ID}` , {
    });
// 调用 showUsers() 函数展示剩余用户信息
    await showUsers();
// 若删除操作未成功，则抛出错误
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
  } catch (error: unknown) {
// 捕获错误，并根据错误类型进行不同的处理
    const err = error as Error;
    if (typeof err === "string" && err === "cancel"){
// 若用户取消删除操作，则弹出提示消息
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
//新增用户表单




// 定义一个函数，用于响应用户选择每页显示条数的事件
const handleSizeChange = (val: number) => {
// 将当前页显示的条数设置为用户选择的条数
  pageSize.value = val
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showUsers();
}

// 定义一个函数，用于响应用户选择页码的事件
const handleCurrentChange = (val: number) => {
// 将当前页码设置为用户选择的页码
  currentPage.value = val;
// 调用 showUsers() 函数展示符合当前搜索条件的用户信息
  showUsers();
};

// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息
onMounted(() => {
  showUsers();
});


</script>


<style scoped>
.userTable {
     margin-top: 15px;
}

</style>