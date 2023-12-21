<script setup lang="ts">
import { defineComponent, toRef } from 'vue';
import type { PropType } from 'vue';

interface Comment {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    DiscussID: number;
    author: string;
    content: string;
    like_num: number;
    view_num: number;
}

const comment = defineProps({
    comments: {
        type: Array as PropType<Comment[]>,
        required: true
    }
})
// console.log(comment);
// const commentsRef = toRef(comment, 'comments');
// console.log("commentsRef", commentsRef.value);

</script>
<template>
    <div class="comment-content">
        comment content
        <div class="comment" v-for="(item, index) in comments" :key="index">
            {{ item.author }}
            {{ item.content }}
            <!-- <div class="comment-item" v-for="(item, index) in commentsRef" :key="index">
                <div class="comment-author">{{ item.author }}</div>
                <div class="comment-content">{{ item.content }}</div>
            </div> -->
        </div>
    </div>
</template>
<style lang="scss" scoped>
.comment-content {
    width: 100%;
    background-color: var(--bg1);
    border-radius: 10px;
    padding: 20px 25px;
}
</style>
```