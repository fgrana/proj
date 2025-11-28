package service

import (
	"fgrana/auth-project/internal/model"
	"fgrana/auth-project/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	constName    = "name"
	constPong    = "pong"
	constValue   = "value"
	constNoValue = "no value"
	constUser    = "user"
	constStatus  = "status"
)

type Service struct {
	db repository.DB
}

func NewService(db repository.DB) *Service {
	s := &Service{
		db: db,
	}
	return s
}

func (s Service) Ping(c *gin.Context) {
	c.String(http.StatusOK, constPong)
}

func (s Service) GetUser(c *gin.Context) {
	value, ok := s.db.GetUser(c)
	if ok {
		c.JSON(http.StatusOK, gin.H{constUser: c.Params.ByName(constName), constValue: value})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{constUser: c.Params.ByName(constName), constStatus: constNoValue})
	}
}

func (s Service) AddUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	if c.Bind(&model.Json) == nil {
		if user == model.Json.Name {
			_, ok := s.db.AddUser(user)
			if !ok {
				fmt.Println("could not add user")
			}
			fmt.Printf("%s was successfully added", ok)
			c.JSON(http.StatusOK, gin.H{constStatus: "ok"})
		}
	}
}

func (s Service) GetAllUsers(c *gin.Context) {
	value, ok := s.db.GetAllUsers(c)
	if ok {
		c.JSON(http.StatusOK, gin.H{constUser: c.Params.ByName(constName), constValue: value})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{constUser: c.Params.ByName(constName), constStatus: constNoValue})
	}

}
