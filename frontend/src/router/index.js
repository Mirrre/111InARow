import Vue from 'vue'
import VueRouter from 'vue-router'
// 导入组件
import Home from '../views/Home.vue'
import Recommend from "../views/index/Recommend.vue"

import Discover from "../views/discover/Discover.vue"
import Hot from "../views/hot/Hot.vue"
import Physical from "../views/physical/Physical.vue"
import Anime from "../views/anime/Anime.vue"
import Game from "../views/game/game.vue"
import Mine from "../views/mine/Mine.vue"

import PlayerVideo from "../views/player/player.vue"
import Search from "../views/search/search.vue"


Vue.use(VueRouter)

const routes = [                          
  {
    path: '/',
    component: Home,
    children: [
      // path 为空，表示默认展示该组件
      {path:'/search',component:Search},
      { path: "Discover", component: Discover }, // 首页
      { path: "", component: Hot }, // 热点
      { path: "Physical", component: Physical }, // 体育
      { path: "Anime", component: Anime }, // 动漫
      { path: "Game", component: Game }, // 游戏
      { path: "Mine", component: Mine }, // 我的
      { path: "Search", component: Search }, // 搜索
    ]
  },
  {
    path: '/player',
    name: "Player",
    component: PlayerVideo
  }
]

const router = new VueRouter({
  routes
})

export default router
