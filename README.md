# 第二届七牛云1024创作节项目--三人成行队
>架构文档: https://awq7m8b63wy.feishu.cn/docx/BgyXdR7CToDXw8xCZaccdTswnHe
## 1. 简介
这是一款基于七牛云存储和七牛视频产品的创新应用，旨在为用户提供一个交互式的短视频观看平台，通过前后端分离的方式实现。

前端**自行使用typescript封装了一套node后端框架**，采用了vue、webpack、axios、element-ui等技术,以实现视频的播放、暂停、进度条拖拽等基础功能，并提供内容分类页，如热门视频、体育频道等，使用户能够方便地浏览和查看感兴趣的视频。

后端部分我们选择了Go语言作为开发语言，并采用了Gin框架和GORM库，以快速构建高性能的API接口，还使用阿里云上的数据库和Redis库，以确保数据的可靠性和高效性。

除了基础功能外，我们还提供了一些高级功能供用户选择，如账户系统，允许用户注册登录并收藏视频，以便他们能够保存自己喜欢的视频。我们还参考常见的短视频应用，自由增加功能，提升了应用的完善度，例如点赞、收藏、关注、搜索等功能。

我们致力于通过创新的技术和出色的用户界面，为用户带来乐趣和娱乐。无论是在休闲时间还是在寻找精彩视频的时候，我们的应用都是用户的理想选择。

## 2. 贡献者
### 三人成行队
- [jujubest](https://github.com/jujubest)(队长)
- [ql17774641928](https://github.com/ql17774641928)
- [Mirrre](https://github.com/Mirrre)

## 3. 使用说明
### 3.1 开发环境配置
#### 前端
1.  vue         v2.6.11
2.  vue-router  v3.2.0
#### 后端
1.  go     v1.20
2.  gin    v1.9.1
3.  gorm   v1.25.5
4.  mysql  v1.5.2
5.  redis  v9.2.1
6.  下载ffmpeg (生成封面功能）
下载地址 ：  https://ffbinaries.com/downloads

下载好对应的版本后将`ffmpeg.exe`文件放到本地Go环境的bin目录下
### 3.2 部署环境
本地部署
### 3.3 开启服务
拉取本仓库的`main`分支到本地，
#### 前端
进入到`frontend`目录
```bash
cd fronotend
```
运行`npm`
```bash
npm run serve
```
#### 后端
1.进入到`backend`目录
```bash
cd backend
```
运行`main.go`文件
```bash
go run main.go
```
2.进入到`backend_node`目录
```bash
cd backend_node
```
运行`yarn`
```bash
run start
```
即可开启所有服务

