package services

import (
	"chatroom/internal/common"
	"chatroom/internal/config"
	"chatroom/internal/oss"
	"errors"
	"fmt"
	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"
)

func NewOssService() *OssService {
	return &OssService{}
}

type OssService struct {
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

func (s *OssService) Signature(dirtype string, contentType string) (string, *common.CodeErr) {
	var (
		err     error
		dir     string
		ext     string
		ok      bool
		signURL string
	)
	if dir, ok = dirMapping[dirtype]; !ok {
		return "", common.NewCodeErr(common.StatusInvalidArgument, errors.New("目录类型错误"))
	}
	var options = []alioss.Option{
		alioss.Expires(time.Now().Add(time.Hour)), // 设置签名过期时间
		alioss.ContentType(contentType),
	}
	if ext, err = getFileExtension(contentType); err != nil {
		return "", common.NewCodeErr(common.StatusInvalidArgument, err)
	}
	signURL, err = oss.G_OssBucket.SignURL(fmt.Sprintf("%s/%s/%s.%s", config.G_Config.Oss.PrefixDir, dir, common.GenerateRandomFilename(), ext), alioss.HTTPPut, 3600, options...)
	if err != nil {
		return "", common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return signURL, nil
}
