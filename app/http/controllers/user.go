package controllers

import (
	"chatroom/app/http/form"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) Create(ctx *gin.Context) {
	ctx.ShouldBindJSON(form.RegisterForm{})
}

func (c *UserController) Login(ctx *gin.Context) {

}

func (c *UserController) Logout(ctx *gin.Context) {

}

func (c *UserController) Info(ctx *gin.Context) {

}

func (c *UserController) Update(ctx *gin.Context) {

}
