package services

import (
	"chatroom/app/http/requests"
	"chatroom/app/models"
	"chatroom/internal/auth"
	"chatroom/internal/common"
	"chatroom/internal/db"
	"errors"
	"github.com/gin-gonic/gin"
)

func NewRoomService() *RoomService {
	return &RoomService{}
}

type RoomService struct {
}

func (s *RoomService) Associate(user *models.User, room *models.Room) (err error) {
	userRoom := models.UserRoom{
		UserID: user.ID,
		RoomID: room.ID,
	}
	return db.G_DB.Save(&userRoom).Error
}

func (s *RoomService) Create(c *gin.Context, form requests.RoomForm) (*models.Room, *common.CodeErr) {
	var (
		err  error
		room *models.Room
	)
	room = &models.Room{
		Name:     form.Name,
		Avatar:   form.Avatar,
		Owner:    auth.User(c).ID,
		IsPublic: form.IsPublic == 1,
		//Users:    []*models.User{auth.User(c)},
	}
	if db.G_DB.Create(&room).Error != nil {
		common.RespFail(c, common.StatusInternal, common.ERR_INTERNAL_SERVER)
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if err = s.Associate(auth.User(c), room); err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return room, nil
}

func (s *RoomService) List(c *gin.Context) ([]*models.Room, *common.CodeErr) {
	var (
		err   error
		rooms []*models.Room
	)
	if err = db.G_DB.Model(auth.User(c)).Association("Rooms").Find(&rooms); err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return rooms, nil
}

func (s *RoomService) Delete(c *gin.Context, id uint64) *common.CodeErr {
	var (
		err  error
		room models.Room
	)
	if err = db.G_DB.Limit(1).Find(&room, id).Error; err != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if room.ID == 0 {
		return common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	if room.Owner != auth.User(c).ID {
		return common.NewCodeErr(common.StatusForbidden, common.ERR_FORBIDDEN)
	}
	//暂时保留用户和房间的关联关系，但用户不可访问该聊天室
	if db.G_DB.Delete(&room).Error != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return nil
}

func (s *RoomService) Info(c *gin.Context, id uint64) (*models.Room, *common.CodeErr) {
	var (
		err  error
		room models.Room
	)
	if err = db.G_DB.Limit(1).Find(&room, id).Error; err != nil {
		return nil, common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if room.ID == 0 {
		return nil, common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	return &room, nil
}

func (s *RoomService) Join(c *gin.Context, id uint64) *common.CodeErr {
	var (
		err  error
		room models.Room
	)
	if err = db.G_DB.Limit(1).Find(&room, id).Error; err != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if room.ID == 0 {
		return common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	if err = s.Associate(auth.User(c), &room); err != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return nil
}

func (s *RoomService) Quit(c *gin.Context, id uint64) *common.CodeErr {
	var (
		err  error
		room models.Room
	)
	if err = db.G_DB.Limit(1).Find(&room, id).Error; err != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	if room.ID == 0 {
		return common.NewCodeErr(common.StatusNotFound, common.ERR_NOT_FOUND)
	}
	if room.Owner == auth.User(c).ID {
		return common.NewCodeErr(common.StatusInvalidArgument, errors.New("房主暂不支持退出聊天室"))
	}
	if db.G_DB.Model(auth.User(c)).Association("Rooms").Delete(room) != nil {
		return common.NewCodeErr(common.StatusInternal, common.ERR_INTERNAL_SERVER)
	}
	return nil
}
