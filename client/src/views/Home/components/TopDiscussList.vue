<script setup lang="ts">
import { useDiscussListStore } from '@/stores/discussStore'
import { getDiscussTop } from '@/api/discuss'
import { ref, reactive, onMounted } from 'vue'

// get discuss list
let discussList = ref<{ id: number, title: string }[]>([])
let listNum = ref([1, 2, 3, 4, 5])
const getDiscussList = async (index: number) => {
    const res = await getDiscussTop(index, index + 4)
    discussList.value = res.data
    console.log("res.data", res.data);
    console.log("discusslist.value", discussList.value);
}

// get discuss index from discussStore
const discussIndex = useDiscussListStore()
const changeList = () => {
    console.log("change");
    console.log(discussIndex.index);
    getDiscussList(discussIndex.index)
    listNum.value = listNum.value.map(item => item + 5)
    discussIndex.increment()
}

onMounted(() => {
    getDiscussList(discussIndex.index)
    discussIndex.increment()
})
</script>
<template>
    <div class="top-discuss-list">
        <div class="top">
            <svg class="icon" aria-hidden="true">
                <use xlink:href="#icon-yuedu"></use>
            </svg>
            <div class="title">
                文章榜
            </div>
            <div class="change" @click="changeList">
                <svg class="change-icon" aria-hidden="true">
                    <use xlink:href="#icon-reset"></use>
                </svg>
                <div class="change-text">
                    换一换
                </div>
            </div>
        </div>
        <hr class="divider">
        <div class="discuss-list">
            <div class="item" v-for="(item, index) in discussList" :key="index">
                <div class="num">{{ index + 1 }}</div>
                <div class="title">{{ item.title }}</div>
            </div>
        </div>
        <hr class="divider">
        <div class="more">
            <RouterLink to="/404">
                查看更多
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-gengduo"></use>
                </svg>
            </RouterLink>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.top-discuss-list {
    width: 100%;
    background: var(--bg1);
    border-radius: 5px;
    padding: 15px 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    // gap: 5px;
    margin-top: 20px;

    .top {
        height: 32px;
        width: 100%;
        display: flex;
        padding: 0 7px;

        .icon {
            height: 24px;
            width: 24px;
            margin: 4px 0px;
            margin-left: 5px;
        }

        .title {
            flex: 1;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 16px;
            color: var(--text-color1)
        }

        .change {
            display: flex;
            border-radius: 10px;
            padding: 0 5px;

            .change-icon {
                width: 18px;
            }

            .change-text {
                margin-left: 5px;
                display: flex;
                align-items: center;
                font-size: 14px;
                color: var(--text-color3);

                &:hover {
                    color: var(--primary-100);
                }
            }

            &:hover {
                background-color: var(--bg5);
            }
        }
    }

    .divider {
        width: 85%;
        opacity: 70%;
        border-bottom: 1px solid var(--primary-700);
    }

    .discuss-list {
        height: 120px;
        width: 100%;
        padding: 0 25px;

        .item {
            height: 20%;
            width: 100%;
            display: flex;
            align-items: center;
            font-size: 14px;
            color: var(--text-color1);

            .num {
                width: 20px;
                margin-right: 10px;
                color: var(--primary-100);
            }

            .title {
                width: 80%;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
        }
    }

    .more {
        height: 32px;
        display: flex;
        align-items: center;
        font-size: 14px;
        border-radius: 10px;
        padding: 0 7px;

        a {
            color: var(--text-color3);

            &:hover {
                color: var(--primary-100);
            }
        }

        &:hover {
            background-color: var(--bg5);
        }
    }
}
</style>