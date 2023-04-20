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
        <el-table-column label="权限" prop="role" align="center" width="180">
          <template v-slot="{ row }">
            <el-tag v-if="row.role === 1" type="warning">管理员</el-tag>
            <el-tag v-if="row.role === 2" type="success">普通用户</el-tag>
          </template>
        </el-table-column>

<!--        <el-table-column prop="role" label="权限" align="center" width="180" :formatter="formatRole" />-->
        <!-- 操作列 -->
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <!-- 编辑按钮 -->
            <el-button
              size="small"
              @click="editUserBtn(scope.row)"
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
    <el-dialog v-model="addUserFormVisible" title="新增用户" width="30%" @close="clearForm(addUserRef)">
      <el-form
        :rules="addUserRules"
        :model="addUserForm"
        ref="addUserRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="用户名" :label-width="formLabelWidth" prop="username">
          <el-input v-model="addUserForm.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth" prop="password">
          <el-input
            v-model="addUserForm.password"
            type="password"
            autocomplete="off"
            show-password
          />
        </el-form-item>
        <el-form-item label="确认密码" :label-width="formLabelWidth" prop="confirmPass">
          <el-input
            v-model="addUserForm.confirmPass"
            type="password"
            autocomplete="off"
            show-password
          />
        </el-form-item>
        <el-form-item label="权限" :label-width="formLabelWidth" prop="role">
          <el-select
            v-model="addUserForm.role"
            @change="roleChange"
          >
            <el-option label="普通用户" value="2"  />
            <el-option label="管理员" value="1" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="addUserFormVisible = false" >Cancel</el-button>
        <el-button type="primary" @click="addUserFormVisible;addUser(addUserRef)">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
<!-- 编辑用户表单 -->
    <!-- Form -->
    <el-dialog v-model="editUserFormVisible" title="编辑用户" width="30%" @close="clearForm(editUserRef)">
      <el-form
        :rules="editUserRules"
        :model="editUserForm"
        ref="editUserRef"
        hide-required-asterisk
        status-icon
      >
        <el-form-item label="用户名" :label-width="formLabelWidth" prop="username">
          <el-input
            v-model="editUserForm.username"
            :clearable = true
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth" prop="password">
          <el-input
            v-model="editUserForm.password"
            :clearable = true
            type="password"
            autocomplete="off"
            show-password
          />
        </el-form-item>
        <el-form-item label="权限" :label-width="formLabelWidth" prop="role">
          <el-select
            v-model="editUserForm.role"
            @change="roleChange"
          >
            <el-option label="普通用户" value="2"  />
            <el-option label="管理员" value="1" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="editUserFormVisible = false" >Cancel</el-button>
        <el-button type="primary" @click="editUserFormVisible;editUser(userList,editUserRef)">
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
import UserList from "@/components/user/user-list.vue";

const userList = ref([]); //列表数据
const total = ref(0); //记录总数
const pageSize = ref(10); //分页大小
const currentPage = ref(1); //当前页
const searchData = ref('');   //搜索数据
const addUserFormVisible = ref(false) //添加用户表单对话框
const editUserFormVisible = ref(false) //编辑用户表单对话框

const formLabelWidth = '80px'  //表单宽度
const addUserRef = ref<FormInstance>()
const editUserRef = ref<FormInstance>()
const addUserForm = reactive({
  username: '',
  password: '',
  confirmPass:'',
  role:'2',
})
const editUserForm = reactive({
  id : '',
  username: '',
  password: '',
  role:'',
})
//新增用户采用的规则
const addUserRules = reactive<FormRules>({
  username:[
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur'
    },
    {
      min: 1,
      max:40,
      message: '用户名长度为1-40个字符',
      trigger: 'blur'
    }
  ],
  password:[
    {
      required: true,
      validator: (rule: any, value: any, callback: any) => {
        if (value === '') {
          callback(new Error('密码不能为空'))
        } else {
          if (addUserForm.confirmPass !== '') {
            if (!addUserRef.value) return
            addUserRef.value.validateField('confirmPass', () => null)
          }
          callback()
        }
      },
      trigger: 'blur'
    },
    {
      min: 6,
        max:40,
      message: '密码长度为6-40个字符',
      trigger: 'blur'
    }
  ],
  confirmPass:[
    {
      required: true,
      validator: (rule: any, value: any, callback: any) => {
        if (value === '') {
          callback(new Error('请再次输入密码'))
        } else if (value !== addUserForm.password) {
          callback(new Error("两次密码输入不一致，请检查"))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
})
//编辑用户采用的规则
const editUserRules = reactive<FormRules>({
  username:[
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur'
    },
    {
      min: 1,
      max:40,
      message: '用户名长度为1-40个字符',
      trigger: 'blur'
    }
  ],
  password:[
    {
      validator: (rule: any, value: any, callback: any) => {
        if (value !== '') {
          if ([...editUserForm.password].length < 6 || [...editUserForm.password].length > 40) {
            callback(new Error('密码长度为6-10个字符'))
          } else {
            callback()
          }
        }
          callback()
      },
      trigger: 'blur'
    }
  ]
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

// // 格式化role列，如果用户role为1则显示为管理员，否则显示为普通用户
// const formatRole = (row: any) => {
//   if (row.role === 1) {
//     return '管理员';
//   } else {
//     return '普通用户';
//   }
// };
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
const editUserBtn = async (userList: UserList) => {
// // 打印出被编辑的用户的索引和信息
  editUserFormVisible.value = true;
  editUserForm.id = userList.ID
  editUserForm.username = userList.username;
  editUserForm.role = userList.role.toString()
}
const editUser = async (userList: UserList,formEl: FormInstance | undefined) => {
  console.log(formEl)
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const user_id = parseInt(editUserForm.id)
        console.log(user_id)
        const response  = await axios.put(`user/${user_id}`,{
          username:editUserForm.username,
          password:editUserForm.password || '',
          role:parseInt(editUserForm.role),
        });
        if (response.data.status != 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          editUserFormVisible.value = false
          await showUsers();
          return ElMessage({
            message:"编辑用户成功",
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

const clearForm = (formEl: FormInstance | undefined)=> {
  if (!formEl) return;
  formEl?.resetFields()
}
//新增用户表单提交
const addUser = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const response = await axios.post('user/add', {
          username:addUserForm.username,
          password:addUserForm.password,
          role:parseInt(addUserForm.role),
        });
        if (response.data.status != 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          addUserFormVisible.value = false
          await showUsers();
          return ElMessage({
            message:"新增用户成功",
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