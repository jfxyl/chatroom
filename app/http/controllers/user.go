package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (o *UserController) Create(c *gin.Context) {
	var (
		err  error
		errs map[string]string
		form requests.RegisterForm
		user models.User
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateRegisterForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	//birthday, _ := time.Parse("2006-01-02", form.Birthday)
	user = models.User{
		Name:     form.Name,
		Nickname: fmt.Sprintf("用户%s", common.RandString(6)),
		//Gender:   form.Gender,
		//Avatar:   form.Avatar,
		//Birthday: nil,
		Password: form.Password,
	}
	if global.DB.Save(&user).Error != nil {
		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
		return
	}
	common.RespOk(c, user.Transform())
}

func (o *UserController) Login(c *gin.Context) {
	var (
		err   error
		errs  map[string]string
		form  requests.LoginForm
		user  models.User
		token string
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateLoginForm(form)); errs != nil && len(errs) > 0 {
		fmt.Println(1)
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if err = global.DB.Where(map[string]any{"name": form.Name}).Limit(1).Find(&user).Error; err != nil {
		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
		return
	}
	fmt.Println(user, user.ID, user.ID == 0)
	if user.ID == 0 {
		common.RespFail(c, common.StatusNotFound, common.ERR_NOT_FOUND)
		return
	}
	fmt.Println(form.Password)
	if !user.CheckPassword(form.Password) {
		common.RespFail(c, common.StatusInvalidArgument, "密码不正确")
		return
	}
	fmt.Println("token")
	if token, err = user.GenerateJWT(); err != nil {
		fmt.Println(err)
		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
		return
	}
	fmt.Println("token", token)
	common.RespOk(c, user.LoginTransform(token))
}

func (o *UserController) Logout(c *gin.Context) {

}

func (o *UserController) Info(c *gin.Context) {
	var (
		err   error
		idStr string
		id    uint64
		user  models.User
	)
	idStr = c.Param("id")
	if idStr != "" {
		if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
			common.RespFail(c, common.StatusInvalidArgument, err.Error())
			return
		}
		if err = global.DB.Limit(1).Find(&user, id).Error; err != nil {
			common.RespFail(c, common.StatusInternal, err.Error())
			return
		}
		if user.ID == 0 {
			common.RespFail(c, common.StatusNotFound, common.ERR_NOT_FOUND)
			return
		}
		common.RespOk(c, user.Transform())
	} else {
		common.RespOk(c, auth.User(c).Transform())
	}
}

func (o *UserController) Update(c *gin.Context) {
	//	var (
	//		err   error
	//		errs  map[string]string
	//		idStr string
	//		id    uint64
	//		form  requests.RegisterForm
	//		user  models.User
	//	)
	//	idStr = c.Param("id")
	//	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
	//		common.RespFail(c, common.StatusInvalidArgument, err.Error())
	//		return
	//	}
	//	if err = global.DB.Limit(1).Find(&user, id).Error; err != nil {
	//		common.RespFail(c, common.StatusInternal, err.Error())
	//		return
	//	}
	//	if user.ID == 0 {
	//		common.RespFail(c, common.StatusNotFound, common.ERR_NOT_FOUND)
	//		return
	//	}
	//	if err = c.ShouldBindJSON(&form); err != nil {
	//		common.RespFail(c, common.StatusInvalidArgument, err)
	//		return
	//	}
	//	if errs = common.SimplifyError(requests.ValidateUserForm(form)); errs != nil && len(errs) > 0 {
	//		common.RespFail(c, common.StatusInvalidArgument, errs)
	//		return
	//	}
	//	birthday, _ := time.Parse("2006-01-02", form.Birthday)
	//	user = models.User{
	//		Name:     form.Name,
	//		Nickname: form.Nickname,
	//		Gender:   form.Gender,
	//		Avatar:   form.Avatar,
	//		Birthday: &birthday,
	//		Password: form.Password,
	//	}
	//	if global.DB.Save(&user).Error != nil {
	//		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
	//		return
	//	}
	//	common.RespOk(c, user.Transform())
}
