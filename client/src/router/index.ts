import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/views/Layout/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'layout',
      component: Layout,
      children: [
        {
          // home page，path: ' ' is the default child path
          path: '',
          name: 'home',
          component: () => import('../views/Home/index.vue')
        },
        {
          // events page
          path: '/events',
          name: 'events',
          component: () => import('../views/Events/index.vue')
        },
        {
          // chanllenge page
          path: '/chanllenge',
          name: 'chanllenge',
          component: () => import('../views/Chanllenge/index.vue')
        },
        {
          // course page
          path: '/course',
          name: 'course',
          component: () => import('../views/Course/index.vue')
        }
      ]
    },

    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/index.vue')
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
