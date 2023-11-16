<script setup lang="ts">
import { getArticleList } from '@/api/article'
import { ref, onMounted } from 'vue'

const articleList = ref([])
const getArticle = async () => {
    const res = await getArticleList()
    articleList.value = res.data
    console.log(articleList.value);
}
onMounted(() => {
    getArticle()
})

</script>

<template>
    <div class="hot">
        this is hot page

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