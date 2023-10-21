<script setup lang="ts">
import { reactive } from 'vue'
const state = reactive({
    isLogin: false,
    emailError: false,
    passwordError: false,
    existed: false,

})
const form = reactive({
    username: '',
    useremail: '',
    userpwd: ''
})

// login function
const login = () => {
    console.log('login');
    console.log(form);
}
// register function
const register = () => {
    console.log('register');
    console.log(form);
}
// changeType function
const changeType = () => {
    state.isLogin = !state.isLogin
    form.username = ''
    form.useremail = ''
    form.userpwd = ''
}
</script>

<template>
    <div class="login">
        <div class="contain">
            <div class="big-box" :class="{ active: state.isLogin }">
                <div class="big-contain" v-if="state.isLogin">
                    <div class="btitle">账户登录</div>
                    <div class="bform">
                        <input type="text" placeholder="用户名" v-model="form.username">
                        <!-- <span class="errTips" v-if="state.emailError">* 邮箱填写错误 *</span> -->
                        <input type="password" placeholder="密码" v-model="form.userpwd">
                        <span class="errTips" v-if="state.passwordError">* 密码填写错误 *</span>
                    </div>
                    <button class="bbutton" @click="login">登录</button>
                </div>
                <div class="big-contain" key="bigContainRegister" v-else>
                    <div class="btitle">创建账户</div>
                    <div class="bform">
                        <input type="text" placeholder="用户名" v-model="form.username">
                        <span class="errTips" v-if="state.existed">* 用户名已经存在！ *</span>
                        <input type="email" placeholder="邮箱" v-model="form.useremail">
                        <input type="password" placeholder="密码" v-model="form.userpwd">
                    </div>
                    <button class="bbutton" @click="register">注册</button>
                </div>
            </div>

            <div class="small-box" :class="{ active: state.isLogin }">
                <div class="small-contain" key="smallContainRegister" v-if="state.isLogin">
                    <div class="stitle">你好，朋友!</div>
                    <p class="scontent">开始注册，和我们一起旅行</p>
                    <button class="sbutton" @click="changeType">注册</button>
                </div>
                <div class="small-contain" key="smallContainLogin" v-else>
                    <div class="stitle">欢迎回来!</div>
                    <p class="scontent">去登入账号来进入奇妙世界吧！！！</p>
                    <button class="sbutton" @click="changeType">登录</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.login {
    width: 90vw;
    height: 88vh;
    padding: 2rem;
    box-sizing: border-box;

    .contain {
        width: 60%;
        height: 60%;
        position: relative;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        background-color: #fff;
        border-radius: 20px;
        // box-shadow: 0 0 3px #f0f0f0,0 0 6px #f0f0f0;
        box-shadow: 10px 10px 10px #d1d9e6, -10px -10px 10px #f9f9f9;

        .big-box {
            width: 70%;
            height: 100%;
            position: absolute;
            top: 0;
            left: 30%;
            transform: translateX(0%);
            transition: all 1s;

            .big-contain {
                width: 100%;
                height: 100%;
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;

                .btitle {
                    font-size: 1.5em;
                    font-weight: bold;
                    color: rgb(57, 167, 176);
                }

                .bform {
                    width: 100%;
                    height: 40%;
                    padding: 2em 0;
                    display: flex;
                    flex-direction: column;
                    justify-content: space-around;
                    align-items: center;

                    .errTips {
                        display: block;
                        width: 50%;
                        text-align: left;
                        color: red;
                        font-size: 0.7em;
                        margin-left: 1em;
                    }
                }

                .bform input {
                    width: 50%;
                    height: 30px;
                    border: none;
                    outline: none;
                    border-radius: 10px;
                    padding-left: 2em;
                    background-color: #f0f0f0;
                    box-shadow: inset 2px 2px 4px #d1d9e6, inset -2px -2px 4px #f9f9f9;
                }

                .bbutton {
                    width: 20%;
                    height: 40px;
                    border-radius: 24px;
                    border: none;
                    outline: none;
                    background-color: rgb(57, 167, 176);
                    color: #fff;
                    font-size: 0.9em;
                    cursor: pointer;
                    box-shadow: 8px 8px 16px #d1d9e6, -8px -8px 16px #f9f9f9;
                }

            }
        }

        .big-box.active {
            left: 0;
            transition: all 0.5s;
        }

        .small-box {
            width: 30%;
            height: 100%;
            background: linear-gradient(135deg, rgb(57, 167, 176), rgb(56, 183, 145));
            position: absolute;
            top: 0;
            left: 0;
            transform: translateX(0%);
            transition: all 1s;
            border-top-left-radius: inherit;
            border-bottom-left-radius: inherit;

            .small-contain {
                width: 100%;
                height: 100%;
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;

                .stitle {
                    font-size: 1.5em;
                    font-weight: bold;
                    color: #fff;
                }

                .scontent {
                    font-size: 0.8em;
                    color: #fff;
                    text-align: center;
                    padding: 2em 3.5em;
                    line-height: 1.7em;
                }

                .sbutton {
                    width: 60%;
                    height: 40px;
                    border-radius: 24px;
                    border: 1px solid #fff;
                    outline: none;
                    background-color: transparent;
                    color: #fff;
                    font-size: 0.9em;
                    cursor: pointer;

                }
            }
        }

        .small-box.active {
            left: 100%;
            border-top-left-radius: 0;
            border-bottom-left-radius: 0;
            border-top-right-radius: inherit;
            border-bottom-right-radius: inherit;
            transform: translateX(-100%);
            transition: all 1s;
        }
    }
}
</style>