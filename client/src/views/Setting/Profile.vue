<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { getUserByUsername, usernameUpdate, passwordUpdate, emailUpdate, ageUpdate, summaryUpdate } from '@/api/user';

// message function
const updateSuccessMsg = () => {
    ElMessage({
        message: '更新成功',
        type: 'success',
    })
}
const updateWarningMsg = () => {
    ElMessage({
        message: '请填写信息',
        type: 'warning',
    })
}
const pswWarningMsg = () => {
    ElMessage({
        message: '请输入密码',
        type: 'warning',
    })
}
const errorMsg = () => {
    ElMessage.error('Oops. 出错了')
}
// reload function
const reload = () => {
    window.location.reload()
}
// form info
// old username
const myUsername = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
const form = reactive({
    username: myUsername,
    password: '',
    newusername: '',
    newpassword: '',
    newemail: '',
    newage: '',
    newsummary: '',
    oldUsername: myUsername
})

// change function
const changeUsername = () => {
    if (form.username != '' && form.newusername != '') {
        usernameUpdate(form.username, form.newusername).then(res => {
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else {
        updateWarningMsg()
    }
}
const changePassword = () => {
    if (form.username != '' && form.password != '' && form.newpassword != '') {
        passwordUpdate(form.username, form.password, form.newpassword).then(res => {
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else if (form.password == '') {
        pswWarningMsg()
    } else {
        updateWarningMsg()
    }
}
const changeEmail = () => {
    if (form.username != '' && form.newemail != '') {
        emailUpdate(form.username, form.newemail).then(res => {
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else {
        updateWarningMsg()
    }
}
const changeAge = () => {
    if (form.username != '' && form.newage != '') {
        ageUpdate(form.username, form.newage).then(res => {
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else {
        updateWarningMsg()
    }
}
const changeSummary = () => {
    if (form.username != '' && form.newsummary != '') {
        summaryUpdate(form.username, form.newsummary).then(res => {
            if (res.code == 200) {
                reload()
                updateSuccessMsg()
            } else {
                errorMsg()
            }
        })
    } else {
        updateWarningMsg()
    }
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
    <div class="profile">
        <div class="title">个人资料:</div>
        <div class="content">
            <div class="info">
                <div class="uid">UID&nbsp;:&nbsp;{{ userInfo.ID }}</div>
                <div class="username">用户名&nbsp;:&nbsp;{{ userInfo.username }}</div>
                <div class="email">邮箱&nbsp;:&nbsp;{{ userInfo.email }}</div>
                <div class="age">年龄&nbsp;:&nbsp;{{ userInfo.age }}</div>
                <div class="summary">简介&nbsp;:&nbsp;{{ userInfo.summary }}</div>
            </div>
            <div class="update-box">
                <div class="form origin-msg">
                    <!-- <div class="oldName">用户名：{{ form.oldUsername }}</div> -->
                    <!-- <input type="text" placeholder="原用户名" v-model="form.username"> -->
                    <input type="password" placeholder="输入密码" v-model="form.password">
                </div>
                <div class="form update-username">
                    <input type="text" placeholder="新的用户名" v-model="form.newusername">
                    <button class="changebtn" @click="changeUsername">修改</button>
                </div>
                <div class="form update-password">
                    <input type="password" placeholder="新密码" v-model="form.newpassword">
                    <button class="changebtn" @click="changePassword">修改</button>
                </div>
                <div class="form update-email">
                    <input type="text" placeholder="新的邮箱" v-model="form.newemail">
                    <button class="changebtn" @click="changeEmail">修改</button>
                </div>
                <div class="form update-age">
                    <input type="text" placeholder="新的年龄" v-model="form.newage">
                    <button class="changebtn" @click="changeAge">修改</button>
                </div>
                <div class="form update-summary">
                    <input type="text" placeholder="新的简介" v-model="form.newsummary">
                    <button class="changebtn" @click="changeSummary">修改</button>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.profile {
    width: 100%;
    padding: 15px 25px;
    background: var(--bg1);
    border-radius: 5px;

    .title {
        width: 100%;
        font-size: 20px;
        font-weight: bold;
        padding: 10px 5px;
    }

    .content {
        width: 100%;
        display: flex;
        flex-direction: row;
    }

    .info {
        width: 500px;
        font-size: 18px;
        padding: 0 20px;
        gap: 5px;
        display: flex;
        flex-direction: column;
        gap: 25px;
        margin-top: 70px;

        .username {
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

    .update-box {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        color: var(--text-color1);
        // width: 100%;

        .form {
            width: 100%;
            padding: 10px 0;
            display: flex;
            // align-items: center;
            justify-content: center;
            gap: 20px;

            input {
                width: 180px;
                height: 40px;
                border: none;
                outline: none;
                border-radius: 10px;
                padding-left: 1em;
                background-color: #f0f0f0;
                box-shadow: inset 2px 2px 4px #d1d9e6, inset -2px -2px 4px #f9f9f9;
                // box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
                // margin: 0 20px;
            }

            .oldName {
                // width: 120px;
                height: 40px;
                display: flex;
                align-items: center;
                margin: 0 10px 0 20px;
            }

            .changebtn {
                width: 86px;
                height: 40px;
                border-radius: 24px;
                border: none;
                outline: none;
                background-color: $primary-200;
                color: #fff;
                font-size: 0.9em;
                cursor: pointer;
                box-shadow: 8px 8px 16px #d1d9e6, -8px -8px 16px #f9f9f9;
            }
        }
    }
}
</style>