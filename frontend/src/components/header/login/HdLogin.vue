<template>
    <div class="hd_login">
        <div class="hd_left">
            <!-- 发布与登录 -->
            <div class="box_r" style="display: flex; justify-content: space-between;">
                <!-- <button class="upload_video" @click="uploadVideo">
                    <i class="el-icon-plus"></i>
                    <span> 上传视频</span>
                </button> -->
                <button class="up_video" @click="openVideoModal">
                    <i class="el-icon-plus"></i>
                    <span> 发布视频</span>
                </button>
                <!-- 上传视频 -->
                <div class="uploadVideo">
                    <el-upload class="upload-video"
                        :action="`http://localhost:8081/api/publish/action?user_id=${user_id}&title=${title}`"
                        :on-success="handleUploadSuccess" :show-file-list="false" name="data">
                        <button class="upload_video" @click="uploadVideo">
                            <i class="el-icon-plus"></i>
                            <span> 上传视频</span>
                        </button>
                    </el-upload>

                </div>
                <button class="btn_login" @click="login">登录</button>


                <!-- <img v-if="isLogin && avatar" class="btn_login" :src="avatar" alt="Avatar"> -->
                <el-popover trigger="hover" placement="bottom" width="100" style="z-index: 999;">
                    <img v-if="isLogin && avatar" slot="reference" class="btn_login" :src="avatar" alt="Avatar">
                    <el-button type="text" @click="dialogVisible = true">退出登录</el-button>
                    <el-dialog :visible.sync="dialogVisible" width="20%" :append-to-body="true">
                        <span>您确定要退出登录吗</span>
                        <span slot="footer" class="dialog-footer">
                            <el-button @click="dialogVisible = false">取 消</el-button>
                            <el-button type="primary" @click="isLogout">确 定</el-button>
                        </span>
                    </el-dialog>
                </el-popover>

                <div class="overlay" v-if="showLoginPopup">
                    <div class="pop_login">
                        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px"
                            class="demo-ruleForm">
                            <button class="close-button" @click="closeLogin">
                                <i class="el-icon-close"></i>
                            </button>
                            <el-form-item label="账号" prop="username">
                                <el-input type="username" v-model="ruleForm.username" placeholder="请输入用户名"
                                    autocomplete="off"></el-input>
                            </el-form-item>
                            <el-form-item label="密码" prop="password">
                                <el-input :type="isShowPassword ? 'password' : 'username'" v-model="ruleForm.password"
                                    placeholder="请输入密码" autocomplete="off">
                                    <template #append>
                                        <!-- 在密码框右侧添加按钮 -->
                                        <i class="el-icon-view" @click="showPassword"></i>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" @click="submitForm('ruleForm')">{{ submitButtonText }}</el-button>
                                <el-button @click="resetForm('ruleForm')">重置</el-button>
                            </el-form-item>
                            <el-tag style="margin-left:280px;margin-top:10px" type="success" plain @click="showLogin">{{
                                subimtButtonNoText }}</el-tag>
                        </el-form>
                    </div>
                </div>

                <!-- 视频弹窗 -->
                <div class="center-video">
                    <el-dialog :modal="false" class="videoBox" :visible="showVideoModal"
                        :title="flag == 2 ? '调用摄像头出错，请重试！' : '视频录制'" v-show="flag != 2" :before-close="closeVideoModal">
                        <div class="video-container">
                            <video id="videoElement" autoplay class="video1"></video>
                        </div>
                        <div class="button-container">
                            <el-button class="button-left" type="warning" round @click="closeVideoModal">取消</el-button>
                            <el-button class="button-right" type="primary" round @click="closeVideoModal">确定</el-button>
                        </div>
                    </el-dialog>
                </div>
            </div>
        </div>
    </div>
</template>

<script >
import { bus } from "@/main.js"; // 导入事件总线
import { Login, Register } from '@/api/login.js'
import layui from 'vue-lay'
import { upLoadVideo } from '@/api/video.js'

export default {
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入账号'));
            } else if (value.length < 6 || value.length > 25) {
                callback(new Error('输入的账号长度小于6或大于25'));
            } else {
                if (this.ruleForm.password !== '') {
                    this.$refs.ruleForm.validateField('password');
                }
                callback();
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'));
            } else if (value.length < 6 || value.length > 25) {
                callback(new Error('输入的密码长度小于6或大于25'));
            } else {
                callback();
            }
        };
        return {
            avatar: localStorage.getItem('avatar') || '', // 存储头像信息
            showVideoModal: false,
            showVideoUpload: true,
            flag: 1,
            showSuccess: false,
            showError: false,
            isLogin: false,
            isShowPassword: true,
            showLoginPopup: false, // 初始时不显示登录弹窗
            showLogout: false,
            dialogVisible: false,
            // videoUrl: '',
            user_id: '1',

            title: 'test',
            loginInfo: {
                username: "",
                password: "",
            },
            ruleForm: {
                username: '',
                password: ''
            },
            rules: {
                username: [
                    { validator: validatePass, trigger: 'blur' }
                ],
                password: [
                    { validator: validatePass2, trigger: 'blur' }
                ]
            },
        };
    },
    computed: {
        submitButtonText() {
            return this.isLogin ? '登录' : '注册';
        },
        subimtButtonNoText() {
            return this.isLogin ? '注册' : '登录';
        }
    },
    created() {
        // 从本地存储中获取用户登录状态
        const storedIsLogin = localStorage.getItem('isLogin');
        if (storedIsLogin === 'true') {
            this.isLogin = true;
        }

        // 监听登录成功事件
        bus.$on("user-logged-in", (avatarUrl) => {
            this.isLogin = true;
            this.avatar = avatarUrl;

            // 将用户登录状态存储到本地存储
            localStorage.setItem('isLogin', 'true');
            localStorage.setItem('avatar', avatarUrl);
            console.log(localStorage.getItem('isLogin'));
            console.log(localStorage.getItem('avatar'));
        });

        bus.$on("openLoginForm", this.openLoginForm)
    },
    mounted() {
        this.getUserMsg(),
            setInterval(() => {
                this.getUserMsg()
            }, 3000)
    },
    methods: {
        getUserMsg() {
            const userMsg = JSON.parse(localStorage.getItem('userMsg'))
            // 从组件的 data 中获取 userId 和 title
            this.user_id = userMsg.id;
            //  console.log(this.user_id);
        },
        openLoginForm() {
            // 在这里打开登录表单
            this.showLoginPopup = true;
        },
        open1() {
            const h = this.$createElement;
            this.$notify({
                title: '提交成功!',
                //   message: h('i', { style: 'color: teal'}, '这是提示文案这是提示文案这是提示文案这是提示文案这是提示文案这是提示文案这是提示文案这是提示文案')
            });
        },
        open2() {
            this.$notify({
                title: '错误提交',
                //   message: '这是一条不会自动关闭的消息',
            });
        },
        showLogin() {
            this.isLogin = !this.isLogin
            // console.log(this.isLogin);
        },
        showPassword() {
            this.isShowPassword = !this.isShowPassword
        },
        submitForm(formName) {
            if (this.isLogin) {
                Login(this.ruleForm)
                    .then(response => {
                        // console.log(response);
                        const msg = response.data[0];
                        // 处理登录成功的情况
                        this.open1();
                        this.showLoginPopup = false;
                        // 假设登录成功后获取到头像链接
                        const avatarUrl = 'https://th.bing.com/th/id/OIP.fboUsSme2Jpa8AIazc5nZQAAAA?pid=ImgDet&rs=1';

                        // 将头像链接存储到avatar属性中
                        this.avatar = avatarUrl;
                        bus.$emit("user-logged-in", avatarUrl);
                        // 将ruleForm的数据保存到localStorage中
                        localStorage.setItem('ruleForm', JSON.stringify(this.ruleForm));
                        localStorage.setItem('userMsg', JSON.stringify(msg));

                    })
                    .catch(error => {
                        // 处理登录失败的情况
                        this.open2();
                    });
            } else {
                // 执行注册逻辑
                // console.log(this.ruleForm);
                Register(this.ruleForm)
                    .then((response) => {
                        const msg = response.data[0];
                        // 处理注册成功的情况
                        this.open1();
                        this.showLoginPopup = false;
                        // 假设登录成功后获取到头像链接
                        const avatarUrl = 'https://th.bing.com/th/id/OIP.fboUsSme2Jpa8AIazc5nZQAAAA?pid=ImgDet&rs=1';

                        // 将头像链接存储到avatar属性中
                        this.avatar = avatarUrl;
                        bus.$emit("user-logged-in", avatarUrl);
                        // 将ruleForm的数据保存到localStorage中
                        localStorage.setItem('ruleForm', JSON.stringify(this.ruleForm));
                        localStorage.setItem('userMsg', JSON.stringify(msg));
                    })
                    .catch((error) => {
                        // 处理注册失败的情况
                        this.open2();
                    });
            }
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    //   this.open1();
                    // alert('提交成功!');
                } else {
                    // this.open2();
                    // console.log('error submit!!');
                    return false;
                }
            });
            // this.showLoginPopup = false;
            if (this.flag2 == 1) {
                console.log(222);
            }
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        login() {
            this.showLoginPopup = true;
        },
        closeLogin() {
            this.showLoginPopup = false;
        },
        isLogout() {
            this.dialogVisible = false
            // 处理退出登录成功的情况
            // 清除本地存储的用户信息和登录状态
            localStorage.removeItem('isLogin');
            localStorage.removeItem('avatar');
            localStorage.removeItem('userMsg');
            this.isLogin = false;
            this.avatar = '';

            // 触发用户退出登录的事件
            // bus.$emit("user-logged-out");
        },
        load() {
            const videoElement = document.getElementById("videoElement");
        },
        // 在 handleUploadSuccess 函数中
        handleUploadSuccess(response, file, fileList) {
            if (response.status === 200) {
                this.$notify({
                    title: '上传成功',
                    message: '视频上传成功',
                    type: 'success'
                });
                // this.videoUrl = response.data.videoUrl;

                console.log(2);
                const userMsg = JSON.parse(localStorage.getItem('userMsg'))
                // 从组件的 data 中获取 userId 和 title
                const user_id = userMsg.id;
                const title = this.title;

                // 构建一个对象包含 userId、title 和 videoUrl
                const dataToSend = {
                    user_id,
                    title
                    // videoUrl: this.videoUrl
                };
                console.log(dataToSend);
                // 调用 uploadVideo API 并传递 dataToSend
                upLoadVideo(dataToSend)
                    .then(apiResponse => {
                        // 处理后端 API 的响应
                        console.log(1);
                        console.log('后端 API 响应：', apiResponse.data);
                    })
                    .catch(error => {
                        // 处理请求失败
                        console.log(3);
                        console.error('请求失败：', error);
                    });
            } else {
                console.log(4);
                this.$notify.error('视频上传失败');
            }
        },
        uploadVideo() {
            this.showVideoUpload = true
            // console.log(111);
        },
        openVideoModal() {
            this.showVideoModal = true;
            const videoElement = document.getElementById("videoElement");
            // console.log(videoElement);
            navigator.mediaDevices
                .getUserMedia({ video: true })
                .then(stream => {
                    videoElement.srcObject = stream;
                })
                .catch(error => {
                    console.error("Error accessing the camera: ", error);
                    this.showVideoModal = false;
                });
            this.flag++
        },
        closeVideoModal(done) {
            this.showVideoModal = false;
            const videoElement = document.getElementById("videoElement");
            if (videoElement.srcObject) {
                const tracks = videoElement.srcObject.getTracks();
                tracks.forEach(track => track.stop());
                videoElement.srcObject = null;
            }
            done();
        }
    },
};
</script>

<style lang="less" scoped>
// Element 样式
// 
.hd_left {
    position: relative;
    // margin-right: 20px;

    .box_r {
        position: relative;
        width: 248px;
        height: 40px;
        position: absolute;
        top: 0;
        right: 0;
        margin-right: 20px;

        .uploadVideo{
            margin-right: 25px;
        }
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
            right: 0px;
            transform: translate(0, -50%);
            z-index: 1;
            /* 设置z-index属性 */
        }

        .btn_login:hover {
            background-color: #067be1;
        }

        .upload_video {
            width: 100px;
            height: 40px;
            // padding-right: 150px;
            background-color: #424242;
            border-radius: 6px;
            border: 2px solid #fff;
            color: #fff;
            font-size: 15px;
            padding-right: 15px;
        }

        .up_video {
            width: 100px;
            height: 40px;
            background-color: #424242;
            border-radius: 6px;
            border: 2px solid #fff;
            color: #fff;
            font-size: 15px;
            right: 25px;
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
        }

        .pop_login {
            text-align: center;
            position: absolute;
            z-index: 1;
            background: #fff;
            /* 白色背景 */
            padding-top: 70px;
            padding-bottom: 20px;
            padding-right: 60px;
            border: 1px solid #000;
            /* 可根据需要添加样式 */
            z-index: 2;
            color: #424242;
            /* 使登录表单显示在遮罩层上面 */
            background-color: #ffffff
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

.center-video {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 10vh;
}

.videoBox {
    width: 100%;
    /* 设置弹窗的宽度，根据需要进行调整 */
}

.video-container {
    position: relative;
    padding-top: 56.25%;
    /* 16:9 宽高比，根据需要进行调整 */
    overflow: hidden;
}

.video1 {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

.button-container {
    display: flex;
    justify-content: space-between;
    margin-top: 40px;
    align-items: center;
}

// .button-left,
// .button-right {
//   flex: 1; /* 平均分配宽度 */
// }
.button-left {
    //   flex: 2; /* 左按钮占1/3的宽度 */
    margin-left: 120px;
}

.button-right {
    //   flex: 2; /* 右按钮占1/3的宽度 */
    margin-right: 120px;
}

.avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
}

.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px !important;
    position: relative !important;
    overflow: hidden !important;
}

.avatar-uploader .el-upload:hover {
    border: 1px dashed #d9d9d9 !important;
    border-color: #409eff;
}

.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 300px;
    height: 178px;
    line-height: 178px;
    text-align: center;
}

.avatar {
    width: 300px;
    height: 178px;
    display: block;
}
</style>