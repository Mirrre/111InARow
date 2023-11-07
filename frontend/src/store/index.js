// store.js

import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    isLoggedIN: false, // 登录状态
    avatarImg: '', // 默认头像链接
  },
  mutations: {
    // 定义 mutation 来更新登录状态和头像
    setIsLoggedIn(state, isLoggedIn) {
      state.isLoggedIN = isLoggedIn;
    },
    setAvatarImg(state, avatarImg) {
      state.avatarImg = avatarImg;
    },
  },
  actions: {
    // 定义 action 用于触发 mutation 更新登录状态和头像
    userLoggedIn({ commit }, { avatarUrl }) {
      commit('setIsLoggedIn', true);
      commit('setAvatarImg', avatarUrl);
    },
    userLoggedOut({ commit }) {
      commit('setIsLoggedIn', false);
      commit('setAvatarImg', '');
    },
  },
});

export default store;
