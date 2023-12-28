<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';
import { onMounted, ref } from 'vue';
import { getUserByUsername } from '@/api/user'
import { useRouter } from 'vue-router';
const route = useRouter()

// 模拟登录登出
// const login = () => {
//     console.log("login");
//     useUserStore().login({ username: 'HANG', password: '123456' })
// }
// reload function
const reload = () => {
    window.location.reload()
}
const logout = () => {
    useUserStore().logout()
    console.log("logout");
    reload()
}

// route push to user profile
const toUserProfile = () => {
    route.push('/user/setting/profile')
}

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
let userInfo = ref<MyUser[]>([])

const getUserInfo = async () => {
    let myUserName = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
    const res = await getUserByUsername(myUserName)
    userInfo.value = res.data
}
onMounted(() => {
    getUserInfo()
})
</script>

<template>
    <div class="user-card">
        <div class="this-user">
            <div class="inner">
                <div class="top">
                    <img class="avatar-image" :src="userInfo.avatar_image" alt="avatar image">
                    <div class="info">
                        <div class="uid">UID&nbsp;:&nbsp;{{ userInfo.ID }}</div>
                        <div class="username">用户名&nbsp;:&nbsp;{{ userInfo.username }}</div>
                        <div class="email">邮箱&nbsp;:&nbsp;{{ userInfo.email }}</div>
                        <div class="age">年龄&nbsp;:&nbsp;{{ userInfo.age }}</div>
                        <div class="summary">简介&nbsp;:&nbsp;{{ userInfo.summary }}</div>
                    </div>
                </div>
                <hr class="divider">
                <div class="btn">
                    <div class="logout" @click="logout()">
                        登出
                    </div>
                    <div class="setting" @click="toUserProfile()">
                        设置
                    </div>
                </div>
            </div>
        </div>
        <!-- <div class="login" @click="login()">
            登录
        </div> -->
    </div>
</template>

<style lang="scss" scoped>
.user-card {
    width: 720px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    .this-user {
        color: var(--text-color1);
        width: 100%;
        background-color: var(--bg1);
        border-radius: 10px;
        padding: 20px 25px;

        .inner {
            display: flex;
            flex-direction: column;
            gap: 10px;

            .top {
                display: flex;
                flex-direction: row;
                align-items: center;

                .avatar-image {
                    background-size: cover;
                    background-repeat: no-repeat;
                    width: 100px;
                    height: 100px;
                    cursor: pointer;
                    transition: transform 0.5s ease-in-out;
                    border-radius: 50%;
                }

                .avatar-image:hover {
                    transform: rotate(360deg);
                }

                .info {
                    display: flex;
                    flex-direction: column;
                    font-size: 16px;
                    padding-left: 20px;
                    gap: 5px;

                    .username {
                        // font-weight:bold;
                        // font-size:16px;
                        overflow: hidden;
                        word-break: break-all;
                        text-overflow: ellipsis;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 1;
                    }

                    .email {
                        overflow: hidden;
                        word-break: break-all;
                        text-overflow: ellipsis;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 1;
                    }

                    .summary {
                        overflow: hidden;
                        word-break: break-all;
                        text-overflow: ellipsis;
                        display: -webkit-box;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 2;
                    }
                }
            }

            .divider {
                width: 90%;
                opacity: 70%;
                border-bottom: 1px solid var(--primary-700);
            }

            .btn {
                display: flex;
                justify-content: center;
                gap: 50px;

                .logout {
                    width: 118px;
                    height: 34px;
                    font-size: 16px;
                    background-color: var(--bg1);
                    color: var(--warning1);
                    border: 1px solid var(--warning1);
                    border-radius: 5px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    cursor: pointer;

                    &:hover {
                        opacity: 70%;
                    }
                }

                .setting {
                    width: 118px;
                    height: 34px;
                    font-size: 16px;
                    background-color: var(--bg1);
                    color: var(--text-color4);
                    border: 1px solid var(--primary-100);
                    border-radius: 5px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    cursor: pointer;

                    &:hover {
                        opacity: 70%;
                    }
                }
            }
        }

    }
}
</style>