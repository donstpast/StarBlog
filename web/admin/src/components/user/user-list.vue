<template>
  <div>
    <el-card class="cardBox">
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
          <el-button type="primary">
            新增
          </el-button>
        </el-col>
      </el-row>
      <el-table
        stripe
        style="width: 100%"
        class="userTable"
        :data="userList"
      >
        <el-table-column prop="ID" label="ID" align="center" width="180" />
        <el-table-column prop="username" label="用户名" align="center" width="180" />
        <el-table-column prop="role" label="权限" align="center" width="180" :formatter="formatRole" />
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <el-button
              size="small"
              @click="handleEdit(scope.$index, scope.row)"
            >Edit
            </el-button>
            <el-button
              size="small"
              type="danger"
              @click="deleteUser(scope.row)"
            >Delete
            </el-button>
          </template>
        </el-table-column>
    </el-table>
      <el-pagination
        layout="prev, pager, next, jumper"
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
import { ElMessage, ElPagination, ElMessageBox} from "element-plus";
import axios from "axios";
import { onMounted, ref } from "vue";
import UserList from "@/components/user/user-list.vue";

const userList = ref([]);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const searchData = ref('')

//显示用户列表
const showUsers = async () => {
  try {
    const response = await axios.get("user/list", {
      params: { username : searchData.value, pagesize: pageSize.value, pagenum: currentPage.value },
    });
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    userList.value = response.data.data;
    total.value = response.data.TotalNum;
  } catch (error: unknown) {
    const err = error as Error;
    ElMessage({
      message: err.message,
      type: "error",
    });
  }
};
//格式化role列，如果用户role为1则显示为管理员，否则显示为普通用户
const formatRole = (row: any) => {
  if (row.role === 1) {
    return '管理员';
  } else {
    return '普通用户';
  }
};
const inputEmpty = () => {
  if (searchData.value==""){
    showUsers()
  }
}
const clickSearch = () => {
  showUsers()
}
const handleEdit = (index: number, row: UserList) => {
  console.log(index, row)
}
const deleteUser = async (userList: UserList) => {
  try {
    const confirmResult = await ElMessageBox.confirm(
      "确认要删除该用户信息么？",
      "提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    );
    if (confirmResult != "confirm") {
      ElMessage({
        message: "已取消",
        type: "error",
      })
      return;
    }
    const response = await axios.delete(`user/${userList.ID} `, {
    });
    await showUsers();
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
  } catch (error: unknown) {
    const err = error as Error;
    if (typeof err === "string" && err === "cancel"){
      ElMessage({
        message: "已取消操作",
        type: "warning",
      });
    }else {
      ElMessage({
        message: err.message,
        type: "error",
      });
    }

  }
};
const handleSizeChange = (val: number) => {
  pageSize.value = val
  showUsers();
}
const handleCurrentChange = (val: number) => {
  currentPage.value = val;
  showUsers();
};
onMounted(() => {
  showUsers();
});


</script>


<style scoped>
.userTable {
     margin-top: 15px;
}

</style>