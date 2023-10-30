<script setup lang="ts">
// import LayoutNav from './components/LayoutNav.vue'
import LayoutHeader from './components/LayoutHeader.vue'
import LayoutFooter from './components/LayoutFooter.vue'
// import Header from '@/components/Header.vue'
</script>

<template>
    <div class="layout">
        <!-- <Header></Header> -->
        <!-- <LayoutNav /> -->
        <LayoutHeader />
        <RouterView />
        <LayoutFooter />
    </div>
</template>

<style lang="scss" scoped>

</style>