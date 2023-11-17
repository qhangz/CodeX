<script setup lang="ts">
import { getArticleBackend } from '@/api/article'
import { ref, onMounted } from 'vue'
import ArticleList from './components/ArticleList.vue'

const backendArticle = ref([])
const getArticle = async () => {
    const res = await getArticleBackend()
    backendArticle.value = res.data
    // console.log(articleList.value);
}
onMounted(() => {
    getArticle()
})

</script>

<template>
    <div class="backend">
        <div v-for="(item,index) in backendArticle" :key="index">
            <ArticleList :articleList="item"></ArticleList>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.backend {
    width: 750px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}
</style>