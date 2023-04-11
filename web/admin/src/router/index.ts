import { createRouter, createWebHistory } from 'vue-router'  // 导入 createRouter 和 createWebHistory 函数，用于创建路由器和路由历史
import Login from '../views/Login.vue' // 导入登录页面组件
import Admin from '../views/Admin.vue' // 导入后台管理页面组件

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // 创建路由历史
  routes: [ // 路由配置数组
    {
      path: '/login', // 登录页面路由
      name: 'login', // 路由名称
      component: Login // 组件
    },
    {
      path: '/admin', // 后台管理页面路由
      name: 'admin', // 路由名称
      component : Admin, // 组件
      meta: { requiresAuth: true } // 添加需要认证的标识
    }
  ]
})

router.beforeEach((to, from, next) => { // 全局前置守卫，当每次路由切换时，都会执行该函数
  const token = sessionStorage.getItem('token') // 从 sessionStorage 中获取 token
  if (to.matched.some(record => record.meta.requiresAuth)) { // 判断路由是否需要认证，通过判断 to.matched 数组中每个记录的 meta 对象是否包含 requiresAuth 属性来实现
    if (token) { // token 存在，允许访问
      next()
    } else { // token 不存在，重定向到登录页面
      next({
        path: '/login', // 重定向的路由
        query: { redirect: to.fullPath } // 记录重定向前的路由
      })
    }
  } else { // 不需要认证
    next()
  }
})

export default router // 导出路由器实例
