<template>
    <div class="short-video-player">
        <swiper-video 
            :videoList="videoList"
            :innerWidth="w"
            :innerHeight="h">
        </swiper-video>
        <!-- 关闭按钮 -->
        <div class="to-back" @click="$router.back(-1)">
            <i class="van-icon van-icon-cross"></i>
        </div>
    </div>
</template>

<script>
import SwiperVideo from "@/components/swiper/index.vue";

export default {
    components: {
        SwiperVideo,
    },
    data() {
        return {
            videoList: [],
            w: "100%",
            h: "100%"
        };
    },
     mounted() {
        this._videoList();
        setInterval(()=>{
            this._videoList();
        },300)
    },
    methods: {
        _videoList() {
            // this.$http.video.hotList().then(({ code, data }) => {
            //     console.log(data);
            //     if (code === 200) {
            //         this.videoList = this.videoList.concat(data.list);
            //     }
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
                    this.videoList = data;
                    // console.log(this.videoArr);
                }
            });
        },
    },
    created() {
        const videoData = JSON.parse(sessionStorage.getItem("rowData"));
        this.videoList.unshift(videoData);
        this._videoList();
    },
};
</script>
<style lang="less" scoped>
// 关闭按钮
.to-back {
    width: 60px;
    height: 60px;
    background-color: #fff;
    border-radius: 50%;
    position: absolute;
    top: 5px;
    left: 50px;
    z-index: 1;
    text-align: center;
    opacity: 0.5;

    i {
        color: #000;
        font-size: 30px;
        font-weight: bold;
        line-height: 64px;
    }

    &:hover {
        opacity: 1;
    }
}
</style>