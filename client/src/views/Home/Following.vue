<script setup lang="ts">
import { getDiscussListByCategory } from '@/api/discuss'
import { ref, onMounted } from 'vue'
import DiscussList from './components/DiscussList.vue'

const followDiscuss = ref([])
const getDiscuss = async () => {
    const res = await getDiscussListByCategory('follow')
    followDiscuss.value = res.data
    // console.log(articleList.value);
}
onMounted(() => {
    getDiscuss()
})
</script>

<template>
    <div class="following">
        <div v-for="(item, index) in followDiscuss" :key="index">
            <DiscussList :discussList="item"></DiscussList>
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