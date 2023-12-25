<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getGameList } from '@/api/game'
const router = useRouter()

const loginWaring = () => {
    ElMessage({
        message: '未完待续...',
        type: 'warning',
    })
}

const innerClick = (item) => {
    if (item.route!==''){
        router.push(item.route)
    }else{
        loginWaring()
    }
}

interface GameList {
    Route: string;
    banner_image: string;
    game_name: string;
}
const gameList = ref<GameList[]>([])
const getGameInfo = async () => {
    const res = await getGameList()
    gameList.value = res.data
}

onMounted(() => {
    getGameInfo()
})
</script>
<template>
    <div class="game-list">
        <div class="content">
            <div class="title">
                游戏列表
            </div>
            <div class="game-item" v-for="(item, index) in gameList" :key="index">
                <div class="inner" @click="innerClick(item)">
                    <img :src="item.banner_image" alt="" class="banner">
                    <div class="message">
                        {{ item.game_name }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style scoped lang="scss">
.game-list {
    // width: 100%;
    max-width: 960px;

    .content {
        margin-top: 25px;
        display: flex;
        flex-wrap: wrap;
        align-items: stretch;
        margin-left: -7px;
        margin-right: -7px;

        .title {
            font-size: 20px;
            color: var(--text-color1);
            width: 100%;
            height: 50px;
            text-align: center;
        }

        .game-item {
            height: 270px;
            width: 25%;
            padding-left: 10px;
            padding-right: 10px;
            box-sizing: border-box;
            position: relative;
            margin-bottom: 35px;
            color: var(--text-color1);

            .inner {
                cursor: pointer;
                background-color: var(--bg1);
                box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
                transition: all .2s linear;
                border-radius: 5%;
                // border-top-left-radius: 20%;
                padding: 10px 15px 10px;

                .banner {
                    width: 100%;
                    height: 100%;
                    border-radius: 5%;
                    // border-top-left-radius: 20%;
                }

                .message {
                    margin: 20px 0 10px 0;
                    padding: 0 5px;
                    font-size: 18px;
                    color: var(--text-color1);
                    text-align: center;
                    overflow: hidden;
                    word-break: break-all;
                    text-overflow: ellipsis;
                    display: -webkit-box;
                    -webkit-box-orient: vertical;
                    -webkit-line-clamp: 2;
                }

                &:hover {
                    content: "";
                    transform: translate3d(0, -3.5px, 0);
                    box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
                }
            }
        }
    }
}
</style>