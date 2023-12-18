<script setup lang="ts">
import GreetingAside from './GreetingAside.vue'
</script>
<template>
    <div class="home-aside">
        <GreetingAside />
        
    </div>
</template>

<style lang="scss" scoped>
.home-aside {
    width: 260px;
    height: 100%;
    // background-color: var(--primary-300);
    // padding: 15px 20px;
}

@media(max-width:1200px) {
    .home-aside {
        display: none;
    }
}
</style>