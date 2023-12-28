<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const toMine = () => {
    router.push({
        path: '/user'
    })
}
</script>
<template>
    <div class="header">
        <svg class="icon" aria-hidden="true"  @click="toMine">
            <use xlink:href="#icon-fanhui"></use>
        </svg>
        <div class="back"  @click="toMine">
            返回个人主页
        </div>
    </div>
</template>
<style lang="scss" scoped>
.header {
    width: 100%;
    display: flex;
    align-items: center;
    color: var(--text-color4);
    background-color: var(--bg1);
    border-radius: 5px;
    padding: 15px 25px;
    gap: 10px;
    font-size: 14px;
    .icon{
        cursor: pointer;
    }
    .back{
        cursor: pointer;
    }
}
</style>