<script setup lang="ts">
import { ref } from 'vue'

// const navList=ref()
// 使用ref存储nav列表
const navList = ref([
    {
        key: 0,
        title: 'following',
    },
    {
        key: 1,
        title: 'frontend',
    },
    {
        key: 2,
        title: 'backend',
    }
])

const isActive = 0
const index = 0
const following = () => {

    console.log('following');
}
const frontend = () => {
    console.log('frontend');
}
const backend = () => {
    console.log('backend');
}
</script>

<template>
    <div class="home-nav">
        <!-- <ul class="home-nav-list">
            <li @click="following" :class="{ 'activeItem': index == isActive }">Following</li>
            <li @click="frontend" :class="{ 'activeItem': index == isActive }">Frontend</li>
            <li @click="backend" :class="{ 'activeItem': index == isActive }">Backend</li>
        </ul> -->
        <!-- <div class="home-nav-list">
            <div @click="following" :class="{ 'activeItem': index == isActive }">Following</div>
            <div @click="frontend" :class="{ 'activeItem': index == isActive }">Frontend</div>
            <div @click="backend" :class="{ 'activeItem': index == isActive }">Backend</div>
        </div> -->
        <div class="home-nav-list">
            <div v-for="(item, key) in navList" :key="key">{{ item.title }}</div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.home-nav {
    background-color: var(--bg1);

    .home-nav-list {
        width: 180px;
        height: 135px;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;

        li {
            height: 45px;
            width: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 16px;
        }

        .activeItem {
            color: red;
            background-color: red;
        }
    }
}

@media (max-width: 1230px) {
    .home-nav {
        display: none;
    }
}
</style>