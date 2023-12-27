import httpInstance from '@/utils/http'

// submit comment 
export const submitComment = (discussid: number, author: string, content: string) => {
    return httpInstance({
        url: '/api/comment/publish',
        method: 'post',
        data: {
            author,
            content,
            discussid,
        }
    })
}

// add comment view number by comment id
export const addCommentViewNumber = (id: any) => {
    return httpInstance({
        url: `/api/comment/info/view?id=${id}`,
        method: 'put',
    })
}

// add comment like number by comment id
export const addCommentLikeNumber = (id: any) => {
    return httpInstance({
        url: `/api/comment/info/like?id=${id}`,
        method: 'put',
    })
}