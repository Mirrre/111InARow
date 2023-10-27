import axios from '@/http/http.js'

// 封装一系列 视频相关 的请求方法
const video = {
    // 热门视频
    hotList(data) {
        return axios.get('/hot', data)
    },
}

export default video