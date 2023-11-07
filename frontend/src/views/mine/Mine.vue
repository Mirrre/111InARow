<template>
    <div class="mine">
      <div class="mine_box">
        <div class="mine_detail">
          <div class="some_body">
            <!-- 登录成功后显示头像 -->
            <img
              v-if="isLoggedIN"
              :src="avatarImg"
              alt="Avatar"
              class="some_body"
            />
            <img v-else src="@/assets/下载.svg" alt="" />
          </div>
  
          <div class="some_things">
            <h2
              slot="reference"
              v-if="!isLoggedIN"
              class="not_log"
              @click="isLog"
            >
              未登录
            </h2>
            <h2 v-else slot="reference" class="logged_in">已登录</h2>
  
            <el-button type="text" @click="dialogVisible = true" v-if="isLoggedIN"
              >退出登录</el-button
            >
            <el-dialog
              :visible.sync="dialogVisible"
              width="20%"
              :append-to-body="true"
            >
              <span>您确定要退出登录吗</span>
              <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="isLoginOut">确 定</el-button>
              </span>
            </el-dialog>
  
            <div class="some_things2">
              <div class="some_things3">
                <p class="log_nav">关注</p>
              </div>
              <div class="some_things3">
                <p class="log_nav">粉丝</p>
              </div>
              <div class="some_things3">
                <p class="log_nav">获赞</p>
              </div>
            </div>
          </div>
        </div>
  
        <!-- 导航栏和其他内容 -->
        <div class="detail">
          <div class="mine_nav">
            <div class="nav_" @click="selectNavItem('作品')">
              <span class="nav_span">作品</span>
            </div>
            <div class="nav_" @click="selectNavItem('喜欢')">
              <span class="nav_span">喜欢</span>
            </div>
            <div class="nav_" @click="selectNavItem('收藏')">
              <span class="nav_span">收藏</span>
            </div>
          </div>
  
          <!-- 根据选中的导航项来显示内容 -->
          <div class="mine_blank" v-if="selectedNavItem === '作品'">11</div>
          <div class="mine_blank" v-if="selectedNavItem === '喜欢'">
            <template>
              <div class="discover">
                <div class="video_content">
                  <el-row :gutter="16">
                    <el-col
                      :span="4"
                      v-for="(item, index) in videoArr"
                      :key="index"
                    >
                      <div class="grid-content bg-purple">
                        <div
                          class="img_box"
                          @mouseover="
                            videoPlayback(index, videoArr[index].video_url)
                          "
                          @mouseout="
                            videoStopped(index, videoArr[index].video_url)
                          "
                          @click="toPlayer(item)"
                        >
                          <img :src="videoArr[index].thumb" alt="" ref="img" />
                          <!-- contextmenu(禁止右键) muted(解决谷歌浏览器不能自动播放)-->
                          <div class="video_box">
                            <video
                              id="video"
                              ref="video"
                              width="100%"
                              height="100%"
                              @contextmenu.prevent
                              controls
                              controlslist="nodownload nofullscreen noremote footbar"
                              disablePictureInPicture
                              muted="muted"
                              style="display: none; object-fit: fill"
                            >
                              <!-- :poster="videoArr[index].thumb"> -->
                              <source
                                type="audio/mp4"
                                :src="videoArr[index].video_url"
                              />
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
                          <img :src="videoArr[index].avatar" alt="" />
                          <span class="uname">{{
                            videoArr[index].username
                          }}</span>
                          <span class="mtime">{{ videoArr[index].mtime }}</span>
                        </div>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </div>
            </template>
          </div>
          <div class="mine_blank" v-if="selectedNavItem === '收藏'">33</div>
          <div class="mine_blank" v-if="selectedNavItem === '观看历史'">
            <!-- 观看历史内容... -->44
          </div>
        </div>
        <div class="mine_blank" v-if="!isLoggedIn">
          <!-- <p>登录后即可观看喜欢、收藏的视频</p> -->
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { bus } from "@/main.js"; // 导入事件总线
  export default {
    data() {
      return {
        isLoggedIN: false, // 登录状态
        avatarImg: "", // 默认头像链接
        dialogVisible: false, // 控制弹出框的显示
        selectedNavItem: "作品",
        videoArr: [],
        videoTime: "",
        timer: "",
      };
    },
    mounted() {
      this.isLoggedIn(),
      this.videoList(),
        setInterval(() => {
          this.isLoggedIn()
        }, 1000);
      // this.avatar()
    },
    methods: {
      videoList() {
        let data1 = {};
        const userMsg = JSON.parse(localStorage.getItem("userMsg"));
        const isLogin = localStorage.getItem("isLogin");
        // console.log(userMsg.id);
        if (isLogin) {
          data1 = {
            user_id: userMsg.id,
            has: 1,
          };
          this.$http.like.getLikeVideoByUser(data1).then(({ data }) => {
            console.log(data);
            this.videoArr = data;
            // this.videoArr = data.list;
            console.log(this.videoArr);
          });
        }
      },
      videoPlayback(index, video_url) {
        const video = this.$refs.video[index];
        if (video_url) {
          // 显示播放器
          video.style.display = "block";
          // 隐藏图片
          this.$refs.img[index].style.display = "none";
          // 定时器
          this.timer = setTimeout(() => {
            // 播放视频
            video.play();
            // 给视频标签添加缓存播放---video标签属性。
            video.setAttribute("autoplay", "autoplay");
            // 给视频标签添加循环播放---video标签属性。
            video.setAttribute("loop", "loop");
            // 初始播放音量
            video.volume = 0;
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
        const video = this.$refs.video[index];
        if (video_url) {
          // 清除定时器
          clearTimeout(this.timer);
          //停止播放
          video.pause();
          // 重置播放时间
          video.currentTime = 0;
          // 隐藏播放器
          video.style.display = "none";
          // 隐藏图片
          this.$refs.img[index].style.display = "block";
        }
      },
      toPlayer(obj) {
        this.$router.push({
          name: "Player",
          query: { id: obj.id },
        });
        sessionStorage.setItem("rowData", JSON.stringify(obj));
      },
      isLoggedIn() {
        let flag = localStorage.getItem("isLogin");
        if (flag) {
          this.avatarImg = localStorage.getItem("avatar");
          // this.videoList();
      //    let timer = setInterval(() => {
      //     this.videoList()
      //   }, 5000);
          this.isLoggedIN = true;
        } else {
          this.isLoggedIN = false;
        }
      },
      isLoginOut() {
        this.dialogVisible = false;
        localStorage.removeItem("isLogin");
        localStorage.removeItem("avatar");
        localStorage.removeItem("userMsg");
  
        this.isLoggedIN = false;
        this.avatarImg = "@/assets/下载.svg";
  
        // 刷新页面
        location.reload();
      },
      isLog() {
        bus.$emit("openLoginForm");
      },
      selectNavItem(item) {
        this.selectedNavItem = item;
      },
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


  .mine {
    width: 100%;
    height: 100%;
    margin-left: 150px;
    // background-color: pink;
  
    .mine_detail {
      height: 160px;
      width: 100%;
      // background-color: aquamarine;
      align-items: center;
      display: flex;
      justify-content: flex-start;
      margin-bottom: 24px;
      margin-top: 32px;
      position: relative;
      width: 100%;
      z-index: 1;
      color: #333534;
  
      .some_body {
        width: 112px;
        height: 112px;
        // background-image: url('~@/assets/下载.svg');
  
        .some_body1 {
          width: 112px;
          height: 112px;
          border-radius: 50px;
        }
      }
  
      .some_things {
        align-items: flex-start;
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin-left: 40px;
        cursor: pointer;
  
        .some_things2 {
          display: flex;
          justify-content: flex-start;
  
          .some_things3 {
            align-items: baseline;
            display: flex;
            justify-content: flex-start;
  
            .log_nav {
              color: var(--color-text-t3);
              font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
              font-weight: 400;
              padding-right: 10px;
            }
  
            .log_ {
              color: var(--color-text-t1);
              font-size: 20px;
              font-weight: 400;
              line-height: 28px;
              margin-left: 6px;
              transform: scaleX(0.5);
            }
          }
        }
      }
    }
  
    .mine_nav {
      height: 100px;
      width: 100%;
      // background-color: burlywood;
      display: flex;
      height: 26px;
      color: #333534;
      font-size: 16px;
      border-bottom: 1px solid #f7f7f7;
  
      .nav_ {
        align-items: center;
        display: flex;
        margin-right: 40px;
  
        .nav_span {
          margin-right: 6px;
        }
      }
  
      .nav_:hover {
        border-bottom: 1px solid #333534;
      }
    }
  
    .mine_blank {
      width: 83%;
      height: 600px;
      // background-color: aquamarine;
      // margin: 150px 300px;
      font-size: 20px;
      color: #808080;
    }
  }
  </style>
  