<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getEventList } from '@/api/event'
const router = useRouter()

const loginWaring = () => {
    ElMessage({
        message: '未完待续...',
        type: 'warning',
    })
}

const innerClick = (item) => {
    if (item.route !== '') {
        router.push(item.route)
    } else {
        if (item.url !== '') {
            window.open(item.url)
        } else {
            loginWaring()
        }
    }
}

interface EventList {
    Route: string;
    url: string;
    name: string;
    time: string;
    place: string;
    status: string;
    banner_image: string;
}
const eventList = ref<EventList[]>([])
const getEventInfo = async () => {
    const res = await getEventList()
    eventList.value = res.data
}

onMounted(() => {
    getEventInfo()
})
</script>

<template>
    <div class="event-list">
        <div class="container">
            <div class="title">
                活动列表
            </div>
            <div class="event-item" v-for="(item, index) in eventList" :key="index">
                <div class="inner" @click="innerClick(item)">
                    <div class="banner" :style="{ backgroundImage: 'url(' + item.banner_image + ')' }"></div>
                    <div class="message">
                        <div class="name">{{ item.name }}</div>
                        <div class="time">
                            <svg class="icon" aria-hidden="true">
                                <use xlink:href="#icon-rili"></use>
                            </svg>
                            {{ item.time }}
                        </div>
                        <div class="place">
                            <svg class="icon" aria-hidden="true">
                                <use xlink:href="#icon-didiandingwei"></use>
                            </svg>
                            {{ item.place }}
                        </div>
                        <div class="detail" v-if="item.status == '已结束'">
                            已结束
                        </div>
                        <div class="join" v-if="item.status == '未开始'">
                            报名参加
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.event-list {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    // min-height: calc(100vh - 188px);
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 960px;

    .container {
        width: 100%;
        height: 100%;
        margin-top: 35px;
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

        .event-item {
            // width: 135px;
            height: 270px;
            width: 25%;
            padding-left: 7px;
            padding-right: 7px;
            box-sizing: border-box;
            position: relative;
            margin-bottom: 30px;
            color: var(--text-color1);

            .inner {
                background-color: var(--bg1);
                box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
                transition: all .2s linear;
                cursor: pointer;
                padding-bottom: 15px;

                .banner {
                    background-size: cover;
                    padding-top: 58.82%;
                    background-color: #ccc;
                    background-repeat: no-repeat;
                }

                .message {
                    padding: 15px 15px 0px;
                    // font-size: 16px;
                    // text-align: center;
                    // color: var(--text-color1);
                    background-color: var(--bg1);

                    .name {
                        color: var(--text-color1);
                        font-size: 16px;
                        display: -webkit-box;
                        overflow: hidden;
                        text-overflow: ellipsis;
                        -webkit-box-orient: vertical;
                        -webkit-line-clamp: 2;
                        height: 40px;
                        font-size: 14px;
                        font-weight: 700;
                        text-align: center;
                    }

                    .time {
                        // vertical-align: middle;
                        color: var(--text-color3);
                        font-size: 12px;
                    }

                    .place {
                        color: var(--text-color3);
                        font-size: 12px;
                    }

                    .join {
                        color: var(--text-color4);
                        font-size: 13px;
                        text-align: center;
                    }

                    .detail {
                        color: var(--text-color3);
                        font-size: 13px;
                        text-align: center;
                    }
                }

                &:hover {
                    content: "";
                    transform: translate3d(0, -3.5px, 0);
                    box-shadow: 1px 5px 8px rgb(0 0 0 / 20%);
                }
            }
        }
    }





    // .event-item:nth-child(4n+1) {
    //     background-color: red;
    // }
    // .event-item::before {
    //     content: "";
    //     position: absolute;
    //     left: 7px;
    //     right: 7px;
    //     top: 0;
    //     bottom: 0;
    //     z-index: -1;
    //     border-radius: 2px;
    //     overflow: hidden;
    //     background-color: red;
    //     box-shadow: 1px 1px 1px rgba(0, 0, 0, .15);
    //     transition: all .2s linear;
    // }
}
</style>
```