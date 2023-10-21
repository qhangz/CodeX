import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // 默认路由页面
    {
      path: '/',
      name: 'null',
      component: HomeView
    },
    // home page
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/home/index.vue')
    },
    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login/index.vue')
    },
    
    // test page
    // {
    //   path: '/test',
    //   name: 'test',
    //   component: () => import('../views/login/sf.vue')
    // },

    // about page
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
