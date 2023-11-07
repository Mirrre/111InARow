<template>
  <div
    class="swiper-video-container"
    :style="`width: ${innerWidth}; height: ${innerHeight}`"
  >
    <!-- 轮播区域 -->
    <swiper
      class="my-swipe"
      ref="playerSwiper"
      :options="swiperOptions"
      v-if="videoList.length"
      @slideChange="_slideChange"
      navigation
    >
      <swiper-slide v-for="(item, index) in videoList" :key="item.id">
        <!-- 右侧信息区域 -->
        <div class="right_menus">
          <!-- 作者 -->
          <div class="menuClick">
            <el-tooltip
              class="item"
              effect="dark"
              content="进入作者主页"
              placement="right-start"
            >
              <img class="avatar" :src="item.avatar" alt="" />
            </el-tooltip>
            <el-tooltip
              class="item"
              effect="dark"
              content="关注"
              placement="right-start"
            >
              <div class="follow">
                <i class="van-icon van-icon-plus"></i>
              </div>
            </el-tooltip>
          </div>
          <!-- 点赞 -->
          <div class="click-info" @click="like(item)" v-if="item.is_follow=='1'">
            <el-tooltip
              class="item"
              effect="dark"
              content="点赞"
              placement="right-start"
            >
              <i class="van-icon van-icon-like" style="color: red"></i>
            </el-tooltip>
            <div class="text">{{ item.like_num }}</div>
          </div>
          <!-- 点赞 -->
          <div class="click-info" @click="like(item)" v-else>
            <el-tooltip
              class="item"
              effect="dark"
              content="点赞"
              placement="right-start"
            >
              <i class="van-icon van-icon-like"></i>
            </el-tooltip>
            <div class="text">{{ item.like_num }}</div>
          </div>
          <!-- 评论 -->
          <div class="click-info">
            <el-tooltip
              class="item"
              effect="dark"
              content="评论"
              placement="right-start"
            >
              <i class="van-icon van-icon-chat"></i>
            </el-tooltip>
            <div class="text">{{ item.comment_num }}</div>
          </div>
          <!-- 收藏 -->
          <div class="click-info" @click="collect(item)" v-if="item.is_follow2=='1'">
            <el-tooltip
              class="item"
              effect="dark"
              content="收藏"
              placement="right-start"
            >
              <i class="van-icon van-icon-star" style="color: yellow"></i>
            </el-tooltip>
            <div class="text">{{ item.collect_num }}</div>
          </div>
          <div class="click-info" @click="collect(item)" v-else>
            <el-tooltip
              class="item"
              effect="dark"
              content="收藏"
              placement="right-start"
            >
              <i class="van-icon van-icon-star"></i>
            </el-tooltip>
            <div class="text">{{ item.collect_num }}</div>
          </div>
          <!-- 分享 -->
          <div class="click-info">
            <el-tooltip
              class="item"
              effect="dark"
              content="分享"
              placement="right-start"
            >
              <i class="van-icon van-icon-share"></i>
            </el-tooltip>
            <!-- <div class="text">{{ item.share_num }}</div> -->
          </div>
          <!-- 切换按钮 -->
          <div class="toggle-swiper">
            <div class="swiper-button-prev" slot="button-prev">
              <i class="van-icon van-icon-arrow-up"></i>
            </div>
            <div class="swiper-button-next" slot="button-next">
              <i class="van-icon van-icon-arrow-down"></i>
            </div>
          </div>
        </div>
        <!-- 底部文本区域 -->
        <div class="text-container">
          <div class="nickname">
            @{{ item.nickname ? item.nickname : item.username }}
          </div>
          <div class="msg">{{ item.title }}</div>
        </div>
        <!-- 播放器 -->
        <video
          class="video-content"
          :id="`player${index}`"
          :src="item.video_url"
          preload="auto"
          controls
          loop
          width="100%"
          height="100%"
          type="video/mp4"
          playsinline="true"
          x5-playsinline="true"
          webkit-playsinline="true"
          x-webkit-airplay="allow"
          oncontextmenu="return false;"
          controlslist="nodownload noremoteplayback"
          :disablePictureInPicture="true"
        ></video>
      </swiper-slide>
    </swiper>
  </div>
</template>

<script>
import { Like, isLike } from "@/api/like.js";
// 导入一个节流函数库，例如 lodash.throttle
import throttle from 'lodash/throttle';
import debounce from 'lodash/debounce';
export default {
  props: ["videoList", "innerWidth", "innerHeight"],
  data() {
    return {
      swiperOptions: {
        direction: "vertical",
        mousewheel: {
          invert: true,
        },
        mousewheel: true,
        navigation: {
          nextEl: ".swiper-button-next",
          prevEl: ".swiper-button-prev",
        },
      },
    };
  },
  computed: {
    swiper() {
      return this.$refs.playerSwiper.$swiper;
    },
  },
  mounted() {
    // 使用 debounce 包装鼠标滚轮事件处理函数，设置等待时间为0，以确保立即响应事件
    this.handleMouseWheelDebounced = debounce(this.handleMouseWheel, 0);
    // 添加键盘时间监听器
    document.addEventListener('keydown', this.handleKeyDown)

    // 添加事件监听
    window.addEventListener('mousewheel', this.handleMouseWheelThrottled);
  },
  methods: {
      collect(item) {
      const userMsg = JSON.parse(localStorage.getItem('userMsg'))  
      const data1 = {
        user_id: userMsg.id,
        video_id: item.id,
      };
      let result = [];
      this.$http.collect.isCollect(data1).then(({ code, data }) => {
        if (code === 200) {
          console.log(data);
          result = data;
          if (result.length != 0) {
          console.log(1);
            this.$http.collect.deleCollect(data1).then(({ code, data }) => {
            if (code === 200) {
            // this.videoList = this.videoList.concat(data);
            console.log(data);
            }
        });
      } else {
          console.log(2);
                console.log(data1);
                this.$http.collect.Collect(data1).then(({ code, data }) => {
                if (code === 200) {
                // this.videoList = this.videoList.concat(data);
                console.log(data);
                }
            });
      }
        }
      });
    },
    handleKeyDown(event) {
      switch (event.key) {
        case 'ArrowUp':
          this.slideToPrevious()
          break
        case 'ArrowDown':
          this.slideToNext()
          break
      }
    },
    slideToPrevious() {
      const currentIndex = this.swiper.activeIndex
      if (currentIndex > 0) {
        this.swiper.slideTo(currentIndex - 1)
      }
    },
    // 切换到下一个视频
    slideToNext() {
      const currentIndex = this.swiper.activeIndex;
      if (currentIndex < this.videoList.length - 1) {
        this.swiper.slideTo(currentIndex + 1);
      }
    },
      // 鼠标滚轮事件处理函数
  handleMouseWheel(event) {
    // 在处理事件时设置标志位，以避免重复触发
    if (this.handlingMouseWheel) {
      return;
    }

    this.handlingMouseWheel = true;

    // 处理鼠标滚轮事件的代码

    // 处理完成后重置标志位
    this.handlingMouseWheel = false;
  },
    like(item) {
      const userMsg = JSON.parse(localStorage.getItem('userMsg'))  
      const data1 = {
        user_id: userMsg.id,
        video_id: item.id,
      };
      let result = [];
      this.$http.like.isLike(data1).then(({ code, data }) => {
        if (code === 200) {
          console.log(data);
          result = data;
          if (result.length != 0) {
          console.log(1);
            this.$http.like.deleLike(data1).then(({ code, data }) => {
            if (code === 200) {
            // this.videoList = this.videoList.concat(data);
            console.log(data);
            }
        });
      } else {
          console.log(2);
                console.log(data1);
                this.$http.like.Like(data1).then(({ code, data }) => {
                if (code === 200) {
                // this.videoList = this.videoList.concat(data);
                console.log(data);
                }
            });
      }
        }
      });
    },
    // swiper 切换
    _slideChange() {
      const index = this.swiper.activeIndex;
      var player = document.getElementById(`player${index}`);
      if (player.paused) {
        this.videoList.forEach((item, s_index) => {
          if (s_index == index) {
            document.getElementById(`player${s_index}`).play();
            item.play = true;
          } else {
            document.getElementById(`player${s_index}`).pause();
            item.play = false;
          }
        });
      } else {
        this.videoList.forEach((item, s_index) => {
          document.getElementById(`player${s_index}`).pause();
          item.play = false;
        });
      }
    },
  },
};
</script>

<style lang="less" scoped>
.swiper-video-container {
  position: fixed;
  // width: 88%;
  // height: calc(100% - 130px);
  border-radius: 6px;
  overflow: hidden;
  cursor: pointer;

  .my-swipe {
    position: relative;
    width: 90%;
    // 解决滑动轮播时有白色线条的Bug
    height: 102%;
    margin-right: 95px;

    .swiper-slide {
      position: relative;
      height: 100%;

      // right: 40px;
      // padding-right: 20px;
      .video-content {
        width: 100%;
        height: 100%;
        background-color: #000;
        // object-fit: cover;
        transform: translateZ(0);

      }

      .right_menus {
        position: absolute;
        width: 50px;
        bottom: 80px;
        right: 20px;
        z-index: 1;
        text-align: center;

        .menuClick {
          position: relative;
          width: 100%;
          margin-bottom: 30px;

          .avatar {
            width: 42px;
            height: 42px;
            border-radius: 50%;
          }

          .follow {
            position: absolute;
            width: 20px;
            height: 20px;
            left: 50%;
            bottom: -5px;
            transform: translateX(-50%);
            font-size: 12px;
            color: #fff;
            background-color: #fe2c55;
            border-radius: 50%;
            line-height: 20px;
          }
        }

        .click-info {
          position: relative;
          margin-bottom: 20px;

          i {
            font-size: 24px;
            color: #fff;
          }

          .text {
            color: #fff;
            font-size: 16px;
            margin-top: 5px;
          }
        }
      }

      .text-container {
        position: absolute;
        width: 30%;
        left: 0;
        bottom: 80px;
        padding: 0 30px;
        box-sizing: border-box;
        z-index: 1;

        .nickname {
          width: 100%;
          font-size: 24px;
          color: #fff;
          text-align: left;
        }

        .msg {
          position: relative;
          width: 100%;
          box-sizing: border-box;
          font-size: 18px;
          text-align: left;
          color: #fff;
        }
      }
    }
  }

  // 切换按钮
  .toggle-swiper {
    position: absolute;
    top: -120px;
    left: 50%;
    transform: translateX(-50%);
    width: 36px;
    height: 80px;
    background-color: #323442;
    border-radius: 36px;
    display: flex;
    flex-direction: column;
    opacity: 0.7;

    &:hover {
      opacity: 1;
    }
  }

  .swiper-button-prev,
  .swiper-button-next {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 36px;
    height: 40px;
    position: static;
    color: #fff;
    margin-top: 0;

    &::after {
      content: "";
    }

    &:hover {
      color: #fe2c55;
    }

    i {
      font-weight: bold;
    }
  }
}
</style>

