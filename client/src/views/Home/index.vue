<script setup lang="ts">
import HomeNav from './components/HomeNav.vue'
import ArticleList from './components/ArticleList.vue'
import HomeAside from './components/HomeAside.vue'
import BackTop from '@/components/BackTop.vue'
</script>

<template>
    <div class="home">
        <HomeNav />
        <ArticleList />
        <HomeAside />
        <BackTop />
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    padding-top: 20px;
    gap: 20px;
}

.home-nav {
    height:1000px;
    
}


// @media (min-width: 1024px) {
//     .home {
//         display: flex;
//         justify-content: center;
//         align-items: center;
//         flex-direction: column;
//     }
// }
</style>