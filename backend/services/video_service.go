package services

import (
	"app/consts"
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func UploadVideo(localPath string, key string) (videoUrl string, err error) {
	//存储空间名称
	bucket := "lips"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 创建七牛云认证对象
	mac := qbox.NewMac(consts.AccessKey, consts.SecretKey)
	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	// 创建存储配置对象
	cfg := storage.Config{}
	// 创建表单上传对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 创建PutExtra对象，用于指定文件在存储空间中的键名
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"key": key,
		},
	}
	// 上传视频到 七牛云
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localPath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	videoUrl = "http://s3315jx3y.hn-bkt.clouddn.com/" + key
	fmt.Println("上传成功，视频访问链接：", "http://s3315jx3y.hn-bkt.clouddn.com/"+key)
	return
}

// GetSnapshot 生成视频缩略图并保存（作为封面）
func GetSnapshot(videoPath, snapshotPath string, title string, frameNum int) (snapshotName string) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	err = imaging.Save(img, snapshotPath+title+".jpeg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	snapshotName = fmt.Sprintf("./uploads/cover/" + title + ".jpeg")
	return
}

func UploadCover(localPath string, key string) (coverUrl string, err error) {
	//存储空间名称
	bucket := "lips"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 创建七牛云认证对象
	mac := qbox.NewMac(consts.AccessKey, consts.SecretKey)
	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	// 创建存储配置对象
	cfg := storage.Config{}
	// 创建表单上传对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 创建PutExtra对象，用于指定文件在存储空间中的键名
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"key": key,
		},
	}
	// 上传截图到七牛云
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localPath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	coverUrl = "http://s3315jx3y.hn-bkt.clouddn.com/" + key
	fmt.Println("上传成功，截图访问链接：", "http://s3315jx3y.hn-bkt.clouddn.com/"+key)
	return
}

// 搜索
