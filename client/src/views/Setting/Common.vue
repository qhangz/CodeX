<script setup lang="ts">

</script>
<template>
    <div class="common">
        <div class="title">通用设置:</div>
        <hr class="divider">
        <div class="content">
            <h4>
                页面设置
            </h4>
            <h5>
                该设置仅在当前浏览器生效，目前已支持部分核心页面，更多页面适配持续进行中，如有建议可点击页面右下角反馈
            </h5>
            <img src="@/assets/images/commonSetting.jpg" alt="img">
        </div>
    </div>
</template>
<style lang="scss" scoped>
.common {
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

    .divider {
        width: 100%;
        opacity: 70%;
        border-bottom: 1px solid var(--primary-700);
    }

}
</style>