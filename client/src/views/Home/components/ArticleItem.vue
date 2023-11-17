<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();
// 获取路由跳转携带的id
const articleId = router.currentRoute.value.params.id;
console.log(articleId);
</script>

<template>
    <div class="article-item">
        this is article detials page

    </div>
</template>

<style lang="scss" scoped>
.article-item {
    width: 750px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

</style>