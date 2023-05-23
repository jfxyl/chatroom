package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/common"
	"chatroom/internal/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
}

func (o *UserController) Create(c *gin.Context) {
	var (
		err    error
		errs   map[string]string
		form   requests.RegisterForm
		user   models.User
		dbResp *gorm.DB
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, http.StatusBadRequest, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateUserForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, http.StatusBadRequest, errs)
		return
	}
	birthday, _ := time.Parse("2006-01-02", form.Birthday)
	user = models.User{
		Name:     form.Name,
		Gender:   form.Gender,
		Avatar:   form.Avatar,
		Birthday: &birthday,
		Password: form.Password,
	}
	if dbResp = global.DB.Save(&user); dbResp.Error != nil {
		common.RespFail(c, http.StatusInternalServerError, dbResp.Error.Error())
		return
	}
	common.RespOk(c, user.Transform())
}

func (o *UserController) Login(c *gin.Context) {

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
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = global.DB.Limit(1).Find(&user, id).Error; err != nil {
		common.RespFail(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.ID == 0 {
		common.RespFail(c, http.StatusNotFound, "用户不存在")
		return
	}
	common.RespOk(c, user.Transform())
}

func (o *UserController) Update(c *gin.Context) {

}
