<script setup lang="ts">

</script>

<template>
    <div class="android">
        this is android page

    </div>
</template>

<style lang="scss" scoped>
.android {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}

@media (min-width: 1024px) {
    .android {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
}
</style>