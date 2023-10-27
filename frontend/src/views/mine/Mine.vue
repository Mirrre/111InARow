<template>
    <div class="mine">
        <div class="mine_box">
            <div class="mine_detail">
                <div class="some_body">
                    <div>

                    </div>
                </div>
                <div class="some_things">
                    <h2 class="not_log">未登录</h2>
                    <div class="some_things2">
                        <div class="some_things3">
                            <p class="log_nav">关注</p>
                        </div>
                        <div class="some_things3">
                            <p class="log_nav">粉丝</p>
                        </div>
                        <div class="some_things3">
                            <p class="log_nav">获赞</p>
                        </div>
                    </div>

                </div>
            </div>


            <div class="mine_nav">
                <div class="nav_">
                    <span class="nav_span">作品</span>
                </div>
                <div class="nav_">
                    <span class="nav_span">喜欢</span>
                </div>
                <div class="nav_">
                    <span class="nav_span">收藏</span>
                </div>
                <div class="nav_">
                    <span class="nav_span">观看历史</span>
                </div>
            </div>
            <div class="mine_blank">
                <p>登录后即可观看喜欢、收藏的视频</p>
            </div>
        </div>

    </div>
</template>

<script>
export default {

}
</script>

<style lang="less">
.mine {
    width: 100%;
    height: 100%;
    margin-left: 150px;
    // background-color: pink;

    .mine_detail {
        height: 160px;
        width: 100%;
        // background-color: aquamarine;
        align-items: center;
        display: flex;
        justify-content: flex-start;
        margin-bottom: 24px;
        margin-top: 32px;
        position: relative;
        width: 100%;
        z-index: 1;
        color: #333534;

        .some_body {
            width: 112px;
            height: 112px;
            background-image: url('~@/assets/下载.svg');
        }

        .some_things {
            align-items: flex-start;
            display: flex;
            flex-direction: column;
            justify-content: center;
            margin-left: 40px;
            cursor: pointer;

            .some_things2 {
                display: flex;
                justify-content: flex-start;

                .some_things3 {
                    align-items: baseline;
                    display: flex;
                    justify-content: flex-start;

                    .log_nav {
                        color: var(--color-text-t3);
                        font-family: PingFang SC, DFPKingGothicGB-Regular, sans-serif;
                        font-weight: 400;
                        padding-right: 10px;
                    }

                    .log_ {
                        color: var(--color-text-t1);
                        font-size: 20px;
                        font-weight: 400;
                        line-height: 28px;
                        margin-left: 6px;
                        transform: scaleX(.5);
                    }
                }
            }

        }
    }

    .mine_nav {
        height: 100px;
        width: 100%;
        // background-color: burlywood;
        display: flex;
        height: 26px;
        color: #333534;
        font-size: 16px;
        border-bottom: 1px solid #f7f7f7;
        .nav_ {

            align-items: center;
            display: flex;
            margin-right: 40px;

            .nav_span {
                margin-right: 6px;
            }
        }
    }

    .mine_blank{
        width: 100%;
        height: 20px;
        margin: 200px 500px;
        font-size: 20px;
        color: #808080;
    }
}
</style>