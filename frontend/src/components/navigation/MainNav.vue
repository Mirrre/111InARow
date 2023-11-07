<template>
    <div class="MainNav-container">
        <div class="logo">
            <img src="../../assets/img-slogan-blue-en.png" alt="" />
        </div>
        <ul>
            <li v-for="(item, index) in navArr" :key="item.index">
                <router-link :to="item.url">
                    <img :src="item.filePath" alt="" class="icon" :class="{ active: index == current }" />
                    <span :class="{ active: index == current }">{{ item.title }}</span>
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            current: "",
            navArr: [
                {
                    title: "首页",
                    url: "/discover",
                    filePath: require("../../assets/首页.png"),
                },
                {
                    title: "热点",
                    url: "/",
                    filePath: require("../../assets/校园热点.png"),
                },
                {
                    title: "体育",
                    url: '/physical',
                    filePath:require("../../assets/体育-01.png")
                },
                {
                    title: "动漫",
                    url: '/anime',
                    filePath:require("../../assets/golang.png")
                },
                {
                    title: "游戏",
                    url: '/game',
                    filePath:require("../../assets/游戏机.png")
                },
                {
                    title: "我的",
                    url: '/mine',
                    filePath:require("../../assets/我的.png")
                },
            ],
        };
    },
    watch: {
        // 监听路由：用于控制导航栏的选中状态
        $route: {
            handler(newVal) {
                switch (newVal.path) {
                    case "/":
                        this.current = 1
                        break
                    case "/discover":
                        this.current = 0
                        break
                    // case "/hot":
                    //     this.current = 2
                    //     break
                    case "/physical":
                        this.current = 2
                        break
                    case "/anime":
                        this.current = 3
                        break
                    case "/game":
                        this.current = 4
                        break
                    case "/mine":
                        this.current = 5
                        break    
                }
            },
            immediate: true // 进入页面执行一次
        }
    }
};
</script>

<style lang="less" scoped>
.logo {
    width: 172px;
    height: 60px;
    text-align: center;

    img {
        width: 70%;
        height: 48px;
    }
}

ul li {
    display: block;
    width: 84px;
    height: 36px;
    margin: 30px 48px 0 40px;
    list-style: none;
    cursor: pointer;

    a {
        display: inline-block;
        font-size: 20px;
        line-height: 36px;
        color: #06a7e1;
    }

    .icon {
        width: 36px;
        height: 36px;
        vertical-align: bottom;
        opacity: 0.5;
    }
}

li:hover a {
    color: #067be1;
}

li:hover .icon {
    opacity: 1;
    opacity: 1;
}

.active {
    color: #067be1;
    opacity: 1 !important;
}
</style>