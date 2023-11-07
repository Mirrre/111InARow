import axios from '../http/http'
const collect = {
    Collect(data){
        //点赞
        return axios.post('/collect_video/addCollectVideo', data)
    },
    isCollect(data){
        //当前用户是否已经点赞，如点赞，则为变红星，is_follow变1
        return axios.post('/collect_video/getCollectVideoByLimit', data)
    },
    deleCollect(data){
        return axios.post('/collect_video/deleteCollectVideoByLimit', data)
    }
} 
export default collect


