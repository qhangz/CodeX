<script setup lang="ts">
// import { getArticleHot } from '@/api/article'
import { getDiscussListByCategory } from '@/api/discuss'
import { ref, onMounted } from 'vue'
import DiscussList from './components/DiscussList.vue'

// const hotArticle = ref([])
// const getArticle = async () => {
//     const res = await getArticleHot()
//     hotArticle.value = res.data
//     console.log(hotArticle.value);
// }

const hotDiscuss = ref([])
const getDiscuss = async () => {
    const res = await getDiscussListByCategory('hot')
    hotDiscuss.value = res.data
    // console.log(hotDiscuss.value);
}
onMounted(() => {
    // getArticle()
    getDiscuss()
})

</script>

<template>
    <div class="hot">
        <div v-for="(item, index) in hotDiscuss" :key="index">
            <DiscussList :discussList="item"></DiscussList>
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