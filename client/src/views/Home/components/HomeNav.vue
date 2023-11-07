<script setup lang="ts">

</script>

<template>
    <div class="home-nav">
        <ul class="home-nav">
            <li>
                <router-link to="/following">Following</router-link>
            </li>
            <li>
                <router-link to="/frontend">Frontend</router-link>
            </li>
            <li>
                <router-link to="/backend">Backend</router-link>
            </li>
        </ul>
    </div>
</template>

<style lang="scss" scoped>
.home-nav {
    width: 180px;
    height: 560px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color: var(--bg1);
}

// @media (min-width: 1024px) {
//     .home-nav {
//         display: flex;
//         justify-content: center;
//         align-items: center;
//         flex-direction: column;
//     }
// }
</style>