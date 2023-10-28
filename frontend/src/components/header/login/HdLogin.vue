<template>
    <div class="hd_login">
        <div class="hd_left">
            <!-- 发布与登录 -->
            <div class="box_r">
                <button class="up_video">
                    <i class="el-icon-plus"></i>
                    <span> 发布视频</span>
                </button>
                <button class="btn_login" @click="login">登录</button>
                <div class="overlay" v-if="showLoginPopup">
                    <div class="pop_login">
                        <!-- 登录表单 -->
                        <form>
                            <!-- 关闭按钮 -->
                            <button class="close-button" @click="closeLogin">关闭</button>
                            <input type="text" id="username" name="username" placeholder="请输入用户名"><br><br>
                            <input type="password" id="password" name="password" placeholder="请输入密码"><br><br>
                            <button @click="performAuthentication(isLogin)" class="isLoginBtn">{{ !isLogin?"登录":"注册" }}</button>
                            <span class="registerBtn" @click="isLogin = !isLogin">{{ isLogin ? '登录' : '注册' }}</span>
                            <!-- 显示头像 -->
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import {Login,Register} from '@/api/login.js'
export default {
    data() {
        return {
            showLoginPopup: false, // 初始时不显示登录弹窗
            isisLogin: false,
            loginInfo:{
                username:"",
                password:""
            }
        };
    },
    methods: {
        login() {
            this.showLoginPopup = true;
        },
        closeLogin() {
            this.showLoginPopup = false;
        },
        isLogin() {
            this.showLoginPopup = false;
            this.isisLogin = true; // 用户已登录
        },
        performAuthentication(isLogin) {
            if (isLogin) {
                Login(this.loginInfo)
                    .then(response => {
                        // 处理登录成功的情况
                        console.log('登录成功', response);
                    })
                    .catch(error => {
                        // 处理登录失败的情况
                        console.error('登录失败', error);

                    });
            } else {
                // 执行注册逻辑
                Register(this.loginInfo)
                    .then(response => {
                        // 处理注册成功的情况
                        console.log('注册成功', response);
                    })
                    .catch(error => {
                        // 处理注册失败的情况
                        console.error('注册失败', error);
                    });
            }
        }
    },
};
</script>

<style lang="less" scoped>
// Element 样式
// 
.hd_left {
    position: relative;

    .box_r {
        position: relative;
        width: 228px;
        height: 40px;
        position: absolute;
        top: 0;
        right: 0;

        .btn_login {
            width: 38px;
            height: 38px;
            line-height: 38px;
            background-color: #06a7e1;
            border-radius: 20px;
            color: #fff;
            font-size: 15px;
            position: absolute;
            top: 50%;
            right: 30px;
            transform: translate(0, -50%);
        }

        .btn_login:hover {
            background-color: #067be1;
        }

        .up_video {
            width: 108px;
            height: 40px;
            background-color: #424242;
            border-radius: 6px;
            border: 2px solid #fff;
            color: #fff;
            font-size: 15px;
        }

        .overlay {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 0, 0, .8);
            display: flex;
            justify-content: center;
            align-items: center;
            // z-index: 1;
        }

        .user-avatar {
            color: #000;
            width: 38px;
            height: 38px;
            line-height: 38px;
            // background-color: #06a7e1;
            border-radius: 20px;
            position: absolute;
            top: 50%;
            right: 30px;
            transform: translate(0, -50%);
            background-color: #fff;
            background-image: url('~@/assets/19.png');
        }

        .pop_login {
            position: absolute;
            z-index: 1;
            background: #fff;
            /* 白色背景 */
            padding: 80px;
            border: 1px solid #000;
            /* 可根据需要添加样式 */
            z-index: 2;
            color: #424242;
            /* 使登录表单显示在遮罩层上面 */
            background-image: url('~@/assets/sp.png');
            background-color: #ffffff;
        }

        .close-button {
            position: absolute;
            top: 10px;
            right: 10px;
            background: none;
            border: none;
            cursor: pointer;
            color: #000;
        }

        /* 关闭按钮的颜色 */
        input {
            background-color: #f2f2f4;
            color: #424242;
            border: none;
            padding: 10px;
            margin: 5px;
            border-radius: 5px;
        }

        .isLoginBtn {
            width: 104px;
            height: 38px;
            line-height: 38px;
            background-color: #06a7e1;
            cursor: pointer;
            margin: 20px 40px;
            border-radius: 10px;
        }

        .registerBtn {
            display: block;
            color: #06a7e1;
            font-size: 30rpx;
            bottom: 0px;
            text-align: end;
            cursor: pointer;
        }
    }
}

// 公共样式
.btn_add,
.btn_login,
.up_video {
    font-family: "PingFang SC", "Lantinghei SC", "Microsoft YaHei",
        "Helvetica Neue", "Open Sans", Arial, "Hiragino Sans GB", sans-serif;
    cursor: pointer;
}
</style>