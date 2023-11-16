import httpInstance from '@/utils/http'
export const getArticleList = () => {
    return httpInstance({
        url: '/article/list',
        method: 'get',
    })
}
