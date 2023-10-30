<script setup lang="ts">

</script>

<template>
    <div class="home">
        
        this is home page
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    div{
        width: 100px;
        height: 100px;
        background-color: red;
        margin: 10px;
    }
}

@media (min-width: 1024px) {
    .home {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }
}
</style>