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
        url: `/api/discuss/list?category=${category}`,
        method: 'get',
    })
}
// 根据query参数category获取讨论列表
export const getDiscussListByCategoryAndPage = (category: string, page: number) => {
    return httpInstance({
        url: `/api/discuss/list?category=${category}&page=${page}`,
        method: 'get',
    })
}
// 根据query参数id获取讨论详情
export const getDiscussDetailById = (id: any) => {
    return httpInstance({
        url: `/api/discuss/info?id=${id}`,
        method: 'get',
    })
}

// get the top topic of discuss list (from pre to end )
export const getDiscussTop = (pre: number, end: number) => {
    return httpInstance({
        url: `/api/discuss/toplist?pre=${pre}&end=${end}`,
        method: 'get',
    })
}

// add discuss view number by discuss id
export const addDiscussViewNumber = (id: any) => {
    return httpInstance({
        url: `/api/discuss/info/view?id=${id}`,
        method: 'put',
    })
}

// add discuss like number by discuss id
export const addDiscussLikeNumber = (id: any) => {
    return httpInstance({
        url: `/api/discuss/info/like?id=${id}`,
        method: 'put',
    })
}

// get mine discuss list
export const getMineDiscussList = (username: string) => {
    return httpInstance({
        url: `/api/discuss/minelist?username=${username}`,
        method: 'get',
    })
}