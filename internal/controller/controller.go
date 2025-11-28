package controller

import (
	"fgrana/auth-project/internal/controller/routes"
	"fgrana/auth-project/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	engine  *gin.Engine
}

func NewHandler(srv service.Service, engine *gin.Engine) *Handler {
	h := &Handler{
		service: srv,
		engine:  engine,
	}

	h.engine.GET(routes.RoutePing, srv.Ping)
	h.engine.GET(routes.RoutesGetUser, srv.GetUser)
	h.engine.POST(routes.RoutesAddUser, srv.AddUser)
	h.engine.GET(routes.RoutesGetAllUsers, srv.GetAllUsers)
	return h
}
