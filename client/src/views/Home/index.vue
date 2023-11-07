<script setup lang="ts">
import HomeNav from './components/HomeNav.vue'
import HomeAside from './components/HomeAside.vue'
</script>

<template>
    <div class="home">
        <HomeNav />
        <router-view />
        <HomeAside/>
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction:row;
    padding-top:20px;
    gap:20px;
}

.home-nav {

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