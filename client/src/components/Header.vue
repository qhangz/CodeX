<template>
    <div class="header">
        <el-menu :default-active="activeIndex" class="el-menu-header" mode="horizontal" @select="handleSelect">
            <!-- <span style="width: 50px;"></span> -->
            <el-menu-item index="1">
                <span>
                    <img style="width: 50px;height: 50px" src="../../public/favicon.ico">
                </span>
            </el-menu-item>


            <el-menu-item index="2">
                <template #title>社区</template>
            </el-menu-item>

            <el-menu-item index="3">
                <template #title>说说</template>
            </el-menu-item>

            <div class="h-search">
                <el-autocomplete style="height: 50px; width: 350px;margin-top: 12px;" v-model="state1"
                    :fetch-suggestions="querySearch" clearable placeholder="探索社区" @select="handleSelect" />
                <el-button style="margin-top: 12px;margin-left: 5px;" :icon="Search" circle />
            </div>

            <div class="h-show">
                <div class="custom-loader"></div>
            </div>

            <el-sub-menu index="4">
                <template #title>
                    <el-button type="primary">创作中心</el-button>
                </template>
                <div class="w-center">
                    <div class="w-c-item">
                        <img class="w-c-it-img" src="~@/assets/image/write.png">
                        <div class="w-c-it-text">
                            写文章
                        </div>
                    </div>

                    <div class="w-c-item">
                        <img class="w-c-it-img" src="~@/assets/image/shuoshuo.png">
                        <div class="w-c-it-text">
                            写说说
                        </div>
                    </div>

                    <div class="w-c-item">
                        <img class="w-c-it-img" src="~@/assets/image/chaogao.png">
                        <div class="w-c-it-text">
                            草稿箱
                        </div>
                    </div>

                </div>
            </el-sub-menu>

            <div class="message-tip">
                <el-badge :value="100" :max="10" class="item">
                    <img class="message-img" src="~@/assets/image/notic.png">
                </el-badge>
            </div>

            <div class="avatar hover-rotate">
                <img class="avatar-img" src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg">
            </div>

        </el-menu>
    </div>
</template>


<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'


const activeIndex = ref('1')
const state1 = ref('')

// 自动补全输入框
interface RestaurantItem {
    value: string
    link: string
}
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
        { value: 'vue3', link: 'https://github.com/vuejs/vue' },
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


<style scoped lang="scss">
.header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 42px;
    background-color: #fff;
    /* 设置背景颜色 */
    padding: 20px;
    /* 添加内边距 */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    /* 添加阴影效果 */
    z-index: 999;
    /* 设置堆叠顺序，确保 header 位于其他内容之上 */
}

.avatar {
    cursor: pointer;
    margin-left: 60px;
    width: 40px;
    height: 40px;
    border-radius: 100px;
}

.hover-rotate {
    transition: transform 0.5s ease-in-out;
}

.hover-rotate:hover {
    transform: rotate(360deg);
}

.avatar-img {
    border-radius: 100px;
    margin-top: 10px;
    width: 100%;
    height: 100%;
}

.message-img {
    cursor: pointer;
    width: 25px;
    height: 30px;
}

.message-img:hover {
    border-radius: 10px;
    -webkit-box-shadow: 0px 0px 24px 6px rgba(5, 143, 235, 1);
    -moz-box-shadow: 0px 0px 24px 6px rgba(5, 143, 235, 1);
    box-shadow: 0px 0px 24px 6px rgba(5, 143, 235, 1);
}

.message-tip {
    margin-top: 12px;
    width: 30px;
}

.w-c-it-img {
    margin-top: 10px;
    width: 40%;
    height: 40px;
}

.w-c-it-text {
    margin-top: 10px;
    width: 100%;
    height: 30px;
    font-size: 14px;
    text-align: center;
}

.w-c-item:hover {
    border-radius: 10px;
    -webkit-box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
    -moz-box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
    box-shadow: 5px 5px 50px 0px rgba(105, 170, 214, 1);
}

.w-c-item {
    cursor: pointer;
    margin-top: 20px;
    width: 80px;
    height: 80px;
    display: flex;
    flex-direction: column;
    justify-items: center;
    align-items: center;
}

.w-center {
    border-radius: 30px;
    height: 120px;
    display: flex;
    padding-left: 60px;
    gap: 20px;
    width: 350px;
}

.custom-loader {
    margin-top: 12px;
    width: 120px;
    height: 22px;
    border-radius: 40px;
    color: #E4E4ED;
    position: relative;
    overflow: hidden;
}

.custom-loader::before {
    content: "";
    position: absolute;
    margin: 2px;
    width: 14px;
    top: 0;
    bottom: 0;
    left: -20px;
    border-radius: inherit;
    background: #2898dd;
    box-shadow: -10px 0 12px 3px #83daf1;
    clip-path: polygon(0 5%, 100% 0, 100% 100%, 0 95%, -30px 50%);
    animation: ct4 1s infinite linear;
}

@keyframes ct4 {
    100% {
        left: calc(100% + 20px)
    }
}

.h-search {
    display: flex;

}

.el-menu-header {
    gap: 20px;
    padding-left: 100px;
}
</style>