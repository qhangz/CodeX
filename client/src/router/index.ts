import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/views/Layout/index.vue'
import { useUserStore } from '@/stores/userStore'
import { createPinia } from 'pinia';

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
    // article item page
    {
      path: '/article/:id',
      name: 'article',
      component: () => import('../views/Home/components/ArticleItem.vue')
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

let userStore: any = null
// 前置拦截器
const interceptor = router.beforeEach(async (to, from, next) => {
  // userStore = useUserStore()
  if (userStore===null) {
    userStore =await useUserStore()
  }
  console.log("router.ts:",userStore.userState.isLogin);
  if (to.meta.isAuth) { // 对需要对登录的页面进行判断
    if (userStore.userState.isLogin === true) {
      next()  // 路由页面防止
    } else {
      next('/login')
      loginMsg()
    }
  } else {
    next()
  }
})
router.beforeEach(interceptor)

// 设置路由守卫（判断登录状态）
// router.beforeEach((to, from, next) => {
//   // let isLogin = useUserStore().userState.isLogin
//   const userStore = useUserStore()
//   console.log(userStore.userState.isLogin);
//   if (to.meta.isAuth) { // 对需要对登录的页面进行判断
//     if (userStore.userState.isLogin === true) {
//       next()  // 路由页面防止
//     } else {
//       next('/login')
//       loginMsg()
//     }
//   } else {
//     next()
//   }
// })

// router.beforeEach((to, from, next) => {
//   const store = useUserStore()
//   if (to.meta.isAuth && !store.userState.isLogin) {
//     next('/login')
//     loginMsg()
//   } else {
//     next()
//   }
// })

// router.beforeEach((to, from, next) => {
//   if (to.matched.length === 0) {
//     // 路由不存在
//     next('/404');
//   } else {
//     next();
//   }
// });

// router.beforeEach((to) => {
//   // ✅ 这样做是可行的，因为路由器是在其被安装之后开始导航的，
//   // 而此时 Pinia 也已经被安装。
//   const store = useUserStore()

//   if (to.meta.isAuth && !store.userState.isLogin) return '/login'
// })
// router.beforeEach((to, from, next) => {
//   // 我们想要在这里使用 store
//   if (store.userState.isLogin) next()
//   else next('/login')
// })
// router.beforeEach((to, from, next) => {
//   if (to.meta.isAuth) { // 对需要对登录的页面进行判断
//     if (localStorage.isLogin === 'true') {
//       next()  // 路由页面防止
//     } else {
//       next('/login')
//       loginMsg()
//     }
//   } else {
//     next()
//   }
// })

export default router
