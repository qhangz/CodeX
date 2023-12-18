import httpInstance from '@/utils/http'

// 发布discuss
export const publishDiscuss = (author: string, title: string, summary: string, content: string, category: string) => {
    return httpInstance({
        url: '/api/discuss/publish',
        method: 'post',
        data: {
            author,
            title,
            summary,
            content,
            category,
        }
    })
}

// 根据query参数category获取讨论列表
export const getDiscussListByCategory = (category: string) => {
    return httpInstance({
        url: `/api/discuss/discusslist?category=${category}`,
        method: 'get',
    })
}
// 根据query参数category获取讨论列表
export const getDiscussListByCategoryAndPage = (category: string, page: number) => {
    return httpInstance({
        url: `/api/discuss/discusslist?category=${category}&page=${page}`,
        method: 'get',
    })
}
// 根据query参数id获取讨论详情
export const getDiscussDetailById = (id: string) => {
    return httpInstance({
        url: `/api/discuss/discussdetail?id=${id}`,
        method: 'get',
    })
}