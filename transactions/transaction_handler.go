package transaction

import (
	"context"

	"github.com/gin-gonic/gin"
)

type RealHandler struct {
	router  *gin.Engine
	usecase *Usecase
}

type Handler interface {
	Standby(context.Context)
	IsActive()
	NewTransaction()
	ViewTransaction()
	ViewAll()
}

func NewHandler(uc *Usecase, router *gin.Engine) *RealHandler {
	return &RealHandler{
		router:  router,
		usecase: uc,
	}
}

func (db *RealHandler) Standby(ctx context.Context) {

}

func (db *RealHandler) IsActive() {

}

func (db *RealHandler) NewTransaction() {

}

func (db *RealHandler) ViewTransaction() {

}

func (db *RealHandler) ViewAll() {

}
