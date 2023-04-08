import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'



const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login', 
      name: 'login',
      component: Login
    },
    {
      path: '/admin',
      name: 'admin', 
      component : Admin
      //路由懒加载：
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // component: () => import('../views/Admin.vue')
    }
  ]
})

export default router
