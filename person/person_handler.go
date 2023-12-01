package person

import (
	"context"
	"fmt"
	"net/http"
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
	GetFriends(*gin.Context)
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
	hdl.router.GET("/friends", hdl.GetFriends)
}

func (hdl *RealHandler) NewPerson(c *gin.Context) {
	var NewPerson entity.Person
	c.BindJSON(&NewPerson)
	fmt.Println(NewPerson)
	code := hdl.usecase.NewPerson(c, NewPerson)
	c.IndentedJSON(code, "Insert Succesfull")
}

func (hdl *RealHandler) Login(c *gin.Context) {
	var person entity.Login_req
	c.BindJSON(&person)
	code := hdl.usecase.Login(c, person)
	c.IndentedJSON(code, "Successfully Login")
}

func (hdl *RealHandler) Logout(c *gin.Context) {
	var person entity.Login_req
	c.BindJSON(&person)
	code := hdl.usecase.Logout(c, person)
	c.IndentedJSON(code, "Succesfully Logout")
}

func (hdl *RealHandler) GetFriends(c *gin.Context) {
	var person entity.Login_req
	c.BindJSON(&person)
	result := hdl.usecase.GetFriends(c, person)
	c.IndentedJSON(http.StatusOK, result)
}
