// 管理用户数据相关
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userLogin } from '@/api/user'
import { createPinia } from 'pinia'

export const useUserStore = defineStore('user', () => {
    // 管理用户登录状态的state
    const userState = {
        isLogin: false,  //是否登录
        // userInfo: {},    //用户信息
        // token:'',       //用户token
    }
    // 管理用户数据的state
    const userInfo = ref({})

    // 登录
    const login = async ({ username, password }: { username: string, password: string }) => {
        const res = await userLogin(username, password)
        userState.isLogin = true
        userInfo.value = res
        // this.token = res.token
    }

    // 登出
    const logout = () => {
        userState.isLogin = false
        userInfo.value = {}
        // this.token = ''
    }

    // 改变登录状态
    const changeLoginState = (state: boolean) => {
        userState.isLogin = state
    }

    // 获取接口数据的action函数
    const getUserInfo = async ({ username, password }: { username: string, password: any }) => {
        const res = await userLogin(username, password)
        userInfo.value = res
    }
    // 退出时清除用户信息
    const clearUserInfo = () => {
        userInfo.value = {}
        userState.isLogin = false
    }

    return {
        userState,
        userInfo,
        login,
        logout,
        changeLoginState,
        getUserInfo,
        clearUserInfo,
    }
})

// const userState = {
//     isLogin: false,  //是否登录
//     userInfo: {},    //用户信息
//     // token:'',       //用户token
// }

// export const useUserStore = defineStore('user', {
//     state: () => userState,
//     getters: {
//         isLogin: (state) => state.isLogin,
//         userInfo: (state) => state.userInfo,
//         // token: (state) => state.token,
//     },
//     actions: {
//         // 登录
//         async login({ username, password }: { username: string, password: string }) {
//             const res = await userLogin(username, password)
//             this.isLogin = true
//             this.userInfo = res
//             // this.token = res.token
//         },
//         // 退出登录
//         logout() {
//             this.isLogin = false
//             this.userInfo = {}
//             // this.token = ''
//         },
//     },
// })

// export const useUserStore = defineStore('user', () => {
//     // 管理用户数据的state
//     const userInfo = ref({})
//     // 获取接口数据的action函数
//     const getUserInfo = async ({ username, password }: { username: string, password: any }) => {
//         const res = await userLogin(username, password)
//         userInfo.value = res
//     }
//     return {
//         userInfo,
//         getUserInfo
//     }
// })



// import { defineStore } from 'pinia'
// import { ref } from 'vue'
// import { loginAPI } from '@/apis/user'
// import { useCartStore } from './cartStore'
// import { mergeCartAPI } from '@/apis/cart'
// export const useUserStore = defineStore('user', () => {
//     const cartStore = useCartStore()
//     // 1. 定义管理用户数据的state
//     const userInfo = ref({})
//     // 2. 定义获取接口数据的action函数
//     const getUserInfo = async ({ account, password }) => {
//         const res = await loginAPI({ account, password })
//         userInfo.value = res.result
//         // 合并购物车的操作
//         await mergeCartAPI(cartStore.cartList.map(item => {
//             return {
//                 skuId: item.skuId,
//                 selected: item.selected,
//                 count: item.count
//             }
//         }))
//         cartStore.updateNewList()
//     }

//     // 退出时清除用户信息
//     const clearUserInfo = () => {
//         userInfo.value = {}
//         // 执行清除购物车的action
//         cartStore.clearCart()
//     }
//     // 3. 以对象的格式把state和action return
//     return {
//         userInfo,
//         getUserInfo,
//         clearUserInfo
//     }
// }, {
//     persist: true,
// })