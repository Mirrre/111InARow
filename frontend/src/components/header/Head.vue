<template>
    <div class="Head-container">
        <!-- 搜索栏 -->
        <hd-search />
        <!-- 右侧登录 -->
        <div class="hd_l">
            <hd-login />
        </div>
    </div>
</template>

<script>
import hdSearch from "@/components/header/search/HdSearch.vue"
import hdLogin from "@/components/header/login/HdLogin.vue"

export default {
    components: {
        hdSearch,
        hdLogin
    }
}
</script>

<style lang="less" scoped>
.Head-container {
    position: relative;
    .hd_l {
        width: 514px;
        height: 40px;
        position: absolute;
        top: 0;
        right: 30px;
    }
}
</style>