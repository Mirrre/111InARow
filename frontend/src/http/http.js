import axios from 'axios'

// 请求根路径
const baseURL = axios.create({ baseURL: 'http://localhost:8901/api' })

baseURL.interceptors.request.use((response) => {
    return response;
}, (error) => {
    // 处理错误响应
    return Promise.reject(error)
})
// 响应拦截器
baseURL.interceptors.response.use(
    // 请求成功
    (res) => {
        if (res.status === 200) {
            return Promise.resolve(res.data)
        } else {
            return Promise.reject(res)
        }
    },
    // 请求失败
    (error) => {
        return Promise.reject(error)
    })

export default baseURL