import Vue from 'vue'
import App from './App.vue'
import router from './router'
import http from './http'
import store from './store'

// 引入 Element UI 组件库
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

// Vant
import { Icon } from 'vant';

// 轮播
import VueAwesomeSwiper from 'vue-awesome-swiper'
import 'swiper/css/swiper.css'

Vue.use(ElementUI);

Vue.prototype.$http = http

Vue.config.productionTip = false

Vue.use(Icon);

Vue.use(VueAwesomeSwiper)

// 创建事件总线
export const bus = new Vue();

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
