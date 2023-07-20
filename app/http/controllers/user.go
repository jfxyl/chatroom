package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/http/services"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewUserController() *UserController {
	return &UserController{
		UserService: services.NewUserService(),
	}
}

type UserController struct {
	UserService *services.UserService
}

func (o *UserController) Create(c *gin.Context) {
	var (
		err     error
		codeErr *common.CodeErr
		errs    map[string]string
		form    requests.RegisterForm
		user    *models.User
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateRegisterForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if user, codeErr = o.UserService.Create(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, user.Transform())
}

func (o *UserController) Login(c *gin.Context) {
	var (
		err     error
		errs    map[string]string
		codeErr *common.CodeErr
		form    requests.LoginForm
		user    models.User
		token   string
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateLoginForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if token, codeErr = o.UserService.Login(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, user.LoginTransform(token))
}

func (o *UserController) Logout(c *gin.Context) {

}

func (o *UserController) Info(c *gin.Context) {
	var (
		err     error
		idStr   string
		id      uint64
		codeErr *common.CodeErr
		user    *models.User
	)
	idStr = c.Param("id")
	if idStr != "" {
		if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
			common.RespFail(c, common.StatusInvalidArgument, err.Error())
			return
		}
		if user, codeErr = o.UserService.Info(c, id); codeErr != nil {
			common.RespFail(c, codeErr.Code, codeErr.Err)
			return
		}
		common.RespOk(c, user.Transform())
	} else {
		common.RespOk(c, auth.User(c).Transform())
	}
}

func (o *UserController) Update(c *gin.Context) {
	var (
		err     error
		errs    map[string]string
		form    requests.UpdateForm
		user    *models.User
		codeErr *common.CodeErr
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateUserForm(form, c)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if user, codeErr = o.UserService.Update(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, user.Transform())
}

func (o *UserController) UpdatePassword(c *gin.Context) {
	var (
		err     error
		errs    map[string]string
		form    requests.UpdatePasswordForm
		user    *models.User
		codeErr *common.CodeErr
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateUpdatePasswordForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if user, codeErr = o.UserService.UpdatePassword(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, user.Transform())
}
