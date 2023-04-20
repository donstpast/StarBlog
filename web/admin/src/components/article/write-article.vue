<template>
  <div>
    <el-card class="cardBox">
      <el-form
        ref="articleFormRef"
        :model="formData"
        :rules="rules"
        class="articleForm"
      >
        <el-row :gutter="20">
          <el-col :span="16">
            <el-form-item>
              <el-input
                v-model="formData.title"
                :clearable = true
                type="text"
                placeholder="输入文章标题"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <!-- 发布按钮 -->
            <el-button type="primary" @click="postArticleBtn(articleFormRef)">
              发布文章
            </el-button>
          </el-col>
        </el-row>
        <el-form-item label="分类：">
          <el-select
            v-model="formData.cid"
            clearable
            default-first-option
            placeholder="选择文章分类"
            :reserve-keyword="false"
            @change="cateChange"
          >
            <el-option
              v-for="item in categoryList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item >
        <el-form-item label="摘要：">
          <el-input
            v-model="formData.desc"
            placeholder="输入文章摘要"
            type="textarea"
            :clearable = true
          />
        </el-form-item>
        <el-form-item label="缩略图：">
          <el-input
            v-model="formData.img"
            placeholder="输入图片URL，该图片会用于主页文章列表的显示"
            type="text"
            :clearable = true
          />
        </el-form-item>
        <el-form-item>
          <MdEditor v-model="formData.content" />
        </el-form-item>
      </el-form>
    </el-card>
  </div>

</template>

<script lang="ts" setup>
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { onMounted, reactive, ref } from "vue";
import type { FormInstance, FormRules } from 'element-plus';
import { ElMessage} from "element-plus";
import axios from "axios";
import router from "@/router";
const categoryList = ref([]);
const articleFormRef = ref<FormInstance>();
const formData: {[key: string]: string} =  reactive({
  title: '',
  cid: '',
  uid:'',
  desc:'',
  content:'',
  img:''
})
const rules = reactive<FormRules>({
  cid:[
    {
      required: true,
      message: '文章分类不能为空',
      trigger: ['blur', 'change']
    }
  ]
})
const props = defineProps({
  id: {
    type: String,
  }
});

// 如果文章id存在，则在表中显示文章相关内容
const showArticle = async (id :any) => {
  try {
    if(id === undefined){
      console.log("test")
    }else {
      const response = await axios.get( `article/${id}`);
      if (response.data.status !== 200) {
        throw new Error(response.data.message);
      }
      for (const key in formData) {
        //通过 Object.prototype.hasOwnProperty 方法来检查 response.data.data 对象中是否也有这个属性名，如果有则将其值赋给 formData 中相应的属
        if (Object.prototype.hasOwnProperty.call(response.data.data, key)) {
          formData[key] = response.data.data[key];
        }
      }

    }
    } catch (error: unknown) {
    const err = error as Error;
    if(err.message === "文章不存在"){
      // 显示错误信息
      ElMessage({
        message: err.message,
        type: "error",
      });
      await router.push(`/admin/write-article`)
    } else {
      // 显示错误信息
      ElMessage({
        message: err.message,
        type: "error",
      });
    }

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
//新增文章表单提交
const addArticle = async (formEl: FormInstance | undefined) => {
  console.log(formEl)
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const response = await axios.post(`article/add`, formData);
        if (response.data.status !== 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          await router.push(`/admin/Article`)
          return ElMessage({
            message:"发布文章成功",
            type:"success",
          })
        }
      } catch (error) {
        console.log("err:",error); // 处理错误响应
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
//编辑文章表单提交
const editArticle = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        const response = await axios.put(`article/${props.id}`, formData);
        if (response.data.status !== 200){
          return ElMessage({
            message:response.data.message,
            type:"error",
          })
        } else {
          await router.push(`/admin/Article`)
          return ElMessage({
            message:"更新文章成功",
            type:"success",
          })
        }
      } catch (error) {
        console.log("err:",error); // 处理错误响应
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

//点击发布文章按钮之后的事件
const postArticleBtn = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid) => {
    if (valid) {
      try{
        try{
          if (props.id !== undefined){
          const response = await axios.get( `article/${props.id}`);
          if (response.data.status !== 200) {
            throw new Error(response.data.message);
          }
          await editArticle(formEl)
          } else {
            await addArticle(formEl)
          }
        } catch (error: unknown) {
          const err = error as Error;
          if(err.message === "文章不存在"){
            await addArticle(formEl)
          }
        }
      } catch (error) {
        console.log("err:",error); // 处理错误响应
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
const cateChange = (value:any) => {
  if(value !== ''){
    console.log(value)
  }
}
// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息
onMounted(() => {
  showArticle(props.id)
  showCategories();
});
</script>


<style scoped>
</style>