<template>
    <div class="hd_search">
        <input type="text" placeholder="搜索你感兴趣的内容" />
        <button class="search_btn">
            <span>搜索</span>
        </button>
    </div>
</template>

<script>
export default {};
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
        font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
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