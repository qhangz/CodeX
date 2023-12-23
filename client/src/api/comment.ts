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
