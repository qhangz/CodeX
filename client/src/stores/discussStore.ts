import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDiscussListStore = defineStore('discuss', () => {
    const index = ref(1)
    function increment() {
        index.value = index.value % 5 + 1;
    }

    return {
        index,
        increment,
    }
})