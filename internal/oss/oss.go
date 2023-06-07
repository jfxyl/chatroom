package oss

import (
	"chatroom/internal/common"
	"chatroom/internal/config"
	"chatroom/internal/global"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"
)

func InitOss() (err error) {
	// 创建 OSS 客户端
	global.OssClient, err = oss.New(config.G_Config.Oss.Entpoint, config.G_Config.Oss.AccessKeyID, config.G_Config.Oss.AccessKeySecret)
	if err != nil {
		return err
	}
	global.OssBucket, err = global.OssClient.Bucket(config.G_Config.Oss.BucketName)
	return err
}

const USER_AVATAR_DIR = "user/avatar"
const ROOM_AVATAR_DIR = "room/avatar"
const MSG_IMG_DIR = "msg/img"

var dirMapping = map[string]string{
	"user_avatar": USER_AVATAR_DIR,
	"room_avatar": ROOM_AVATAR_DIR,
	"msg_img":     MSG_IMG_DIR,
}

var contentTypeMapping = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
	"image/bmp":  ".bmp",
}

func getFileExtension(contentType string) (ext string, err error) {
	// 如果有对应的文件类型映射，则返回文件后缀；否则返回默认后缀 ".dat"
	if ext, ok := contentTypeMapping[contentType]; ok {
		return ext, nil
	}
	return "", errors.New("不支持的文件类型")
}

func Signature(dirtype string, contentType string) (signURL string, err error) {
	var (
		dir string
		ext string
		ok  bool
	)
	if dir, ok = dirMapping[dirtype]; !ok {
		return "", errors.New("目录类型错误")
	}
	var options = []oss.Option{
		oss.Expires(time.Now().Add(time.Hour)), // 设置签名过期时间
		oss.ContentType(contentType),
	}
	if ext, err = getFileExtension(contentType); err != nil {
		return "", err
	}
	signURL, err = global.OssBucket.SignURL(fmt.Sprintf("%s/%s/%s.%s", config.G_Config.Oss.PrefixDir, dir, common.GenerateRandomFilename(), ext), oss.HTTPPut, 3600, options...)
	fmt.Print(signURL)
	return
}
