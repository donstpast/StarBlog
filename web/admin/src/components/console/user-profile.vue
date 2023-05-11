<template>
  <el-row :gutter="24">
    <el-col :span="24">
      <el-card class="topCard">
        <el-row :gutter="24">
          <el-col :span="2">
            <el-avatar
              class="Avatar"
              :size="100"
              :src=editProfileForm.avatar
            />
          </el-col>
          <el-col :span="4">
              <div class="text-box">
                <a>Hello,{{editProfileForm.nickName}}</a>
                <br>
                <span>这是一个用来占位的测试文本</span>
              </div>
            </el-col>
          <el-col :offset="15" :span="3">
            <el-button type="primary" class="saveButton"><el-icon class="buttonIcon"><Finished /></el-icon>保存信息</el-button>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="bottomLeftCard bottomCard">
        <a class="bottomCardTitle">个人资料</a>
        <el-form
          :model="editProfileForm"
          :rules="editProfileRules"
          ref="editProfileRef"
          class="bottomLeftCardForm"
          :label-position="'left'"
          label-width="100px"
          hide-required-asterisk

          style="max-width: 460px"
          size="large"
        >
          <el-form-item label="昵称" prop="nickName" >
            <el-input v-model="editProfileForm.nickName" />
          </el-form-item>
          <el-form-item label="描述" prop="desc">
            <el-input v-model="editProfileForm.desc" />
          </el-form-item>
          <el-form-item label="头像" prop="avatar">
            <el-input v-model="editProfileForm.avatar" />
          </el-form-item>
          <el-form-item label="主页地址" prop="site">
            <el-input v-model="editProfileForm.site" />
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="bottomRightCard bottomCard">
        账号安全
        <el-form
          :rules="editProfileRules"
          :model="editProfileForm"
          ref="editProfileRef"
          class="bottomRightCardForm"
          :label-position="'left'"
          label-width="100px"
          hide-required-asterisk

          style="max-width: 460px"
          size="large"
        >
          <el-form-item label="用户名" prop="username">
            <el-input v-model="editProfileForm.username" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="editProfileForm.email" />
          </el-form-item>
          <el-form-item label="旧密码" prop="oldPass">
            <el-input
              v-model="editProfileForm.oldPass"
              type="password"
              autocomplete="off"
              show-password
              placeholder="请输入您原先的密码"
            />
          </el-form-item>
          <el-form-item v-if="showPasswordFields" label="修改密码" prop="password">
            <el-input
              v-model="editProfileForm.password"
              type="password"
              autocomplete="off"
              show-password
              placeholder="请输入新密码"
            />
          </el-form-item>
          <el-form-item v-if="showConfirmPassField" label="确认密码" prop="confirmPass">
            <el-input
              v-model="editProfileForm.confirmPass"
              type="password"
              autocomplete="off"
              show-password
              placeholder="请重新输入新密码"
            />
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">

import { onMounted, reactive, ref, watch } from "vue";
import axios from "axios";
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from "element-plus";
import UserList from "@/components/user/user-list.vue";
const editProfileRef = ref<FormInstance>()
// 控制是否显示密码字段和确认密码字段
const showPasswordFields = ref(false);
const showConfirmPassField = ref(false);
const editProfileForm = reactive({
  nickName:"",
  desc:"",
  avatar: "",
  site: "",
  username:"",
  email:"",
  oldPass:"",
  password: '',
  confirmPass: '',
});

//编辑个人资料采用的规则
const editProfileRules = reactive<FormRules>({
  username: [
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur'
    },
    {
      min: 1,
      max: 40,
      message: '用户名长度为1-40个字符',
      trigger: 'blur'
    }
  ],
  oldPass: [
    {
      validator: (rule: any, value: any, callback: any) => {
        if (value === '') {
// 旧密码为空时，隐藏密码字段和确认密码字段
          showPasswordFields.value = false;
          showConfirmPassField.value = false;
          callback();
        } else if (value.length >= 6 && value.length <= 40) {
          showPasswordFields.value = true;
          callback();
        } else {
          showPasswordFields.value = false;
          showConfirmPassField.value = false;
          callback(new Error('密码长度为6-40个字符'));
        }
      },
      trigger: 'blur'
    },
    {
      min: 6,
      max: 40,
      message: '密码长度为6-40个字符',
      trigger: 'blur',
      ...(editProfileForm.oldPass !== '' && {
        required: true,
        message: '密码不能为空',
        trigger: 'blur'
      })
    }
  ],
  password: [
    {
      validator: (rule: any, value: any, callback: any) => {
        if (value === '') {
// 密码为空时，隐藏确认密码字段
          showConfirmPassField.value = false;
          callback();
        } else if (value.length >= 6 && value.length <= 40) {
          showConfirmPassField.value = true;
          callback();
        } else {
          showConfirmPassField.value = false;
          callback(new Error('密码长度为6-40个字符'));
        }
      },
      trigger: 'blur',
      //如果editProfileForm的password属性不为空，则将大括号中的属性包含在新对象中。否则，排除它们。
      ...(editProfileForm.oldPass !== '' && {
        required: true,
        message: '密码不能为空',
        trigger: 'blur'
      })
    }
  ],
  confirmPass: [
    {
      validator: (rule: any, value: any, callback: any) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== editProfileForm.password) {
          callback(new Error('两次密码输入不一致，请检查'));
        } else {
          callback();
        }
      },
      trigger: 'blur',
      ...(editProfileForm.password !== '' && {
        required: true,
        message: '请重新输入新密码',
        trigger: 'blur'
      })
    }
  ],
  email: [
    {
      required: true,
      message: '邮箱不能为空',
      trigger: 'blur'
    },
    {
      type: 'email',
      message: '请输入有效的邮箱地址',
      trigger: ['blur', 'change']
    }
  ],

  site: [
    {
      type: 'url',
      message: '请输入有效的网站地址',
      trigger: ['blur', 'change']
    }
  ]
});


// 显示个人资料
const showProfile = async () => {
  try {
    const response = await axios.get(`user/profile`)
    if (response.data.status !== 200) {
      throw new Error(response.data.message)
    }
    editProfileForm.nickName = response.data.data.nickName
    editProfileForm.avatar = response.data.data.avatar
    editProfileForm.desc = response.data.data.desc
    editProfileForm.site = response.data.data.site
    editProfileForm.username = response.data.data.username
    editProfileForm.email = response.data.data.email

  } catch (error: unknown) {
    const err = error as Error
    // 显示错误信息
    ElMessage({
      message: err.message,
      type: 'error'
    })
  }
}

//编辑个人资料表单提交
const editProfile = async (userList: UserList, formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        const user_id = parseInt(editUserForm.id)
        console.log(user_id)
        const response = await axios.put(`user/${user_id}`, {
          username: editUserForm.username,
          password: editUserForm.password || '',
          role: parseInt(editUserForm.role)
        })
        if (response.data.status != 200) {
          return ElMessage({
            message: response.data.message,
            type: 'error'
          })
        } else {
          editUserFormVisible.value = false
          await showUsers()
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



// 监听旧密码字段的变化
watch(() => editProfileForm.oldPass, (newVal) =>{
  if (newVal !== '') {
    if (newVal.length >= 6 && newVal.length <= 40) {
      showPasswordFields.value = true;
      if (editProfileForm.password !== '') {
        if (editProfileForm.confirmPass !== '') {
          showConfirmPassField.value = true;
        } else {
          showConfirmPassField.value = false;
        }
      } else {
        showConfirmPassField.value = false;
      }
    } else {
      showPasswordFields.value = false;
      showConfirmPassField.value = false;
    }
  } else {
    showPasswordFields.value = false;
    showConfirmPassField.value = false;
  }
});

// 监听修改密码字段的变化
watch(() => editProfileForm.password, (newVal) => {
  if (newVal !== '') {
    if (editProfileForm.oldPass !== '' && newVal.length >= 6 && newVal.length <= 40) {
      if (editProfileForm.confirmPass === '') {
        showConfirmPassField.value = true;
      } else {
        editProfileForm.confirmPass = ''
        showConfirmPassField.value = false;
      }
    } else {
      showConfirmPassField.value = false;
    }
  } else {
    showConfirmPassField.value = false;
  }
});

// 监听确认密码字段的变化
watch(() => editProfileForm.confirmPass, (newVal) => {
  if (newVal !== '') {
    if (
      editProfileForm.oldPass !== '' && editProfileForm.password !== ''
    ) {
      showConfirmPassField.value = true;
    } else {
      showConfirmPassField.value = false;
    }
  } else {
    if (editProfileForm.oldPass !== '' && editProfileForm.password !== ''){
      showConfirmPassField.value = true;
    } else {
      showConfirmPassField.value = false;
    }
  }
});

onMounted(() =>{
  showProfile()
})
</script>

<style scoped>
.topCard {
    height: 140px;
}
.Avatar {
    border: 1px solid rgba(223,226,230,0.1);
}
.text-box {
    margin: 10px 0 0 20px;
    width: 400px;
}
.text-box a {
    font-size: 20px;
}
.text-box span {
    font-size: 10px;
    background-image: linear-gradient(to right, red, green,blue);
    -webkit-background-clip: text; /* Safari/Chrome浏览器上需要添加此属性 */
    color: transparent;
}
.buttonIcon {
    padding-right: 5px;
}

.bottomCardTitle {
    font-size: 15px;
    font-weight: bold;
}
.bottomLeftCardForm {
    margin-top: 30px;
}
.bottomRightCardForm {
    margin-top: 30px;
}
.bottomCard{
    margin-top: 30px;
}

</style>