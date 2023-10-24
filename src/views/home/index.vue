<script setup lang="ts">

</script>

<template>
    <div class="home">
        this is home page
        <router-link to="/about">About</router-link>
        <el-button type="primary">button</el-button>
    </div>
</template>

<style lang="scss" scoped>
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