// import axios from '@/http/http.js'

// // 封装一系列 视频相关 的请求方法
// const video = {
//     // 热门视频
//     hotList(data) {
//         return axios.get('/hot', data)
//     },
// }

// export default video

import axios from '@/http/http.js'

// 封装一系列 视频相关 的请求方法
const video = {
    // 热门视频
    allList(data) {
        return axios.post('/all_video/findAllVideo', data)
    },
    hotList(data){
        return axios.post('/hot_video/getHotVideo', data)
    },
    physicalList(data){
        return axios.post('/physical_video/getPhysicalVideo', data)
    },
    animeList(data){
        return axios.post('/anime_video/getAnimeVideo', data)
    },
    gameList(data){
        return axios.post('/game_video/getGameVideo', data)
    },
    findAllVideoByTitle(data){
        return axios.post('/all_video/findAllVideoByTitle', data)
    },
    findUploadVideoById(data){
        return axios.post('/all_video/findUploadVideoById', data)
    },
}

export default video