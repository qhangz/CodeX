<script setup lang="ts">
import { defineProps, onMounted, ref } from 'vue';
import { getUserByUsername } from '@/api/user';

// get author info from parent component
const authorInfo = defineProps(['author'])
// console.log('author:', authorInfo.author);

// get user info by author
interface MyUser {
    ID: Int16Array;
    createTime: Date;
    updateTime: Date;
    deleteTime: Date;
    username: any;
    email: string;
    age: Int16Array;
    summary: string;
    avatar_iamge: string;
}
const userInfo = ref<MyUser[]>([])
const getUserInfo = async () => {
    const res = await getUserByUsername(authorInfo.author)
    // console.log(res.data);
    userInfo.value = res.data
}
onMounted(() => {
    getUserInfo()
})

// focus function
const focus = () => {
    ElMessage({
        message: '关注成功',
        type: 'success',
    })
}

// communicate function
const communicate = () => {
    ElMessage({
        message: '私信成功',
        type: 'success',
    })
}
</script>
<template>
    <div class="user-form">
        <div class="top">
            <img :src="userInfo.avatar_image" alt="avatar" class="avatar-image">
            <div class="info">
                <div class="username">
                    {{ userInfo.username }}
                </div>
                <div class="summary">
                    {{ userInfo.summary }}
                </div>
            </div>
        </div>
        <div class="buttom">
            <div class="focus" @click="focus">
                关注
            </div>
            <div class="communicate" @click="communicate">
                <RouterLink to="/404">
                    私信
                </RouterLink>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.user-form {
    width: 100%;
    background-color: var(--bg1);
    border-radius: 10px;
    padding: 20px 25px;

    .top {
        height: 58px;
        width: 100%;
        display: flex;
        gap: 15px;
        padding: 0 10px;

        .avatar-image {
            height: 58px;
            width: 58px;
            border-radius: 50%;
            transition: transform 0.5s ease-in-out;
        }

        .avatar-image:hover {
            transform: rotate(360deg);
        }

        .info {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            width: 100%;
            gap: 5px;

            .username {
                display: flex;
                align-items: center;
                width: 100%;
                // height: 50%;
                font-size: 18px;
                font-weight: 600;
                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

            .summary {
                display: flex;
                align-items: center;
                width: 100%;
                // height: 50%;
                font-size: 14px;
                overflow: hidden;
                word-break: break-all;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 1;
            }

        }
    }

    .buttom {
        // height: 58px;
        width: 100%;
        margin-top: 25px;
        display: flex;
        gap: 20px;
        justify-content: center;
        align-items: center;

        .focus {
            width: 100px;
            height: 34px;
            background-color: var(--primary-100);
            border-radius: 5px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 16px;
            color: var(--text-color2);
            transition: transform 0.2s ease-in-out;
            cursor: pointer;

            &:hover {
                transform: translate3d(0, -2.5px, 0);
                box-shadow: 1px 5px 8px rgb(0 0 0 / 15%);
            }
        }

        .communicate {
            width: 100px;
            height: 34px;
            background-color: var(--primary-500);
            border-radius: 5px;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 16px;
            color: var(--text-color4);
            transition: transform 0.2s ease-in-out;
            cursor: pointer;

            &:hover {
                transform: translate3d(0, -2.5px, 0);
                box-shadow: 1px 5px 8px rgb(0 0 0 / 15%);
            }
        }
    }
}
</style>
```