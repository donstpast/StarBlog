<template>
  <el-card
    class="commentCard box-card"
    :body-style="cardBodyStyle"
  >
    <template #header>
      <div class="card-header">
        <h3>{{formData.title}}</h3>
      </div>
    </template>
    <div>
      <MdEditor
      :previewOnly = true
      v-model="formData.content"
      :preview-theme="'mk-cute'" >
      </MdEditor>
    </div>


  </el-card>
</template>
<script setup lang="ts">
import { onMounted, reactive} from "vue";
import MdEditor from 'md-editor-v3'
import '@/assets/css/editor.css'
import router from "@/router";
import axios from "axios";
import { ElMessage } from "element-plus";
const formData: { [key: string]: string } = reactive({
  title: '',
  cid: '',
  uid: '',
  desc: '',
  content: '',
  img: ''
})
const cardBodyStyle:any = {
  padding: '0px',
};

// 如果文章id存在，则在表中显示文章相关内容
const showArticle = async (id: any) => {
  try {
    if (id === undefined) {
      console.log('test')
    } else {
      const response = await axios.get(`article/${id}`)
      if (response.data.status !== 200) {
        throw new Error(response.data.message)
      }
      for (const key in formData) {
        //通过 Object.prototype.hasOwnProperty 方法来检查 response.data.data 对象中是否也有这个属性名，如果有则将其值赋给 formData 中相应的属
        if (Object.prototype.hasOwnProperty.call(response.data.data, key)) {
          formData[key] = response.data.data[key]
        }
      }
    }
  } catch (error: unknown) {
    const err = error as Error
    if (err.message === '文章不存在') {
      // 显示错误信息
      ElMessage({
        message: err.message,
        type: 'error'
      })
      await router.push(`/`)
    } else {
      // 显示错误信息
      ElMessage({
        message: err.message,
        type: 'error'
      })
    }
  }
}
const props = defineProps({
  id: {
    type: String
  }
})
// 在组件挂载完成后，调用 showUsers() 函数展示所有用户信息
onMounted(() => {
  showArticle(props.id)
})
</script>

<style scoped>

</style>