<script setup lang="ts">
import { defineProps } from 'vue';
import { useRouter } from 'vue-router'
const discussList = defineProps(['discussList'])
// console.log(discussList);
const router = useRouter()
const toDiscuss = () => {
    // router.push({
    //     path: '/article/',
    //     name: 'article',
    //     params: {
    //         id: articleList.articleList.id
    //     }
    // })
    const routeData = router.resolve({
        path: '/discuss/',
        name: 'discuss',
        params: {
            id: discussList.discussList.id
        }
    })

    window.open(routeData.href, '_blank')
}

</script>

<template>
    <div class="discuss-list" @click="toDiscuss">
        <div class="discuss-item" v-for="(item, index) in discussList" :key="index">
            <div class="title">{{ item.title }}</div>
            <div class="summary">{{ item.summary }}</div>
            <div class="footer">
                <div class="author">{{ item.author }}</div>
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
                <!-- 可能会存在多个标签 -->
                <!-- <div class="label">
                    <div v-for="(item, index) in discussLabel" :key="index">
                        {{ item }}
                    </div>
                </div> -->
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.discuss-list {
    width: 720px;
    background-color: var(--bg1);
    border-radius: 10px;
    margin-bottom: 20px;
    transition: all .2s linear;

    .discuss-item {
        width: 100%;
        padding: 10px 25px;
        cursor: pointer;

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

            .author {
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

    &:hover {
        content: "";
        transform: translate3d(0, -3.5px, 0);
        box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
        border-radius: 10px;
    }
}

@media(max-width:750px) {
    .discuss-list {
        width: 95vw;
        // height:100vh;
    }
}
</style>
