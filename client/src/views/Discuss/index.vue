<script setup lang="ts">
import { useRouter } from 'vue-router';
import { getDiscussDetailById } from '@/api/discuss';
import { ref, onMounted } from 'vue'
import DiscussContent from './components/DiscussContent.vue'
import CommentContent from './components/CommentContent.vue'
import UserForm from './components/UserForm.vue'
import { addDiscussViewNumber } from '@/api/discuss';

const router = useRouter();
// 获取路由跳转携带的id和author
const discussId = router.currentRoute.value.params.id;
const author = router.currentRoute.value.query.author;

// get the details of this discuss by id
interface Comment {
    ID: number;
    CreatedAt: string;
    // UpdatedAt: string;
    // DeletedAt: string | null;
    DiscussID: number;
    author: string;
    content: string;
    like_num: number;
    view_num: number;
}
interface Discuss {
    ID: number;
    CreatedAt: string;
    // UpdatedAt: string;
    // DeletedAt: string | null;
    author: string;
    title: string;
    summary: string;
    content: string;
    category: string;
    like_num: number;
    view_num: number;
}

let discussInfo = ref<Discuss>({
    ID: 0,
    CreatedAt: "",
    // UpdatedAt: "",
    // DeletedAt: null,
    author: "",
    title: "",
    summary: "",
    content: "",
    category: "",
    like_num: 0,
    view_num: 0,
})
let commentInfo = ref<Comment[]>([])

const getDiscussInfo = async () => {
    // add view_num by the article id
    const addRes = await addDiscussViewNumber(discussId)

    // get the discussinfo
    const res = await getDiscussDetailById(discussId)
    commentInfo.value = res.data.Comment
    for (let i = 0; i < commentInfo.value.length; i++) {
        commentInfo.value[i].CreatedAt = commentInfo.value[i].CreatedAt.slice(0, 10)
    }
    discussInfo.value = {
        ID: res.data.ID,
        CreatedAt: res.data.CreatedAt,
        // UpdatedAt: res.data.UpdatedAt,
        // DeletedAt: res.data.DeletedAt,
        author: res.data.author,
        title: res.data.title,
        summary: res.data.summary,
        content: res.data.content,
        category: res.data.category,
        like_num: res.data.like_num,
        view_num: res.data.view_num,
    };
    // cut out the time ten digits
    discussInfo.value.CreatedAt = discussInfo.value.CreatedAt.slice(0, 10)
    // 给content添加换行符
    // discussInfo.value.content = discussInfo.value.content.replace(/\n/g, "<br />");
}


onMounted(() => {
    getDiscussInfo()
})
</script>

<template>
    <div class="article">
        <div class="content">
            <DiscussContent :discusses="discussInfo"></DiscussContent>
            <CommentContent :comments="commentInfo" :discussID="discussInfo.ID"></CommentContent>
        </div>
        <div class="aside">
            <UserForm :author="author"></UserForm>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.article {
    width: 100%;
    display: flex;
    justify-content: center;
    padding-top: 20px;
    gap: 20px;

    .content {
        width: 720px;
        gap: 300px;

        .comment-content {
            margin-top: 30px;
        }
    }

    .aside {
        width: 300px;
    }
}
</style>