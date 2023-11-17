<script setup lang="ts">
import { getArticleFrontend } from '@/api/article'
import { ref, onMounted } from 'vue'
import ArticleList from './components/ArticleList.vue'

const frontendArticle = ref([])
const getArticle = async () => {
    const res = await getArticleFrontend()
    frontendArticle.value = res.data
    // console.log(articleList.value);
}
onMounted(() => {
    getArticle()
})
</script>

<template>
    <div class="frontend">
        <div v-for="(item,index) in frontendArticle" :key="index">
            <ArticleList :articleList="item"></ArticleList>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.frontend {
    width: 750px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

</style>