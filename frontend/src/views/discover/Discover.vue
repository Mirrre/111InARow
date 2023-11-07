<template>
    <div class="discover">
        <div class="video_content">
            <el-row :gutter="16">
                <el-col
                    :span="4"
                    v-for="(item, index) in videoArr"
                    :key="index">
                    <div class="grid-content bg-purple">
                        <div class="img_box" 
                        @mouseover="videoPlayback(index, videoArr[index].video_url)" 
                        @mouseout="videoStopped(index, videoArr[index].video_url)"
                        @click="toPlayer(item)">
                            <img :src="videoArr[index].thumb" alt="" ref="img"/>
                            <!-- contextmenu(禁止右键) muted(解决谷歌浏览器不能自动播放)-->
                            <div class="video_box">
                                <video id="video"
                                ref="video"
                                width="100%"
                                height="100%"
                                @contextmenu.prevent
                                controls
                                controlslist="nodownload nofullscreen noremote footbar"
                                disablePictureInPicture
                                muted="muted"
                                style="display: none; object-fit: fill">
                                <!-- :poster="videoArr[index].thumb"> -->
                                <source type="audio/mp4" :src="videoArr[index].video_url" />
                                </video>
                            </div>
                            <span class="video_time">{{ videoTime }}</span>
                        </div>
                        <!-- 标题 -->
                        <div class="tit">
                            <span>{{ videoArr[index].title }}</span>
                        </div>
                        <!-- 用户及视频上传时间 -->
                        <div class="actcont-auto">
                            <img :src="videoArr[index].avatar" alt="">
                            <span class="uname">{{ videoArr[index].username }}</span>
                            <span class="mtime">{{ videoArr[index].mtime }}</span>
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            videoArr: [],
            videoTime: "",
            timer: ""
        };
    },
     mounted() {
        this.videoList();
        setInterval(()=>{
            this.videoList();
        },1000)
    },
    methods: {
        videoList() {
            // this.$http.video.allList().then(({ data }) => {
            //     console.log(data);
            //     this.videoArr = data;
            //     // this.videoArr = data.list;
            //     console.log(this.videoArr);
            // });
            let data1 = {}
            const userMsg = JSON.parse(localStorage.getItem('userMsg'))
            const isLogin = localStorage.getItem('isLogin')
            // console.log(userMsg.id);
            if(isLogin){
                data1 = {
                    user_id: userMsg.id,
                    has: 1
                }
            }
            this.$http.video.allList(data1).then(({ code, data }) => {
                if (code === 200) {
                    // console.log(data);
                    this.videoArr = data;
                    // console.log(this.videoArr);
                }
            });
            
        },
        videoPlayback(index, video_url) {
            const video = this.$refs.video[index]
            if(video_url) {
                // 显示播放器
                video.style.display = "block"
                // 隐藏图片
                this.$refs.img[index].style.display = "none"
                // 定时器
                this.timer = setTimeout(() => {
                    // 播放视频
                    video.play()
                    // 给视频标签添加缓存播放---video标签属性。
                    video.setAttribute("autoplay", "autoplay")
                    // 给视频标签添加循环播放---video标签属性。
                    video.setAttribute("loop", "loop")
                    // 初始播放音量
                    video.volume = 0
                }, 200);
                // 视频当前的播放时间(进度)
                // setInterval(() => {
                //     console.log(video.currentTime);
                // }, 0);
                // 获取视频时长
                // let minutes = parseInt(video.duration / 60); // 分
                // let seconds = parseInt(video.duration % 60); // 秒
                // minutes >= 10 ? minutes = minutes : minutes = "0" + minutes
                // seconds >= 10 ? seconds = seconds : seconds = "0" + seconds
                // this.videoTime = minutes + ":" + seconds
                // console.log(this.videoTime);
            }
        },
        videoStopped(index, video_url) {
            const video = this.$refs.video[index]
            if(video_url) {
                // 清除定时器
                clearTimeout(this.timer)
                //停止播放
                video.pause()
                // 重置播放时间
                video.currentTime = 0
                // 隐藏播放器
                video.style.display = "none"
                // 隐藏图片
                this.$refs.img[index].style.display = "block"
            }
        },
        toPlayer(obj) {
            this.$router.push({
                name: "Player",
                query: { id: obj.id},
            })
            sessionStorage.setItem("rowData", JSON.stringify(obj))
        }
    },
    created() {
        this.videoList();
    },
};
</script>

<style lang="less" scoped>
.discover {
    width: 100%;
    height: 100%;
    .video_content {
        width: 100%;
        height: 100%;
    }
}
.el-row {
    margin-bottom: 20px;
    &:last-child {
        margin-bottom: 0;
    }
}
.el-col {
    border-radius: 4px;
    margin-bottom: 24px;
}
.bg-purple-dark {
    background: #99a9bf;
}
.bg-purple {
    background: none;
}
.bg-purple-light {
    background: #e5e9f2;
}
.grid-content {
    border-radius: 4px;
    min-height: 36px;
    overflow: hidden;
}
.row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
}
.el-col-4 {
    width: 20%;
}
// 视频信息样式
.img_box {
    width: 100%;
    max-height: 240px;
    border-radius: 4px;
    overflow: hidden;
    cursor: pointer;
    img, .video_box {
        width: 280px;
        height: 300px;
        object-fit: cover;
    }
    img {
        // 解决图片底部默认空白缝隙问题
        display: block;
    }
}
.tit {
    width: 100%;
    height: 46px;
    margin-top: 5px;
    overflow: hidden;
    text-overflow:ellipsis;
    span {
        font-size: 16px;
        line-height: 24px;
        color: black;
    }
    span:hover {
        color: black;
    }
}
.actcont-auto {
    position: relative;
    width: 100%;
    height: 34px;
    margin-top: 8px;
    line-height: 34px;
    img {
        width: 34px;
        height: 34px;
        border-radius: 50%;
    }
    .uname {
        display: inline-block;
        max-width: 120px;
        height: 100%;
        color: black;
        overflow: hidden;
        text-overflow:ellipsis;
        padding-left: 10px;
    }
    .uname:hover {
        color: black;
    }
    .mtime {
        max-width: 120px;
        height: 100%;
        position: absolute;
        top: 0;
        right: 0;
        color: rgba(255,255,255,.5);
        overflow: hidden;
    }
}
// 播放器样式
video::-webkit-media-controls-timeline { // 进度条
    width: 100%;
    padding: 0;
}
//观看的当前时间
video::-webkit-media-controls-current-time-display{
    display: none !important;
}
//剩余时间
video::-webkit-media-controls-time-remaining-display {
    display: none !important;
}
//音量按钮
video::-webkit-media-controls-mute-button {
    display: none !important;
}
video::-webkit-media-controls-toggle-closed-captions-button {
    display: none !important;
}
//音量的控制条
video::-webkit-media-controls-volume-slider {
    display: none !important;        
}
// 全屏按钮
video::-webkit-media-controls-fullscreen-button {
    display: none !important;
}
//播放按钮
video::-webkit-media-controls-play-button {
    display: none !important;
}
// 播放控件
video::-webkit-media-controls { // 进度条
    opacity: 1;
    // 让鼠标指针一直显示
    cursor: pointer !important;
}
</style>