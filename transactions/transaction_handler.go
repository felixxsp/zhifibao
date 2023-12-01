package transaction

import (
	"context"
	"net/http"
	"zhifubao/domain/entity"
	"zhifubao/domain/usecases"

	"github.com/gin-gonic/gin"
)

type RealHandler struct {
	router  *gin.Engine
	usecase usecases.Transaction_Usecase
}

type Handler interface {
	Standby(context.Context)
	NewTransaction(*gin.Context)
	ViewTransaction(*gin.Context)
	ViewMulti(*gin.Context)
}

func NewHandler(uc usecases.Transaction_Usecase, router *gin.Engine) *RealHandler {
	return &RealHandler{
		router:  router,
		usecase: uc,
	}
}

func (hdl *RealHandler) Standby(ctx context.Context) {
	hdl.router.POST("/newtransaction", hdl.NewTransaction)
	hdl.router.GET("/transaction/single", hdl.ViewTransaction)
	hdl.router.GET("/transaction/multi", hdl.ViewMulti)
}

func (hdl *RealHandler) NewTransaction(c *gin.Context) {
	var newtransaction entity.Transaction
	c.BindJSON(&newtransaction)
	code, err := hdl.usecase.NewTransaction(c, newtransaction)
	c.IndentedJSON(code, err.Error())
}

func (hdl *RealHandler) ViewTransaction(c *gin.Context) {
	var request entity.Trc_req_one
	c.BindJSON(&request)
	output, err := hdl.usecase.ViewTransaction(c, request)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, output)
	}
}

func (hdl *RealHandler) ViewMulti(c *gin.Context) {
	var request entity.Trc_req_multi
	c.BindJSON(&request)
	output, err := hdl.usecase.ViewMulti(c, request)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, output)
	}
}
