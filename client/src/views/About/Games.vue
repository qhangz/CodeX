<template>
    <div>
        <h1>{{ article.title }}</h1>
        <p>{{ article.content }}</p>
        <p>浏览量：{{ views }}</p>
    </div>
</template>
  
<script>
import { ref, onMounted } from 'vue';

export default {
    setup() {
        const article = {
            title: '文章标题',
            content: '文章内容',
        };
        const views = ref(0); // 使用 ref 创建响应式变量

        onMounted(() => {
            incrementViews(); // 组件加载后立即调用增加浏览量的方法
        });

        function incrementViews() {
            views.value++; // 每次调用该方法，浏览量加一
        }

        return {
            article,
            views,
        };
    },
};
</script>