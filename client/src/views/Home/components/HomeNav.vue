<script setup lang="ts">

</script>

<template>
    <div class="home-nav">
        <div class="home-nav-list">
            <router-link to="/" :class="{ 'active': $route.path === '/' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-_meitiyunying"></use>
                </svg>
                综合
            </router-link>
            <router-link to="/following" :class="{ 'active': $route.path === '/following' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-wodeguanzhu"></use>
                </svg>
                关注
            </router-link>
            <router-link to="/frontend" :class="{ 'active': $route.path === '/frontend' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-_qianduankaifa"></use>
                </svg>
                前端
            </router-link>
            <router-link to="/backend" :class="{ 'active': $route.path === '/backend' }">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-_houduankaifa"></use>
                </svg>
                后端
            </router-link>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.home-nav {
    width: 150px;
    height: 100%;

    .home-nav-list {
        background: var(--bg1);
        border-radius: 5px;
        padding: 10px 0;

        a {
            width: 130px;
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
                margin-right: 15px;
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