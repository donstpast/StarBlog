// 导入 createApp 方法和 createPinia 方法
import { createApp } from 'vue'
import { createPinia } from 'pinia'
// 导入 ElementPlus 样式和组件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 导入 App 组件和 router 路由对象
import App from './App.vue'
import router from './router'
import axios from 'axios'
// 导入自定义的全局样式
import './assets/css/base.css'
//导入所有图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
axios.defaults.baseURL = 'http://localhost:8383/api/v1'
//axios请求拦截器
axios.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${sessionStorage.getItem('token')}`
  return config
})
// 创建 Vue 应用实例
const app = createApp(App)
//进行图标全局注册
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
// 安装 Pinia 状态管理插件
app.use(createPinia())
// 安装 ElementPlus 组件库
app.use(ElementPlus)
// 安装 Vue Router 路由插件
app.use(router)
// 挂载应用实例到页面
app.mount('#app')

