package person

import (
	"context"
	"fmt"
	"zhifubao/domain/entity"
	"zhifubao/domain/usecases"

	"github.com/gin-gonic/gin"
)

type RealHandler struct {
	router  *gin.Engine
	usecase usecases.Person_Usecase
}

type Handler interface {
	Standby(context.Context)
	NewPerson(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
}

func NewHandler(uc usecases.Person_Usecase, router *gin.Engine) *RealHandler {
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
	var NewPerson entity.Person
	c.BindJSON(&NewPerson)
	fmt.Println(NewPerson)
	code := hdl.usecase.NewPerson(c, NewPerson)
	c.IndentedJSON(code, "Insert Succesfull")
}

func (hdl *RealHandler) Login(c *gin.Context) {
	var newperson entity.Login_req
	c.BindJSON(&newperson)
	code := hdl.usecase.Login(c, newperson)
	c.IndentedJSON(code, "Successfully Login")
}

func (hdl *RealHandler) Logout(c *gin.Context) {
	var newperson entity.Login_req
	c.BindJSON(&newperson)
	code := hdl.usecase.Logout(c, newperson)
	c.IndentedJSON(code, "Succesfully Logout")
}
