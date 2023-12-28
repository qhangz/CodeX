<script setup lang="ts">

</script>

<template>
    <div class="home-nav">
        <div class="home-nav-list">
            <router-link to="/user/setting/profile" :class="{ 'active': $route.path === '/user/setting/profile' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-gerenziliao"></use>
                </svg>
                个人资料
            </router-link>
            <router-link to="/user/setting/account" :class="{ 'active': $route.path === '/user/setting/account' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-zhanghuguanli"></use>
                </svg>
                账户设置
            </router-link>
            <router-link to="/user/setting/common" :class="{ 'active': $route.path === '/user/setting/common' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-shezhi"></use>
                </svg>
                通用设置
            </router-link>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.home-nav {
    width: 190px;
    // height: 100%;

    .home-nav-list {
        background: var(--bg1);
        border-radius: 5px;
        padding: 10px 0;

        a {
            width: 170px;
            height: 45px;
            display: flex;
            // justify-content: center;
            align-items: center;
            flex-direction: row;
            color: var(--text-color1);
            margin: 0 10px;
            padding-left: 15px;
            border-radius: 5px;
            color: var(--text-color3);

            .icon {
                margin-right: 25px;
                font-size: 20px;
            }
        }

        .active {
            background-color: var(--bg5);
            color: var(--primary-100);
        }
    }
}

@media(max-width:935px) {
    .home-nav {
        display: none;
    }
}
</style>