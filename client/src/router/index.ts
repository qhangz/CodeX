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
          path: '/',
          name: 'home',
          component: () => import('../views/Home/index.vue'),
          children: [
            {
              path: '/',
              name: 'hot',
              component: () => import('../views/Home/Hot.vue')
            },
            {
              path: '/following',
              name: 'following',
              component: () => import('../views/Home/Following.vue')
            },
            {
              path: '/frontend',
              name: 'frontend',
              component: () => import('../views/Home/Frontend.vue')
            },
            {
              path: '/backend',
              name: 'backend',
              component: () => import('../views/Home/Backend.vue')
            }
          ]
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
        },
        {
          // games page
          path: '/games',
          name: 'games',
          component: () => import('../views/Games/index.vue')
        },
        {
          // user page
          path: '/user',
          name: 'user',
          component: () => import('../views/User/index.vue')
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
      component: () => import('../views/About/AboutView.vue'),
      meta: {
        isAuth: true
      }
    }
  ]
})

const loginMsg = () => {
  ElMessage.error('请先登录！')
}
// 设置路由守卫（判断登录状态）
router.beforeEach((to, from, next) => {
  if (to.meta.isAuth) { // 对需要对登录的页面进行判断
    if (localStorage.isLogin === 'true') {
      next()  // 路由页面防止
    } else {
      next('/login')
      loginMsg()
    }
  } else {
    next()
  }
})

export default router
