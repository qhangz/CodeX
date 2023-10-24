import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/views/layout/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'layout',
      component: Layout,
      children: [
        {
          // home page
          path: '',
          name: 'home',
          component: () => import('../views/home/index.vue')
        }
      ]
    },

    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login/index.vue')
    },


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
