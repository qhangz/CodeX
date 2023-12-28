<script setup lang="ts">
import { getMineDiscussList } from '@/api/discuss'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
// to discuss detail
const toDiscussDetail = (item: any) => {
    router.push({
        path: '/discuss/',
        name: 'discuss',
        params: {
            id: item.id
        },
        query:{
            author: item.author
        }
    })
}

interface Discuss {
    ID: number;
    created_at: string;
    author: string;
    title: string;
    summary: string;
    category: string;
    like_num: number;
    view_num: number;
}

const discussList = ref<Discuss[]>([])

const username = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
const getDiscussList = async () => {
    const res = await getMineDiscussList(username)
    discussList.value = res.data
}
onMounted(() => {
    getDiscussList()
})

// determine whether there are discusses
const hasDiscuss = ref(false)
if (discussList.value.length > 0) {
    hasDiscuss.value = true
}
</script>
<template>
    <div class="mine-discuss">
        <div class="top">
            <div class="title">我的说说：</div>
            <div class="num">{{ discussList.length }}</div>
        </div>
        <div class="discuss-list">
            <div class="discuss-item" v-for="(item, index) in discussList" :key="index" @click="toDiscussDetail(item)">
                <div class="title">{{ item.title }}</div>
                <div class="summary">{{ item.summary }}</div>
                <div class="footer">
                    <div class="category">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-leibie"></use>
                        </svg>
                        {{ item.category }}
                    </div>
                    <div class="line"></div>
                    <div class="viewnum">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-chakan"></use>
                        </svg>
                        {{ item.view_num }}
                    </div>
                    <div class="likenum">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-icon"></use>
                        </svg>
                        {{ item.like_num }}
                    </div>
                    <div class="created-at">
                        {{ item.created_at }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style scoped lang="scss">
.mine-discuss {
    width: 720px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    .top {
        width: 100%;
        display: flex;
        align-items: center;
        margin-bottom: 20px;
        font-size: 18px;
        color: var(--text-color1);
        background-color: var(--bg1);
        border-radius: 10px;
        padding: 20px 25px;
        gap: 10px;
    }

    .discuss-list {
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: 10px;

        .discuss-item {
            width: 100%;
            padding: 10px 25px;
            cursor: pointer;
            color: var(--text-color1);
            background-color: var(--bg1);
            border-radius: 10px;

            .title {
                width: 100%;
                height: 30px;
                font-size: 20px;
                line-height: 30px;
                color: var(--text-color1);
                font-weight: bold;
            }

            .summary {
                height: 30px;
                font-size: 13px;
                line-height: 30px;
                color: var(--text-color3);
                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

            .footer {
                display: flex;
                width: 100%;
                align-items: center;
                font-size: 13px;
                color: var(--text-color3);

                .category {
                    height: 30px;
                    line-height: 30px;
                    margin-right: 15px;
                }

                .line {
                    height: 13px;
                    display: flex;
                    border-right: 1px solid var(--text-color3);
                    margin-right: 15px;
                }

                .viewnum {
                    height: 30px;
                    line-height: 30px;
                    margin-right: 13px;
                }

                .likenum {
                    height: 30px;
                    line-height: 30px;
                }

                .created-at {
                    height: 30px;
                    line-height: 30px;
                    margin-left: auto;
                }
            }


        }

    }
}
</style>