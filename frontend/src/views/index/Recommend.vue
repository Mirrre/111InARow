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
    methods: {
        _videoList() {
            this.$http.video.hotList().then(({ code, data }) => {
                if (code === 200) {
                    this.videoList = this.videoList.concat(data.list);
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