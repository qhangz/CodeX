import httpInstance from '@/utils/http'

// get game list
export const getEventList = () => {
    return httpInstance({
        url: '/api/event/list',
        method: 'get',
    })
}