<script setup lang="ts">
import { getDiscussListByCategory } from '@/api/discuss'
import { ref, onMounted } from 'vue'
import DiscussList from './components/DiscussList.vue'

const backendDiscuss = ref([])
const getDiscuss = async () => {
    const res = await getDiscussListByCategory('backend')
    backendDiscuss.value = res.data
    // console.log(articleList.value);
}
onMounted(() => {
    getDiscuss()
})

</script>

<template>
    <div class="backend">
        <div v-for="(item, index) in backendDiscuss" :key="index">
            <DiscussList :discussList="item"></DiscussList>
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