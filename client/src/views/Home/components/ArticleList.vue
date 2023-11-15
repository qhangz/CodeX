<script setup lang="ts">

</script>

<template>
    <div class="article-list">
        article list
    </div>
</template>

<style lang="scss" scoped>
.article-list{
    width: 720px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color:bisque;
}
@media(max-width:730px){
    .article-list{
        width:95vw;
        // height:100vh;
    }
}
</style>
