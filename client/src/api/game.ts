import httpInstance from '@/utils/http'

// get game list
export const getGameList = () => {
    return httpInstance({
        url: '/api/game/list',
        method: 'get',
    })
}