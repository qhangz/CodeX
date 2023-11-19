<script setup lang="ts">
import { defineProps } from 'vue';
import { useRouter } from 'vue-router'
const articleList = defineProps(['articleList'])
console.log(articleList);
const router = useRouter()
const toArticle = () => {
    // router.push({
    //     path: '/article/',
    //     name: 'article',
    //     params: {
    //         id: articleList.articleList.id
    //     }
    // })
    const routeData = router.resolve({
        path: '/article/',
        name: 'article',
        params: {
            id: articleList.articleList.id
        }
    })

    window.open(routeData.href, '_blank')
}

</script>

<template>
    <div class="article-list" @click="toArticle">
        <div class="article-item" v-for="item in articleList" :key="item.id">
            <div class="title">{{ item.title }}</div>
            <div class="summary">{{ item.summary }}</div>
            <div class="footer">
                <div class="author">{{ item.author }}</div>
                <div class="line"></div>
                <div class="viewnum">
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-chakan"></use>
                    </svg>
                    {{ item.viewNum }}
                </div>
                <div class="likenum">
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-icon"></use>
                    </svg>
                    {{ item.likeNum }}
                </div>
                <!-- 可能会存在多个标签 -->
                <!-- <div class="label">
                    <div v-for="(item, index) in articleLabel" :key="index">
                        {{ item }}
                    </div>
                </div> -->
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.article-list {
    width: 720px;
    background-color: var(--bg1);
    border-radius: 10px;
    margin-bottom: 20px;

    .article-item {
        width: 100%;
        padding: 10px 25px;

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
        }
    }
}

@media(max-width:750px) {
    .article-list {
        width: 95vw;
        // height:100vh;
    }
}
</style>
