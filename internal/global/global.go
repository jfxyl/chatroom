package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	OssClient *oss.Client
	OssBucket *oss.Bucket
)
