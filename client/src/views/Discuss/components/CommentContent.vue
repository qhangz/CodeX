<script setup lang="ts">
import { ref } from 'vue';
import type { PropType } from 'vue';
import { submitComment } from '@/api/comment';
import { addCommentLikeNumber, addCommentViewNumber } from '@/api/comment';

// get the comment list from index.vue
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
    },
    discussID: {
        type: Number,
        required: true
    }
})

// get the user's avatar
const userAvatar = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).avatar_image : ''

// publish comment
const submitWaring = () => {
    ElMessage({
        message: '请输入评论内容',
        type: 'warning',
    })
}
const submitSuccess = () => {
    ElMessage({
        message: '发送成功',
        type: 'success',
    })
}
const submitFail = () => {
    ElMessage.error('发送失败')
}
const userNotLogin = () => {
    ElMessage.error('请先登录')
}
// reload function(refresh page)
const reload = () => {
    window.location.reload()
}
const textareaValue = ref('')
const publishComment = () => {
    console.log("publish comment");
    console.log(textareaValue.value);
    console.log(comment.discussID);
    if (textareaValue.value === '') {
        submitWaring()
    } else {
        if (localStorage.getItem('isLogin')) {
            const author = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
            const res = submitComment(comment.discussID, author, textareaValue.value).then(res => {
                if (res.code == 200) {
                    reload()
                    submitSuccess()
                } else {
                    submitFail()
                }
            })
        } else {
            userNotLogin()

        }
    }
}

// determine whether there are comments
const hasComment = ref(false)
if (comment.comments.length > 0) {
    hasComment.value = true
}

// textarea value
const textarea = ref('')

// add like number
const addLikeNumber = async (item: any) => {
    const res = await addCommentLikeNumber(item.ID)
    if (res.code === '200') {
        // comment.comments[id].like_num += 1
        item.like_num += 1
    }
}

// add view number
const addViewNumber = async (item: any) => {
    const res = await addCommentViewNumber(item.ID)
    if (res.code === '200') {
        item.view_num += 1
    }
}

</script>
<template>
    <div class="comment-content">
        <div class="top">
            <div class="title">评论</div>
            <div class="num">{{ comments.length }}</div>
        </div>
        <div class="publish-box">
            <img :src="userAvatar" alt="" class="avatar">
            <textarea placeholder="平等表达，友善交流" v-model="textareaValue" class="input-box" name="" id="input-filed" cols="70"
                rows="5" maxlength="250"></textarea>
            <div class="publish-btn" @click="publishComment">发送</div>
        </div>
        <div class="comment-list" v-if="hasComment">
            评论列表：
        </div>
        <div class="comment" v-for="(item, index) in comments" :key="index">
            <img :src="item.avatar" alt="" class="avatar-image">
            <div class="comment-content">
                <div class="comment-info">
                    <div class="comment-author">{{ item.author }}</div>
                </div>
                <div class="comment-content">{{ item.content }}</div>
                <div class="comment-action">
                    <div class="comment-time">{{ item.CreatedAt }}</div>
                    <div class="view-num">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-chakan"></use>
                        </svg>
                        {{ item.view_num }}
                    </div>
                    <div class="like-num" @click="addLikeNumber(item)">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-icon"></use>
                        </svg>
                        {{ item.like_num }}
                    </div>
                    <div class="comment-reply" @click="addViewNumber(item)">
                        <span>回复</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.comment-content {
    width: 100%;
    background-color: var(--bg1);
    border-radius: 10px;
    padding: 20px 25px;

    .top {
        width: 100%;
        display: flex;
        align-items: center;
        margin-bottom: 20px;
        font-size: 18px;
        color: var(--text-color1);
        gap: 10px;
    }

    .publish-box {
        display: flex;
        flex-direction: row;
        gap: 20px;

        .avatar {
            height: 40px;
            width: 40px;
            border-radius: 50%;
            transition: transform 0.5s ease-in-out;
        }

        .avatar:hover {
            transform: rotate(360deg);
        }

        .input-box {
            background-color: var(--bg2);
            width: 550px;
            border: 1px solid var(--border1);
            transition: border-color 0.3s;
            border-radius: 7px;
            padding: 10px;
            font-size: 14px;
            font: 1em/1.4 'Microsoft Yahei', 'PingFang SC', 'Avenir', 'Segoe UI',
                'Hiragino Sans GB', 'STHeiti', 'Microsoft Sans Serif', 'WenQuanYi Micro Hei',
                sans-serif;
        }

        .input-box:focus {
            border-color: var(--border2);
        }

        #input-filed {
            resize: vertical;
        }

        .publish-btn {
            width: 60px;
            height: 28px;
            background-color: var(--bg6);
            border-radius: 5px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 14px;
            color: var(--text-color2);
            cursor: pointer;
            transition: background-color 0.2s ease-in-out;
        }

        .publish-btn:hover {
            background-color: var(--bg7);
        }

    }

    .comment-list {
        padding: 20px 0 10px 0;
        font-size: 16px;
        color: var(--text-color1);
    }

    .comment {
        display: flex;
        flex-direction: row;
        padding-top: 20px;

        .avatar-image {
            height: 32px;
            width: 32px;
            border-radius: 50%;
        }

        .comment-content {
            padding: 5px 15px;

            .comment-info {
                width: 100%;
                display: flex;
                align-items: center;
                margin-bottom: 5px;
                font-size: 14px;
                color: var(--text-color3);
                gap: 20px;
            }

            .comment-content {
                font-size: 14px;
                color: var(--text-color1);
                line-height: 1.5;
                padding: 5px 0;
            }

            .comment-action {
                width: 100%;
                display: flex;
                align-items: center;
                // margin-bottom: 20px;
                font-size: 12px;
                color: var(--text-color3);
                gap: 20px;

                .view-num {
                    cursor: pointer;
                    transition: color 0.2s ease-in-out;

                    &:hover {
                        color: var(--primary-100);
                    }
                }

                .like-num {
                    cursor: pointer;
                    transition: all .2s linear;

                    &:hover {
                        color: var(--primary-100);
                        background-color: var(--bg5);
                        border-radius: 5px;
                        padding: 0px 7px;
                    }
                }

                .comment-reply {
                    cursor: pointer;
                    transition: all 0.2s ease-in-out;

                    &:hover {
                        color: var(--primary-100);
                        background-color: var(--bg5);
                        border-radius: 5px;
                        padding: 0px 7px;
                    }
                }
            }
        }
    }
}
</style>
```