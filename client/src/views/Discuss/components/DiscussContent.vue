<script setup lang="ts">
import { defineProp } from 'vue';
import type { PropType } from 'vue';

interface Discuss {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    author: string;
    title: string;
    summary: string;
    content: string;
    category: string;
    like_num: number;
    view_num: number;
}
const discussInfo = defineProps({
    discusses: {
        type: Object as PropType<Discuss>,
        required: true
    }
})

// convert time format
const timeFormat = (time: string) => {
    const date = new Date(time)
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()
    const second = date.getSeconds()
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`
}
// discussInfo.discusses.CreatedAt = timeFormat(discussInfo.discusses.CreatedAt)

</script>
<template>
    <div class="discuss-content">
        <div class="title">
            {{ discusses.title }}
        </div>
        <div class="info">
            <div class="author">{{ discusses.author }}</div>
            <div class="created-at">{{ discusses.CreatedAt }}</div>
            <div class="view-num">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-chakan"></use>
                </svg>
                {{ discusses.view_num }}
            </div>
            <div class="like-num">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-icon"></use>
                </svg>
                {{ discusses.like_num }}
            </div>
            <div class="category">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-leibie"></use>
                </svg>
                专栏：{{ discusses.category }}
            </div>
        </div>
        <div class="summary">
            {{ discusses.summary }}
        </div>
        <div class="content">
            {{ discusses.content }}
        </div>
    </div>
</template>
<style lang="scss" scoped>
.discuss-content {
    width: 100%;
    background-color: var(--bg1);
    border-radius: 10px;
    padding: 20px 25px;

    .title {
        font-size: 32px;
        font-weight: bold;
        margin-bottom: 15px;
        width: 100%;
        color: var(--text-color1);
        word-break: break-all;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 1;
        overflow: hidden;
    }

    .info {
        width: 100%;
        display: flex;
        align-items: center;
        margin-bottom: 20px;
        font-size: 14px;
        color: var(--text-color3);
        gap: 20px;

        .view-num {
            display: flex;
            gap: 5px;
        }

        .like-num {
            display: flex;
            gap: 5px;
        }

        .category {
            display: flex;
            gap: 5px;
        }

        .icon {
            height: 18px;
            width: 18px;
            display: flex;
            align-items: center;
        }
    }

    .summary {
        font-size: 14px;
        color: var(--text-color3);
        margin-bottom: 20px;
        word-break: break-all;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
        overflow: hidden;
    }

    .content {
        font-size: 18px;
        color: var(--text-color1);
        line-height: 1.5;
    }
}
</style>
```