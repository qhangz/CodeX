<script setup lang="ts">

</script>
<template>
    <div class="codex-app">
        <div class="left">
            <div class="title">
                <img src="@/assets/images/logo.png" alt="logo" class="logo">
                <div class="name">
                    CodeX
                </div>
            </div>
            <div class="content">一个帮助开发者学习成长的社区</div>
            <img src="@/assets/images/qrcode.jpg" alt="ar-code" class="QR-code">
            <div class="text">扫码关注公众号</div>
        </div>
        <div class="right">
            <div class="images">
                <img src="@/assets/images/OIG.jpg" alt="img0" class="img0">
                <img src="@/assets/images/OIG1.jpg" alt="img1" class="img1">
                <img src="@/assets/images/codex0.png" alt="codex0" class="codex0">
                <img src="@/assets/images/codex3.png" alt="codex3" class="codex3">
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
.codex-app {
    width: 100%;
    min-height: calc(100vh - 60px);
    display: flex;
    justify-content: center;
    flex-direction: row;
    gap: 500px;

    margin: 0;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(45deg, #b3e5fc, #b2dfdb, #b3e5fc, #b2dfdb);
    // background: linear-gradient(45deg, #ff00ff, #00ffff, #ff00ff, #00ffff);
    background-size: 400% 400%;
    animation: gradientAnimation 5s ease infinite;

    @keyframes gradientAnimation {
        0% {
            background-position: 0% 50%;
        }

        50% {
            background-position: 100% 50%;
        }

        100% {
            background-position: 0% 50%;
        }
    }

    .left {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        transform: translate(0px, -50px);

        .title {
            display: flex;
            flex-direction: row;
            gap: 20px;

            .logo {
                width: 100px;
                border-radius: 50%;
                margin-right: 20px;
            }

            .name {
                font-size: 30px;
                color: var(--text-color1);
                font-weight: bold;
                display: flex;
                align-items: center;
            }
        }

        .content {
            font-size: 20px;
            color: var(--text-color1);
            display: flex;
            justify-content: center;
            align-items: center;
            flex-direction: column;
            padding: 24px 0 48px 0;
        }

        .QR-code {
            width: 240px;
            height: 240px;
            border-radius: 20px;
        }

        .text {
            font-size: 24px;
            color: var(--text-color1);
            padding-top: 15px;
        }

    }

    .right {
        .images {
            width: 230px;
            height: 230px;

            .img0 {
                transform: translate(0px, -255px);
                border-radius: 20px;
            }

            .img1 {
                transform: translate(130px, -230px);
                border-radius: 20px;
            }

            .codex0 {
                transform: translate(0px, -200px);
                border-radius: 20px;
            }

            .codex3 {
                transform: translate(-130px, -690px);
                border-radius: 20px;
            }
        }
    }
}

@media (min-width: 1024px) {
    .app {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
}
</style>
```