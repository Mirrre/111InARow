<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>
export default {

}
</script>

<style lang="less">
* {
    padding: 0;
    margin: 0;
}
ul,li {
    padding:0;
    margin:0;
    list-style:none
}
a {
    text-decoration: none;
    color: #fff;
}
body {
    font-family: "PingFang SC", "Lantinghei SC", "Microsoft YaHei","Helvetica Neue", "Open Sans", Arial, "Hiragino Sans GB", sans-serif;
    background-color: #fafafa;
}
button {
    margin: 0;
    padding: 0;
    border: 0;
    outline: none;
}
input {
    outline:none;
    border: 0;
    background:none;
}
// 隐藏滚动条
::-webkit-scrollbar {
    width: 0 !important;
    height: 0 !important;
}
</style>