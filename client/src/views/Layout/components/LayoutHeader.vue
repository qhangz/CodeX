<script setup lang="ts">
import { onMounted, ref, reactive } from 'vue'
import { Search } from '@element-plus/icons-vue'    // 引入搜索图标
import { useUserStore } from '@/stores/userStore'

import { useScroll } from '@vueuse/core'
const { y } = useScroll(window)

// user info interface
useUserStore().isUserLogin()
const isLogin = ref(useUserStore().userState.isLogin)
// const isLogin = ref(localStorage.getItem('isLogin') ? true : false)

// user avatar image
const avatarImage = reactive(localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).avatar_image : '')
// console.log('avatarImage', avatarImage);

// 自动补全输入框
interface RestaurantItem {
    value: string
    link: string
}

const state1 = ref('')

const restaurants = ref<RestaurantItem[]>([])
const querySearch = (queryString: string, cb: any) => {
    const results = queryString
        ? restaurants.value.filter(createFilter(queryString))
        : restaurants.value
    // call callback function to return suggestions
    cb(results)
}
const createFilter = (queryString: string) => {
    return (restaurant: RestaurantItem) => {
        return (
            restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        )
    }
}
const loadAll = () => {
    return [
        { value: 'vue', link: 'https://github.com/vuejs/vue' },
        { value: 'element', link: 'https://github.com/ElemeFE/element' },
        { value: 'cooking', link: 'https://github.com/ElemeFE/cooking' },
        { value: 'mint-ui', link: 'https://github.com/ElemeFE/mint-ui' },
        { value: 'vuex', link: 'https://github.com/vuejs/vuex' },
        { value: 'vue-router', link: 'https://github.com/vuejs/vue-router' },
        { value: 'babel', link: 'https://github.com/babel/babel' },
    ]
}
const handleSelect = (item: RestaurantItem) => {
    console.log(item)
}
onMounted(() => {
    restaurants.value = loadAll()
})
</script>
<template>
    <div class='app-header' :class="{ hidden: y > 200 }">
        <div class="container">
            <!-- logo -->
            <h1 class="logo">
                <RouterLink to="/codexplores">CodeX</RouterLink>
            </h1>
            <!-- header nav -->
            <ul class="app-header-nav">
                <li class="home">
                    <RouterLink to="/" :class="{ 'active': $route.path === '/' }">首页</RouterLink>
                </li>
                <li>
                    <RouterLink to="/events" :class="{ 'active': $route.path === '/events' }">活动</RouterLink>
                </li>
                <li>
                    <RouterLink to="/games" :class="{ 'active': $route.path === '/games' }">游戏</RouterLink>
                </li>
                <li>
                    <RouterLink to="/codexplores" :class="{ 'active': $route.path === '/codexplores' }">用户</RouterLink>
                </li>
                <li>
                    <RouterLink to="/app" :class="{ 'active': $route.path === '/app' }">APP</RouterLink>
                </li>

                <li>
                    <RouterLink to="/about" target="_blank" :class="{ 'active': $route.path === '/about' }">关于</RouterLink>
                </li>
            </ul>
            <!-- header nav dropdown -->
            <div class="app-header-nav-dropdown">
                <div class="dropdown-icon">
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-xuanxiang"></use>
                    </svg>
                </div>
                <ul class="dropdown-content">
                    <li>
                        <RouterLink to="/" :class="{ 'active': $route.path === '/' }">首页</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/events" :class="{ 'active': $route.path === '/events' }">活动</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/games" :class="{ 'active': $route.path === '/games' }">游戏</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/codexplores" :class="{ 'active': $route.path === '/codexplores' }">用户</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/app" :class="{ 'active': $route.path === '/app' }">APP</RouterLink>
                    </li>
                    <li>
                        <RouterLink to="/about" target="_blank" :class="{ 'active': $route.path === '/about' }">关于
                        </RouterLink>
                    </li>
                </ul>
            </div>
            <!-- searcbox -->
            <div class="search">
                <el-col :span="200">
                    <el-autocomplete v-model="state1" :fetch-suggestions="querySearch" clearable class="inline-input w-50"
                        placeholder="Code explore" @select="handleSelect" />
                    <el-button style="margin-left: 5px;" :icon="Search" circle />
                </el-col>
            </div>
            <!-- creaor center -->
            <div class="creator-center">
                <div class="btn">
                    <RouterLink to="/publish">创作中心</RouterLink>
                </div>
                <div class="dropdown-content">
                    <ul class="drop-item">
                        <li>
                            <RouterLink to="/404">
                                <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-yuedu"></use>
                                </svg>
                                <div class="text">
                                    写文章
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/404">
                                <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-menu-dataMining"></use>
                                </svg>
                                <div class="text">
                                    写代码
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/404">
                                <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-edit"></use>
                                </svg>
                                <div class="text">
                                    写笔记
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/404">
                                <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-caogaoxiang"></use>
                                </svg>
                                <div class="text">
                                    草稿箱
                                </div>
                            </RouterLink>
                        </li>
                    </ul>
                </div>
            </div>
            <!-- user info -->
            <div class="user" v-if="isLogin">
                <div class="is-login">
                    <div class="notice">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-xiaoxi"></use>
                        </svg>
                        <div class="dropdown">
                            <ul class="drop-item">
                            </ul>
                        </div>
                    </div>
                    <div class="avatar">
                        <RouterLink to="/user">
                            <img class="icon" :src="avatarImage" />
                        </RouterLink>
                    </div>
                </div>
            </div>
            <div class="user" v-else>
                <div class="not-login">
                    <RouterLink to="/login"><span>登录&nbsp;</span>&nbsp;&nbsp;注册</RouterLink>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.app-header {
    background-color: var(--bg1);
    height: 60px;
    width: 100%;
    padding: 5px 0;
    position: fixed;
    top: 0;
    left: 0;
    transition: top 0.2s ease-in-out;
    z-index: 999;

    &.hidden {
        top: -64px;
    }

    .container {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 20px;
        width: 100%;

        .logo {
            width: 50px;
            margin-left: 0px;

            a {
                display: block;
                height: 50px;
                width: 100%;
                text-indent: -9999px;
                background: url('@/assets/images/logo.png') no-repeat center 0px / contain;
            }
        }

        .app-header-nav {
            width: 560px;
            display: flex;
            padding-left: 20px;
            position: relative;
            z-index: 998;

            li {
                margin-right: 30px;
                width: 38px;
                text-align: center;

                a {
                    font-size: 16px;
                    line-height: 32px;
                    height: 32px;
                    display: inline-block;
                    color: var(--text-color3);

                    &:hover {
                        color: var(--primary-100);
                        border-bottom: 1px solid var(--primary-100);
                    }
                }

                .active {
                    color: var(--primary-100);
                }
            }
        }

        .app-header-nav-dropdown {
            position: relative;
            display: flex;
            font-size: 16px;
            line-height: 32px;
            height: 32px;
            width: 50px;
            justify-content: center;
            margin-right: 150px;
            margin-left: 50px;
            background-color: var(--primary-100);
            color: var(--text-color2);
            border-radius: 4px;
            box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.25);

            .dropdown-icon {
                font-size: 20px;
            }

            .dropdown-content {
                display: none;
                position: absolute;
                background-color: var(--bg3);
                // min-width: 160px;
                box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
                z-index: 1;
                margin-top: 32px;

                li {
                    list-style: none;
                    padding: 10px;
                    display: flex;
                    justify-content: center;
                    width: 100px;
                }

                li:hover {
                    background-color: var(--bg4);
                    cursor: pointer;
                }
            }
        }

        .app-header-nav-dropdown:hover .dropdown-content {
            display: block;
        }

        .search {
            display: flex;
        }

        .creator-center {
            height: 100%;
            width: 100px;
            position: relative;
            display: inline-block;
            z-index: 999;

            .btn {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                background-color: var(--primary-100);
                border-radius: 4px;
                // box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.25);
            }

            .btn a {
                color: var(--text-color2);
                // font-size: 16px;
            }

            .dropdown-content {
                display: none;
                position: absolute;
                background-color: var(--bg3);
                min-width: 160px;
                box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
                z-index: 1;
                border-radius: 5px;
                height: 95px;
                width: 300px;
                // margin-top: 17px;
                left: -100%;

                .drop-item {
                    margin-top: 20px;
                    margin-bottom: 10px;
                    text-align: center;
                    gap: 30px;
                    display: flex;
                    position: relative;
                    justify-content: center;
                    align-items: center;

                    .icon {
                        width: 30px;
                        height: 30px;
                    }

                    .text {
                        font-size: 12px;
                        width: 100%;
                    }
                }

                .drop-item li:hover {
                    border-radius: 10px;
                    -webkit-box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
                    -moz-box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
                    box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
                }
            }

            .dropdown-content img {
                width: 30px;
                height: 30px;
                margin-right: 10px;
            }
        }

        .creator-center:hover .dropdown-content {
            display: block;
        }

        .user {
            .is-login {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                // width: 100px;
                gap: 20px;

                .notice {
                    width: 32px;
                    height: 32px;

                    .icon {
                        width: 100%;
                        height: 100%;
                    }
                }

                .notice:hover {
                    cursor: pointer;
                }

                .avatar {
                    width: 40px;
                    height: 40px;
                    cursor: pointer;
                    transition: transform 0.5s ease-in-out;

                    .icon {
                        width: 100%;
                        height: 100%;
                        // background: url('@/assets/images/myavatar.jpg') no-repeat center 0px / contain;
                        border-radius: 50%;
                    }
                }

                .avatar:hover {

                    transform: rotate(360deg);
                }
            }

            .not-login {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                width: 100px;
                background-color: var(--primary-500);
                border-radius: 4px;
                border: 1px solid var(--primary-300);

                a {
                    color: $primary-100;
                    width: 100%;
                    display: flex;
                    justify-content: center;
                    align-items: center;

                    span {
                        border-right: 1px solid var(--primary-300);
                    }
                }
            }

            .not-login:hover {
                background-color: var(--primary-400);
            }
        }
    }
}

@media(max-width:1140px) {
    .search {
        display: none !important;
    }
}

@media (max-width: 900px) {
    .app-header-nav {
        display: none !important;
    }
}

@media (min-width: 900px) {
    .app-header-nav-dropdown {
        display: none !important;
    }
}

@media (max-width: 630px) {
    .app-header-nav-dropdown {
        margin-left: 0 !important;
    }
}

@media (max-width: 590px) {
    .app-header-nav-dropdown {
        margin-right: 50px !important;
    }
}

@media (max-width: 450px) {
    .app-header-nav-dropdown {
        margin-right: 0px !important;
    }
}

@media (max-width: 380px) {
    .creator-center {
        display: none !important;
    }
}
</style>
