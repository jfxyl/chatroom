package oss

import (
	"chatroom/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	G_OssClient *oss.Client
	G_OssBucket *oss.Bucket
)

func InitOss() (err error) {
	// 创建 OSS 客户端
	G_OssClient, err = oss.New(config.G_Config.Oss.Entpoint, config.G_Config.Oss.AccessKeyID, config.G_Config.Oss.AccessKeySecret)
	if err != nil {
		return err
	}
	G_OssBucket, err = G_OssClient.Bucket(config.G_Config.Oss.BucketName)
	return err
}
