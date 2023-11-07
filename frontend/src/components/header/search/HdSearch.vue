<template>
  <div class="hd_search">
    <input type="text" placeholder="搜索你感兴趣的内容" v-model="title" />
    <button class="search_btn" @click="searchVideo()">
      <router-link to="/search">搜索</router-link>
    </button>
    <!-- <searchVue :videoArr="videoArr"> -->
    <!-- </searchVue> -->
  </div>
</template>

<script>
// import { component } from "vue/types/umd";
// import searchVue from "../../../views/search/search.vue";

export default {
  // components: {
  //   searchVue,
  // },
  data() {
    return {
      title: "",
      videoArr: []
    };
  },
  methods: {
    searchVideo() {
      let data1 = {};
      const userMsg = JSON.parse(localStorage.getItem("userMsg"));
      console.log(userMsg);
      //   const isLogin = localStorage.getItem("isLogin");
      //   console.log(isLogin);
      //   if (isLogin == true) {
      data1 = {
        title: this.title,
        user_id: userMsg.id,
      };
      //   } else if(isLogin == false){
      //     data1 = {
      //       title: this.title,
      //     };
      //   }
      this.$http.video.findAllVideoByTitle(data1).then(({ code, data }) => {
        if (code === 200) {
          console.log(data);
          this.videoArr = data;
          localStorage.setItem('videoArr',JSON.stringify(this.videoArr))
          // console.log(this.videoList);
        }
      });
    },
  },
};
</script>

<style lang="less" scoped>
.hd_search {
  display: flex;
  margin: 10px 30px;
  width: 35.49%;
  height: 40px;
  overflow: hidden;
  border-radius: 4px;
  background-color: #f6f6f6;
  input {
    // width: calc(100% - 93px);
    // height: 100%;
    flex-grow: 1;
    font-size: 14px;
    font-weight: 400;
    line-height: 22px;
    font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
    color: black;
    padding: 0 12px;
    transition: all 0.2s;
    // 光标颜色
    caret-color: red;
  }
  input::-webkit-input-placeholder {
    color: #ccc;
  }
  .search_btn {
    width: 69px;
    height: 100%;
    background-color: #06a7e1;
    color: #fff;
    font-size: 16px;
    font-family: "PingFang SC", "Lantinghei SC", "Microsoft YaHei",
      "Helvetica Neue", "Open Sans", Arial, "Hiragino Sans GB", sans-serif;
    transition: all 0.2s;
    cursor: pointer;
  }
  .search_btn:hover {
    background-color: #067be1;
  }
}
.hd_search:hover input {
  background-color: #fff;
}
</style>
