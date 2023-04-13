<template>
  <div>
    <el-card class="cardBox">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-input placeholder="请输入要查找的用户名" >
            <template #append>
              <el-button icon="Search"/>
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
        <el-table-column prop="ID" label="ID" width="180" />
        <el-table-column prop="username" label="用户名" width="180" />
        <el-table-column prop="role" label="权限" :formatter="formatRole"/>
    </el-table>
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </el-card>

  </div>

</template>

<script lang="ts" setup>
import { ElMessage, ElPagination } from "element-plus";
import axios from "axios";
import { onMounted, ref } from "vue";

const userList = ref([]);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const token = sessionStorage.getItem("token");

const showUsers = async () => {
  try {
    const response = await axios.get("http://127.0.0.1:8383/api/v1/user/list", {
      headers: { Authorization: `Bearer ${token}` },
      params: { pagesize: pageSize.value, pagenum: currentPage.value },
    });
    if (response.data.status !== 200) {
      throw new Error(response.data.message);
    }
    currentPage.value = response.data.currentPage;
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
const handleSizeChange = (val: number) => {
  pageSize.value = val
  showUsers();

}
const handleCurrentChange = async (val: number) => {
  currentPage.value = val;
  await showUsers();
};

const formatRole = (row: any) => {
  if (row.role === 1) {
    return '管理员';
  } else if (row.role === 2) {
    return '普通用户';
  } else {
    return '';
  }
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