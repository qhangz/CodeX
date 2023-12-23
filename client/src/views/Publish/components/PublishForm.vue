<script setup lang="ts">
import { reactive } from 'vue'
import { useUserStore } from '@/stores/userStore';
import { publishDiscuss } from '@/api/discuss'

// do not use same name with ref
const form = reactive({
    title: '',
    summary: '',
    category: '',
    content: '',
})

// submit botton
const submitWaring = () => {
    ElMessage({
        message: '请完成信息输入',
        type: 'warning',
    })
}
const submitSuccess = () => {
    ElMessage({
        message: '发布成功',
        type: 'success',
    })
}
const submitFail = () => {
    ElMessage.error('发布失败')
}
const userNotLogin = () => {
    ElMessage.error('用户未登录')
}
const onSubmit = () => {
    if (form.title != '' && form.summary != '' && form.category != '' && form.content != '') {
        if (localStorage.getItem('isLogin')) {
            let author = localStorage.getItem('userInfo') ? JSON.parse(localStorage.getItem('userInfo')!).username : ''
            const res = publishDiscuss(author, form.title, form.summary, form.content, form.category).then(res => {
                if (res.code == 200) {
                    submitSuccess()
                } else {
                    submitFail()
                }
            })
        } else {
            userNotLogin()
        }
    } else {
        submitWaring()
    }
}

// cancel botton
const onCancel = () => {
    form.title = ''
    form.summary = ''
    form.category = ''
    form.content = ''
}
</script>

<template>
    <div class="publish-form">
        <div class="top-title">Come and share what's new with code explorer!</div>
        <el-form :model="form" label-width="120px" size="large" label-position="left">
            <el-form-item label="Title">
                <el-input v-model="form.title" placeholder="Please input title" />
            </el-form-item>
            <el-form-item label="Summary">
                <el-input v-model="form.summary" placeholder="Please input summary" />
            </el-form-item>
            <el-form-item label="Category">
                <el-radio-group v-model="form.category">
                    <el-radio border label="hot" />
                    <el-radio border label="follow" />
                    <el-radio border label="frontend" />
                    <el-radio border label="backend" />
                </el-radio-group>
            </el-form-item>
            <el-form-item label="Content">
                <el-input v-model="form.content" type="textarea" :autosize="{ minRows: 5, maxRows: 7 }"
                    placeholder="Please input content" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">Publish</el-button>
                <el-button @click="onCancel">Cancel</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<style lang="scss" scoped>
.publish-form {
    display: flex;
    align-items: center;
    flex-direction: column;
    // min-height: calc(100vh - 188px);
    // height: calc(100vh - 60px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 1180px;
    justify-content: center;

    .top-title {
        font-size: 25px;
        font-weight: 600;
        margin-bottom: 50px;
    }

    .el-form {
        .el-form-item {
            font-weight: 600;
            margin-bottom: 30px;
            display: flex;
            align-items: center;
            justify-content: center;

            .el-button {
                display: flex;
                align-items: center;
                margin-left: 65px;
            }
        }
    }
}
</style>
```