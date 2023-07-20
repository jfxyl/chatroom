package controllers

import (
	"chatroom/app/http/requests"
	"chatroom/app/http/services"
	"chatroom/app/models"
	"chatroom/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewRoomController() *RoomController {
	return &RoomController{
		RoomService: services.NewRoomService(),
	}
}

type RoomController struct {
	RoomService *services.RoomService
}

func (o *RoomController) Create(c *gin.Context) {
	var (
		err     error
		errs    map[string]string
		codeErr *common.CodeErr
		form    requests.RoomForm
		room    *models.Room
	)
	if err = c.ShouldBindJSON(&form); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err)
		return
	}
	if errs = common.SimplifyError(requests.ValidateCreateRoomForm(form)); errs != nil && len(errs) > 0 {
		common.RespFail(c, common.StatusInvalidArgument, errs)
		return
	}
	if room, codeErr = o.RoomService.Create(c, form); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, room.Transform())
}

func (o *RoomController) List(c *gin.Context) {
	var (
		codeErr        *common.CodeErr
		room           *models.Room
		rooms          []*models.Room
		transformRooms []map[string]any
	)
	if rooms, codeErr = o.RoomService.List(c); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	for _, room = range rooms {
		transformRooms = append(transformRooms, room.Transform())
	}
	common.RespOk(c, transformRooms)
}

func (o *RoomController) Delete(c *gin.Context) {
	var (
		err     error
		idStr   string
		id      uint64
		codeErr *common.CodeErr
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	if codeErr = o.RoomService.Delete(c, id); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, nil)
}

func (o *RoomController) Info(c *gin.Context) {
	var (
		err     error
		codeErr *common.CodeErr
		idStr   string
		id      uint64
		room    *models.Room
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	if room, codeErr = o.RoomService.Info(c, id); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, room.Transform())
}

func (o *RoomController) Find(c *gin.Context) {
	var (
		err     error
		codeErr *common.CodeErr
		field   string
		content string
		id      uint64
		room    *models.Room
		query   = make(map[string]any)
	)
	field = c.Query("field")
	content = c.Query("content")
	fmt.Println(field, content)
	if field == "id" {
		if id, err = strconv.ParseUint(content, 10, 64); err != nil {
			common.RespFail(c, common.StatusInvalidArgument, "参数错误")
			return
		}
		query["id"] = id
	} else if field == "name" {
		query["name"] = content
	} else {
		common.RespFail(c, common.StatusInvalidArgument, "参数错误")
		return
	}
	fmt.Println("query", query)
	if room, codeErr = o.RoomService.Find(c, query); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, room.Transform())
}

func (o *RoomController) Join(c *gin.Context) {
	var (
		err     error
		codeErr *common.CodeErr
		idStr   string
		id      uint64
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	if codeErr = o.RoomService.Join(c, id); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, nil)
}

func (o *RoomController) Quit(c *gin.Context) {
	var (
		err     error
		codeErr *common.CodeErr
		idStr   string
		id      uint64
	)
	idStr = c.Param("id")
	if id, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		common.RespFail(c, common.StatusInvalidArgument, err.Error())
		return
	}
	if codeErr = o.RoomService.Quit(c, id); codeErr != nil {
		common.RespFail(c, codeErr.Code, codeErr.Err)
		return
	}
	common.RespOk(c, nil)
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
	//		common.RespFail(c, common.StatusInvalidArgument, err.Error())
	//		return
	//	}
	//	if err = db.G_DB.Limit(1).Find(&user, id).Error; err != nil {
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
	//	if db.G_DB.Save(&user).Error != nil {
	//		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
	//		return
	//	}
	//	common.RespOk(c, user.Transform())
}
