import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import Layout from '@/views/Layout/index.vue'
import { useUserStore } from '@/stores/userStore'
import { createPinia } from 'pinia';

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
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
          meta: {
            title: 'CodeX - Home'
          },
          children: [
            {
              path: '/',
              name: 'hot',
              component: () => import('../views/Home/Hot.vue'),
              meta: {
                title: 'CodeX - 综合'
              },
            },
            {
              path: '/following',
              name: 'following',
              component: () => import('../views/Home/Following.vue'),
              meta: {
                title: 'CodeX - 关注'
              },
            },
            {
              path: '/frontend',
              name: 'frontend',
              component: () => import('../views/Home/Frontend.vue'),
              meta: {
                title: 'CodeX - 前端'
              },
            },
            {
              path: '/backend',
              name: 'backend',
              component: () => import('../views/Home/Backend.vue'),
              meta: {
                title: 'CodeX - 后端'
              },
            }
          ]
        },
        {
          // events page
          path: '/events',
          name: 'events',
          component: () => import('../views/Events/index.vue'),
          meta: {
            title: 'CodeX - events'
          },
        },
        {
          // chanllenge page
          path: '/chanllenge',
          name: 'chanllenge',
          component: () => import('../views/Chanllenge/index.vue'),
          meta: {
            title: 'CodeX - chanllenge'
          },
        },
        {
          // course page
          path: '/course',
          name: 'course',
          component: () => import('../views/Course/index.vue'),
          meta: {
            title: 'CodeX - course'
          },
        },
        {
          // games page
          path: '/games',
          name: 'games',
          component: () => import('../views/Games/index.vue'),
          meta: {
            title: 'CodeX - Games'
          },
        },
        {
          // user page
          path: '/user',
          name: 'user',
          component: () => import('../views/User/Mine.vue'),
          meta: {
            title: 'CodeX - user',
            isAuth: true
          },
        },
        {
          // user setting pages
          path: '/user/setting',
          name: 'userprofile',
          component: () => import('../views/Setting/index.vue'),
          meta: {
            title: 'CodeX - setting',
            isAuth: true
          },
          children: [
            {
              path: 'profile',
              name: 'profile',
              component: () => import('../views/Setting/Profile.vue'),
              meta: {
                title: 'CodeX - profile'
              },
            },
            {
              path: 'account',
              name: 'account',
              component: () => import('../views/Setting/Account.vue'),
              meta: {
                title: 'CodeX - Account'
              },
            },
            {
              path: 'common',
              name: 'common',
              component: () => import('../views/Setting/Common.vue'),
              meta: {
                title: 'CodeX - Common'
              },
            },
          ]
        },
        {
          // codexplore page
          path: '/codexplores',
          name: 'codexplores',
          component: () => import('../views/User/codexplores.vue'),
          meta: {
            title: 'CodeX - user'
          },
        },
        {
          // app page
          path: '/app',
          name: 'app',
          component: () => import('../views/About/App.vue'),
          meta: {
            title: 'CodeX - App'
          },
        },
        {
          // publish discuss page
          path: '/publish',
          name: 'publish',
          component: () => import('../views/Publish/index.vue'),
          meta: {
            title: 'CodeX - publish',
            // isAuth: true
          },
        },
        // discuss item page
        {
          path: '/discuss/:id',
          name: 'discuss',
          component: () => import('../views/Discuss/index.vue'),
          meta: {
            title: 'CodeX - 说说详情'
          },
        },
      ]
    },

    // games page
    {
      path: '/games/santaclaus',
      name: 'santaclaus',
      component: () => import('../views/Games/SantaClaus/index.vue'),
      meta: {
        title: 'CodeX - Santa Claus'
      },
    },
    // login and register page
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/index.vue'),
      meta: {
        title: 'CodeX - 登录'
      },
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
        title: 'CodeX - 关于'
      }
    },
    // page not found
    {
      path: '/404',
      name: '404',
      component: () => import('../views/404/index.vue'),
      meta: {
        title: 'CodeX - page not found'
      },
    }
  ]
})

const loginMsg = () => {
  ElMessage.error('请先登录！')
}

// 通过localStorage获取登录状态
router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title as string ? to.meta.title : '加载中';
  }
  // next();
  if (to.meta.isAuth) { // 判断该路由是否需要登录权限
    if (localStorage.isLogin === 'true') {
      next();
    } else {
      next('/login')
      loginMsg()
    }
  } else {
    next();
  }
})

// if page not found
router.beforeEach((to, from, next) => {
  if (to.matched.length === 0) {
    // 路由不存在
    next('/404');
  } else {
    next();
  }
});


// 通过useUserStore()获取userStore，进行登录状态判断
// let userStore: any = null
// router.beforeEach(async (to, from, next) => {
//   if (userStore === null) {
//     userStore = await useUserStore()
//   }
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

export default router
