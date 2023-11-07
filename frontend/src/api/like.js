import axios from '../http/http'
const like = {
    Like(data){
        //点赞
        return axios.post('/like_video/addLikeVideo', data)
    },
    isLike(data){
        //当前用户是否已经点赞，如点赞，则为变红星，is_follow变1
        return axios.post('/like_video/getLikeVideoByLimit', data)
    },
    deleLike(data){
        return axios.post('/like_video/deleteLikeVideoByLimit', data)
    },
    getLikeVideoByUser(data){
        return axios.post('/like_video/getLikeVideoByUser', data)
    },
} 
export default like


