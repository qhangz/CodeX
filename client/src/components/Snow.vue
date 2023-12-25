<template>
    <div class="snow-container">
        <div v-for="snowflake in snowflakes" :key="snowflake.id" class="snowflake" :style="snowflake.style">
            {{ snowflake.text }}
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
const snowflakes = ref([]);
onMounted(() => {
    for (let i = 0; i <= 30; i++) {
        createSnowflake();
    }
});
onUnmounted(() => {
    snowflakes.value = [];
});
function createSnowflake() {
    const snowflake = {
        id: Date.now(),
        text: '❄️',
        style: {
            position: 'fixed',
            color: 'white',
            zIndex: 9999,
            top: '-10%',
            left: `${100 * Math.random()}%`,
            animation: 'snowflakes-fall 10s linear infinite, snowflakes-shake 3s ease-in-out infinite',
            fontSize: `${2 * Math.random()}em`,
            animationDelay: `${10 * Math.random()}s, ${2 * Math.random()}s`,
        },
    };
    snowflakes.value.push(snowflake);
}
</script>
<style>
@keyframes snowflakes-fall {
    0% {
        top: -10%;
    }

    100% {
        top: 100%;
    }
}

@keyframes snowflakes-shake {
    0% {
        transform: translateX(0px);
    }

    50% {
        transform: translateX(80px);
    }

    100% {
        transform: translateX(0px);
    }
}

.snow-container {
    position: fixed;
    width: 100%;
    height: 100%;
}

.snowflake {
    position: fixed;
    color: white;
    z-index: 9999;
    animation-name: snowflakes-fall, snowflakes-shake;
    animation-duration: 10s, 3s;
    animation-timing-function: linear, ease-in-out;
    animation-iteration-count: infinite, infinite;
    animation-play-state: running, running;
}
</style>