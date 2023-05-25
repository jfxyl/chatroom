package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoomController struct {
}

func (o *RoomController) Create(c *gin.Context) {
	var (
		err  error
		errs map[string]string
		form requests.RoomForm
		room models.Room
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, http.StatusBadRequest, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateCreateRoomForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, http.StatusBadRequest, errs)
		return
	}
	room = models.Room{
		Name:     form.Name,
		Owner:    auth.User(c).ID,
		IsPublic: form.IsPublic == 1,
		Users:    []*models.User{auth.User(c)},
	}
	if global.DB.Create(&room).Error != nil {
		common.RespFail(c, http.StatusInternalServerError, common.ERR_INTERNAL_SERVER)
		return
	}
	common.RespOk(c, room.Transform())
}

func (o *RoomController) List(c *gin.Context) {
	var (
		room           *models.Room
		rooms          []*models.Room
		transformRooms []map[string]any
	)
	if global.DB.Model(auth.User(c)).Association("Rooms").Find(&rooms) != nil {
		common.RespFail(c, http.StatusInternalServerError, common.ERR_INTERNAL_SERVER)
		return
	}
	for _, room = range rooms {
		transformRooms = append(transformRooms, room.Transform())
	}
	common.RespOk(c, transformRooms)
}

func (o *RoomController) Detele(c *gin.Context) {
	var (
		err   error
		idStr string
		id    uint64
		room  models.Room
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = global.DB.Limit(1).Find(&room, id).Error; err != nil {
		common.RespFail(c, http.StatusInternalServerError, common.ERR_INTERNAL_SERVER)
		return
	}
	if room.ID == 0 {
		common.RespFail(c, http.StatusNotFound, common.ERR_NOT_FOUND)
		return
	}
	if room.Owner != auth.User(c).ID {
		common.RespFail(c, http.StatusNotFound, common.ERR_FORBIDDEN)
		return
	}
	if global.DB.Delete(&room).Error != nil {
		common.RespFail(c, http.StatusInternalServerError, common.ERR_INTERNAL_SERVER)
		return
	}
	common.RespOk(c, nil)
}

func (o *RoomController) Info(c *gin.Context) {
	var (
		err   error
		idStr string
		id    uint64
		user  models.User
	)
	idStr = c.Param("id")
	if idStr != "" {
		if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
			common.RespFail(c, http.StatusBadRequest, err.Error())
			return
		}
		if err = global.DB.Limit(1).Find(&user, id).Error; err != nil {
			common.RespFail(c, http.StatusInternalServerError, err.Error())
			return
		}
		if user.ID == 0 {
			common.RespFail(c, http.StatusNotFound, common.ERR_NOT_FOUND)
			return
		}
		common.RespOk(c, user.Transform())
	} else {
		common.RespOk(c, auth.User(c).Transform())
	}
}

func (o *RoomController) Update(c *gin.Context) {
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
	//		common.RespFail(c, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//	if err = global.DB.Limit(1).Find(&user, id).Error; err != nil {
	//		common.RespFail(c, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//	if user.ID == 0 {
	//		common.RespFail(c, http.StatusNotFound, common.ERR_NOT_FOUND)
	//		return
	//	}
	//	if err = c.ShouldBindJSON(&form); err != nil {
	//		common.RespFail(c, http.StatusBadRequest, err)
	//		return
	//	}
	//	if errs = common.SimplifyError(requests.ValidateUserForm(form)); errs != nil && len(errs) > 0 {
	//		common.RespFail(c, http.StatusBadRequest, errs)
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
	//		common.RespFail(c, http.StatusInternalServerError, common.ERR_INTERNAL_SERVER)
	//		return
	//	}
	//	common.RespOk(c, user.Transform())
}
