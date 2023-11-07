import Vue from 'vue'
import App from './App.vue'
import router from './router'
import http from './http'

// 引入 Element UI 组件库
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

// Vant
import { Icon } from 'vant';

// 轮播
import VueAwesomeSwiper from 'vue-awesome-swiper'
import 'swiper/css/swiper.css'
// 创建事件总线
export const bus = new Vue();

Vue.use(ElementUI);

Vue.prototype.$http = http

Vue.config.productionTip = false

Vue.use(Icon);

Vue.use(VueAwesomeSwiper)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
