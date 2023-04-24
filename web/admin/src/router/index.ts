import { createRouter, createWebHistory } from 'vue-router' // 导入 createRouter 和 createWebHistory 函数，用于创建路由器和路由历史
import Login from '@/views/Login.vue' // 导入登录页面组件
import Admin from '@/views/Admin.vue' // 导入后台管理页面组件

// 导入页面路由组件
import writeArticle from '@/components/article/write-article.vue'
import ArticleList from '@/components/article/article-list.vue'
import CategoryList from '@/components/category/category-list.vue'
import userList from '@/components/user/user-list.vue'
import CommentList from "@/components/comment/comment-list.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // 创建路由历史
  routes: [
    // 路由配置数组
    {
      path: '/login', // 后台登录页面路由
      name: 'login', // 路由名称
      meta: {
        title: 'login'
      },
      component: Login // 组件
    },
    {
      path: '/', // 后台管理页面路由
      name: 'admin', // 路由名称
      component: Admin, // 组件
      meta: {
        title: 'admin',
        requiresAuth: true
      }, // 添加需要认证的标识
      children: [
        {
          path: 'write-article',
          component: writeArticle,
          meta: {
            title: '添加文章'
          }
        },
        {
          path: 'write-article/:id',
          component: writeArticle,
          props: true
        },
        {
          path: 'Article',
          component: ArticleList,
          name: 'Article',
          meta: {
            title: '文章列表',
          }
        },
        {
          path: 'Category',
          component: CategoryList,
          meta: {
            title: '分类列表'
          }
        },
        {
          path: 'User',
          component: userList,
          meta: {
            title: '用户列表'
          }
        },
        {
          path: 'Comment',
          component: CommentList,
          meta: {
            title: '评论列表'
          }
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  // 全局前置守卫，当每次路由切换时，都会执行该函数
  if (to.meta.title) {
    document.title = to.meta.title as string;
  }
  const token = sessionStorage.getItem('token') // 从 sessionStorage 中获取 token
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // 判断路由是否需要认证，通过判断 to.matched 数组中每个记录的 meta 对象是否包含 requiresAuth 属性来实现
    if (token) {
      // token 存在，允许访问
      next()
    } else {
      // token 不存在，重定向到登录页面
      next({
        path: '/login', // 重定向的路由
        query: { redirect: to.fullPath } // 记录重定向前的路由
      })
    }
  } else {
    // 不需要认证
    next()
  }
})

export default router // 导出路由器实例
