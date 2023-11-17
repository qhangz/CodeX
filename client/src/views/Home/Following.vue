<script setup lang="ts">
import { getArticleFollow } from '@/api/article'
import { ref, onMounted } from 'vue'
import ArticleList from './components/ArticleList.vue'

const followArticle = ref([])
const getArticle = async () => {
    const res = await getArticleFollow()
    followArticle.value = res.data
    // console.log(articleList.value);
}
onMounted(() => {
    getArticle()
})
</script>

<template>
    <div class="following">
        <div v-for="(item, index) in followArticle" :key="index">
            <ArticleList :articleList="item"></ArticleList>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.following {
    width: 750px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}
</style>