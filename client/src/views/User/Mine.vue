<script setup lang="ts">
import UserCard from './components/UserCard.vue'
import MineDiscuss from './components/MineDiscuss.vue'
import GreetingAside from './components/GreetingAside.vue';
</script>

<template>
    <div class="user">
        <div class="content">
            <UserCard />
            <MineDiscuss />
        </div>
        <div class="aside">
            <GreetingAside />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.user {
    width: 100%;
    display: flex;
    justify-content: center;
    padding-top: 20px;
    gap: 20px;
    .content{
        width:720px;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        gap: 30px;
    }
    .aside{
        width:300px;
    }
}
</style>