import httpInstance from '@/utils/http'

// user register
export const userRegister = (username: string, password: any, email: any) => {
    return httpInstance({
        url: '/api/user/register',
        method: 'POST',
        data: {
            username,
            password,
            email,
        }
    })
}
// user login
export const userLogin = (username: string, password: any) => {
    return httpInstance({
        url: '/api/user/login',
        method:'POST',
        data:{
            username,
            password,
        }
    })
}