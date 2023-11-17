import httpInstance from '@/utils/http'
// 获取热门文章
export const getArticleHot = () => {
    return httpInstance({
        url: '/article/hot',
        method: 'get',
    })
}

// 获取关注文章
export const getArticleFollow = () => {
    return httpInstance({
        url: '/article/follow',
        method: 'get',
    })
}

// 获取前端文章
export const getArticleFrontend = () => {
    return httpInstance({
        url: '/article/frontend',
        method: 'get',
    })
}

// 获取后端文章
export const getArticleBackend = () => {
    return httpInstance({
        url: '/article/backend',
        method: 'get',
    })
}
