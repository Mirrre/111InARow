<template>
    <div class="home">
        <!-- 头部 -->
        <div class="header-content">
            <Head></Head>
        </div>
        <!-- 导航栏 -->
        <div class="nav">
            <main-nav></main-nav>
        </div>
        <!-- 内容 -->
        <div class="contentBox">
            <!-- 动态路由：内容显示区域 -->
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import Head from "@/components/header/Head.vue"
import MainNav from "@/components/navigation/MainNav.vue"

export default {
    components: {
        Head,
        MainNav
    }
}
</script>

<style lang="less" scoped>
.home {
    color: black;
    .header-content {
        position: fixed;
        // position: absolute;
        right: 0;
        width: calc(100% - 172px);
        height: 59px;
        background-color: #fff;
        border-bottom: 1px solid #fafafa;
        z-index: 999;
    }
    .nav {
        position: fixed;
        width: 172px;
        height: 100%;
        background-color: #fff;
        z-index: 999;
    }
    .contentBox {
        position: absolute;
        top: 60px;
        left: 172px;
        width: calc(100% - 232px);
        height: 100%;
        padding: 24px 30px;
    }
}
</style>