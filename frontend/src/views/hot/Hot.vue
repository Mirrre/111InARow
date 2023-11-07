<template>
    <div class="short-video-player">
        <swiper-video 
            :videoList="videoList"
            :innerWidth="w"
            :innerHeight="h">
        </swiper-video>
    </div>
</template>

<script>
import SwiperVideo from "@/components/swiper/index.vue"

export default {
    components: {
        SwiperVideo
    },
    data() {
        return {
            videoList: [],
            w: "88%",
            h: "calc(100% - 130px)"
        };
    },
    mounted() {
        this._videoList();
        setInterval(()=>{
            this._videoList();
        },400)
    },
    methods: {
        _videoList() {
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
            this.$http.video.hotList(data1).then(({ code, data }) => {
                if (code === 200) {
                    // console.log(data);
                    this.videoList = data;
                    // console.log(this.videoList);
                }
            });
        },
    },
    created() {
        this._videoList();
    },
};
</script>

<style lang="less" scoped>

</style>>