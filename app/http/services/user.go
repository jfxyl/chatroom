package services

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var defaultAvatars = [12]string{
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/mouse.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/cattle.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/tiger.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/rabbit.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/dragon.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/snake.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/horse.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/sheep.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/monkey.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/kun.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/dog.jpg",
	"https://jfxy.oss-cn-nanjing.aliyuncs.com/chatroom/user/pig.jpg",
}

func NewUserService() *UserService {
	return &UserService{}
}

type UserService struct {
}

func (s *UserService) Create(c *gin.Context, form requests.RegisterForm) (*models.User, *common.CodeErr) {
	var (
		err  error
		user *models.User
	)
	//birthday, _ := time.Parse("2006-01-02", form.Birthday)
	user = &models.User{
		Name:     form.Name,
		Nickname: fmt.Sprintf("用户%s", common.RandString(6)),
		Avatar:   defaultAvatars[common.RandInt(len(defaultAvatars))],
		//Gender:   form.Gender,
		//Birthday: nil,
		Password: form.Password,
	}
	if err = db.G_DB.Save(user).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return user, nil
}

func (s *UserService) Login(c *gin.Context, form requests.LoginForm) (string, *common.CodeErr) {
	var (
		err   error
		user  models.User
		token string
	)
	if err = db.G_DB.Where(map[string]any{"name": form.Name}).Limit(1).Find(&user).Error; err != nil {
		return "", common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if user.ID == 0 {
		return "", common.NewCodeErr(common.StatusInternal, common.ERR_NOT_FOUND)
	}
	if !user.CheckPassword(form.Password) {
		return "", common.NewCodeErr(common.StatusInvalidArgument, errors.New("密码不正确"))
	}
	if token, err = user.GenerateJWT(); err != nil {
		return "", common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return token, nil
}

func (s *UserService) Info(c *gin.Context, id uint64) (*models.User, *common.CodeErr) {
	var (
		err  error
		user models.User
	)
	if err = db.G_DB.Limit(1).Find(&user, id).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, err)
	}
	if user.ID == 0 {
		return nil, common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	return &user, nil
}

func (s *UserService) Update(c *gin.Context, form requests.UpdateForm) (*models.User, *common.CodeErr) {
	var (
		user *models.User
	)
	user = auth.User(c)
	birthday, _ := time.Parse("2006-01-02", form.Birthday)
	user.Name = form.Name
	user.Nickname = form.Nickname
	user.Gender = form.Gender
	user.Avatar = form.Avatar
	user.Birthday = &birthday
	if db.G_DB.Save(&user).Error != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return user, nil
}

func (s *UserService) UpdatePassword(c *gin.Context, form requests.UpdatePasswordForm) (*models.User, *common.CodeErr) {
	var (
		user *models.User
	)
	user = auth.User(c)
	user.Password = form.Password
	if db.G_DB.Save(&user).Error != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return user, nil
}
