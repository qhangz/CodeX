<script setup lang="ts">
import { getArticleHot } from '@/api/article'
import { ref, onMounted } from 'vue'
import ArticleList from './components/ArticleList.vue'

const hotArticle = ref([])
const getArticle = async () => {
    const res = await getArticleHot()
    hotArticle.value = res.data
    console.log(hotArticle.value);
}
onMounted(() => {
    getArticle()
})

</script>

<template>
    <div class="hot">
        <div v-for="(item,index) in hotArticle" :key="index">
            <ArticleList :articleList="item"></ArticleList>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.hot {
    width: 750px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}
</style>