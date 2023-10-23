<script setup lang="ts">
import {ref} from 'vue'
const value=ref(new Date())
</script>

<template>
    <div class="home">
        this is home page
        <router-link to="/about">About</router-link>
        <el-calendar v-model="value" />
    </div>
</template>

<style lang="less" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

@media (min-width: 1024px) {
    .home {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
}
</style>