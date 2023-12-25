<script setup lang="ts">
import GameList from './components/GameList.vue'
</script>

<template>
    <div class="games">
        <GameList/>
    </div>
</template>

<style lang="scss" scoped>
.games {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
}

@media (min-width: 1024px) {
    .games {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
}
</style>