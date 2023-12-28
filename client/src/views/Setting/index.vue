<script setup lang="ts">
import Header from './components/Header.vue'
import Nav from './components/SettingNav.vue'
</script>
<template>
    <div class="setting">
        <div class="top">
            <Header></Header>
        </div>
        <div class="content">
            <Nav></Nav>
            <router-view />
        </div>
    </div>
</template>
<style lang="scss" scoped>
.setting {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    gap: 30px;

    .top {
        width: 1080px;
        margin-top: 20px;

    }

    .content {
        width: 1080px;
        display: flex;
        // justify-content: center;
        // align-items: center;
        flex-direction: row;
        gap:30px;
    }
}
</style>