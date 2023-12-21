<script setup lang="ts">
import HomeNav from './components/HomeNav.vue'
import HomeAside from './components/HomeAside.vue'
import BackTop from '@/components/BackTop.vue'
</script>

<template>
    <div class="home">
        <HomeNav />
        <router-view />
        <HomeAside />
        <BackTop />
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    // align-items: center;
    // flex-direction: row;
    padding-top: 20px;
    gap: 20px;
    // height: 100%;
    // position: relative;
    // max-width: 100%;
    // width: 100%;
    // height: 100vh;
}
</style>