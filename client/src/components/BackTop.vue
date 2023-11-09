<script setup lang="ts">
import { ref } from 'vue'
import { useScroll } from '@vueuse/core'
const { y } = useScroll(window)
const scrollY = ref(200)
// 自定义样式
const customStyle = ref({
    width: '50px',
    height: '50px',
    bottom: '50px',
    right: '50px',
    backgroundColor: 'var(--primary-200)',
})

const handleClick = () => {
    window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'smooth'
    })
}

</script>
<template>
    <div class="backtop" :class="{ hidden: y < scrollY }" @click="handleClick" :style="{
        '--width': customStyle.width,
        '--height': customStyle.height,
        '--bottom': customStyle.bottom,
        '--right': customStyle.right,
        'background-color': customStyle.backgroundColor
    }">
        A
    </div>
</template>
<style scoped lang="scss">
.backtop {
    position: fixed;
    right: var(--right);
    bottom: var(--bottom);
    z-index: 999;
    width: var(--width);
    height: var(--height);
    border-radius: 50%;
    background-color: var(--background-color);
    color: #fff;
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