import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ArticleCard from "@/components/homePage/ArticleCard.vue";
import ArticleDetail from "@/components/homePage/ArticleDetail.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [
        {
          path:'/',
          component:ArticleCard
        },
        {
          path:'article/:id',
          component:ArticleDetail,
          props: true
        }
      ]
    }
  ]
})
// {
//   path: '/about',
//   name: 'about',
//   // route level code-splitting
//   // this generates a separate chunk (About.[hash].js) for this route
//   // which is lazy-loaded when the route is visited.
//   // component: () => import('../views/.vue')
// }
export default router
