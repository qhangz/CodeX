<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Search } from '@element-plus/icons-vue'    // 引入搜索图标

import { useScroll } from '@vueuse/core'
const { y } = useScroll(window)

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
                <RouterLink to="/">CodeX</RouterLink>
            </h1>
            <!-- header nav -->
            <ul class="app-header-nav">
                <li class="home">
                    <RouterLink to="/">首页</RouterLink>
                </li>
                <li>
                    <RouterLink to="/events">活动</RouterLink>
                </li>
                <li>
                    <RouterLink to="/chanllenge">竞赛</RouterLink>
                </li>
                <li>
                    <RouterLink to="/course">课程</RouterLink>
                </li>
                <li>
                    <RouterLink to="/">商城</RouterLink>
                </li>
                <li>
                    <RouterLink to="/">APP</RouterLink>
                </li>
                <li>
                    <RouterLink to="/">游戏</RouterLink>
                </li>
            </ul>
            <!-- searcbox -->
            <div class="search">
                <el-col :span="200">
                    <el-autocomplete v-model="state1" :fetch-suggestions="querySearch" clearable class="inline-input w-50"
                        placeholder="Code explore" @select="handleSelect" />
                    <el-button style="margin-left: 5px;" :icon="Search" circle />
                </el-col>
            </div>
            <!-- creaor center -->
            <div class="creatorCenter">
                <div class="btn">
                    <RouterLink to="/">创作中心</RouterLink>
                </div>
                <div class="dropdown-content">
                    <ul class="drop-item">
                        <li>
                            <RouterLink to="/"> <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-yuedu"></use>
                                </svg>
                                <div class="text">
                                    写文章
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/"> <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-menu-dataMining"></use>
                                </svg>
                                <div class="text">
                                    写代码
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/"> <svg class="icon" aria-hidden="true">
                                    <use xlink:href="#icon-edit"></use>
                                </svg>
                                <div class="text">
                                    写笔记
                                </div>
                            </RouterLink>
                        </li>
                        <li>
                            <RouterLink to="/"> <svg class="icon" aria-hidden="true">
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
        </div>
    </div>
</template>
<style lang="scss" scoped>
.app-header {
    background-color: #fff;
    height: 64px;
    width: 100%;
    padding: 7px 0;
    position: fixed;
    top: 0;
    left: 0;
    transition: top 0.2s ease-in-out;

    &.hidden {
        top: -64px;
    }

    .container {
        display: flex;
        align-items: center;
        gap: 20px;

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

                    &:hover {
                        color: $xtxColor;
                        border-bottom: 1px solid $xtxColor;
                    }
                }

                .active {
                    color: $xtxColor;
                    border-bottom: 1px solid $xtxColor;
                }
            }
        }

        .search {
            display: flex;
        }

        .creatorCenter {
            height: 100%;
            width: 100px;
            position: relative;
            display: inline-block;

            .btn {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 32px;
                background-color: $xtxColor;
                border-radius: 4px;
                // box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.25);
            }

            .btn a {
                color: #fff;
                // font-size: 16px;
                text-decoration: none;
            }

            .dropdown-content {
                display: none;
                position: absolute;
                background-color: #f9f9f9;
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

        .creatorCenter:hover .dropdown-content {
            display: block;
        }
    }
}
</style>
