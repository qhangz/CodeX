<script setup lang="ts">
import { onMounted } from 'vue'

let greeting = ''
const setGreeting = () => {
    const currentTime = new Date().getHours();
    if (currentTime >= 0 && currentTime < 12) {
        greeting = '上午好！';
    } else if (currentTime >= 12 && currentTime < 18) {
        greeting = '下午好！';
    } else {
        greeting = '晚上好!';
    }
}
setGreeting()
</script>
<template>
    <div class="greeting-aside">
        <div class="greeting-word">
            {{ greeting }}
        </div>
        <div class="other">点亮在社区的每一天</div>
    </div>
</template>

<style lang="scss" scoped>
.greeting-aside {
    width: 100%;
    background: var(--bg1);
    border-radius: 5px;
    padding: 15px 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;

    .greeting-word {
        font-size: 20px;
        font-weight: bold;
        color: var(--text-color1);
        transition: all .2s linear;

        &:hover {
            padding: 0px 10px;
            // transform: translate3d(0, -3.5px, 0);
            // box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
            border-radius: 5px;
            background-color: var(--bg5);
            color: var(--primary-100);
        }
    }

    .other {
        font-size: 14px;
        color: var(--text-color1)
    }
}
</style>