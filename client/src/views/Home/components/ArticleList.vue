<script setup lang="ts">
import { defineProps } from 'vue';
import { useRouter } from 'vue-router'
const articleList = defineProps(['articleList'])

const router = useRouter()
const toArticle = () => {
    // router.push({
    //     path: '/article/',
    //     name: 'article',
    //     params: {
    //         id: articleList.articleList.id
    //     }
    // })
    const routeData = router.resolve({
        path: '/article/',
        name: 'article',
        params: {
            id: articleList.articleList.id
        }
    })

    window.open(routeData.href, '_blank')
}

</script>

<template>
    <div class="article-list" @click="toArticle">
        <div class="article-item" v-for="item in articleList" :key="item.id">
            <div class="article-title">{{ item.title }}</div>
            <div class="article-summary">{{ item.summary }}</div>
            <div class="article-author">{{ item.author }}</div>
            <div class="article-viewnum">{{ item.viewNum }}</div>
            <div class="article-likenum">{{ item.likeNum }}</div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.article-list {
    width: 720px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color:var(--bg1);
    border-radius: 10px;
    margin-bottom: 20px;
}

@media(max-width:730px) {
    .article-list {
        width: 95vw;
        // height:100vh;
    }
}
</style>
