package person

import (
	"context"
	"只服宝/entity"

	"github.com/gin-gonic/gin"
)

type RealHandler struct {
	router  *gin.Engine
	usecase Usecase
}

type Handler interface {
	Standby(context.Context)
	NewPerson(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
}

func NewHandler(uc Usecase, router *gin.Engine) *RealHandler {
	return &RealHandler{
		router:  router,
		usecase: uc,
	}
}

func (hdl *RealHandler) Standby(ctx context.Context) {
	hdl.router.POST("/newperson", hdl.NewPerson)
	hdl.router.PATCH("/login", hdl.Login)
	hdl.router.PATCH("/logout", hdl.Logout)
}

func (hdl *RealHandler) NewPerson(c *gin.Context) {
	var newperson entity.Person
	c.BindJSON(&newperson)
	code := hdl.usecase.NewPerson(c, newperson)
	c.IndentedJSON(code, "")
}

func (hdl *RealHandler) Login(c *gin.Context) {
	var newperson entity.DTO_login
	c.BindJSON(&newperson)
	code := hdl.usecase.Login(c, newperson)
	c.IndentedJSON(code, "")
}

func (hdl *RealHandler) Logout(c *gin.Context) {
	var newperson entity.DTO_login
	c.BindJSON(&newperson)
	code := hdl.usecase.Logout(c, newperson)
	c.IndentedJSON(code, "")
}
