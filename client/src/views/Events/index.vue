<script setup lang="ts">
import EventList from './components/EventList.vue'
</script>

<template>
    <div class="events">
        <EventList/>
    </div>
</template>

<style lang="scss" scoped>
.events {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    // padding-top: 90px;
    gap: 20px;
}

// @media (min-width: 1024px) {
//     .events {
//         display: flex;
//         justify-content: center;
//         align-items: center;
//         flex-direction: column;
//     }
// }
</style>