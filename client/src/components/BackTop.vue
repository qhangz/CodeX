<!-- <script setup lang="ts">

</script>
<template>
    <div class="backtop">
        <a href="javascript:;" class="backtop backtop-hidden">
            <i class="iconfont icon-up"></i>
        </a>
    </div>
</template>
<style scoped lang="scss"></style> -->
<!-- backtop 的组件 -->
<script setup lang="ts">
import { ref } from 'vue'
import { useScroll } from '@vueuse/core'
const { y } = useScroll(window)
const scrollY = ref(200)

const handleClick = () => {
    window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'smooth'
    })
}


</script>
<template>
    <div class="backtop"  :class="{ hidden: y < scrollY }" @click="handleClick">
        A
    </div>
</template>
<style scoped lang="scss">
.backtop {
    position: fixed;
    right: 50px;
    bottom: 50px;
    z-index: 999;
    width: 50px;    
    height: 50px;
    border-radius: 50%;
    background-color: var(--primary-200);
    color:#fff;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    transition: all 0.3s ease-in-out;

    &.hidden {
        opacity: 0;
        visibility: hidden;
    }
}
</style>