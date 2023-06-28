package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/mytwitter/handlers/dto"
	"github.com/hi-supergirl/mytwitter/handlers/services"
	"github.com/hi-supergirl/mytwitter/logging"
	"github.com/hi-supergirl/mytwitter/middlewares"
	"go.uber.org/zap"
)

type TwitterUserHandler interface {
	// register a twitter user
	Register(c *gin.Context)
	// unregister a twitter use which has been registered
	Unregister(c *gin.Context)
	// Login for a twitter user
	Login(c *gin.Context)
	// logout for a twitter user
	Unlogin(c *gin.Context)
	// Check the details of one existing twitter user
	CheckUserInfo(c *gin.Context)
	// return the number of the existing twitter users
	CheckUserNumbers(c *gin.Context)
}

func NewTwitterUserHandler(logger *zap.Logger, twitterUserService services.TwitterUserService) TwitterUserHandler {
	return &twitterUserHandler{twitterUserService: twitterUserService}
}

type twitterUserHandler struct {
	twitterUserService services.TwitterUserService
}

func (tuh *twitterUserHandler) Register(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		var reg dto.RegisterDTO
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.Register start ...")
		if err := c.ShouldBindJSON(&reg); err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		tx, err := tuh.twitterUserService.Register(c.Request.Context(), reg.Name, reg.Password)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterUserHandler.Register ended")
		return NewResponse(http.StatusOK, success, gin.H{"tx": tx})
	})
}

func (tuh *twitterUserHandler) Unregister(c *gin.Context) {

}

func (tuh *twitterUserHandler) CheckUserInfo(c *gin.Context) {

}

func (tuh *twitterUserHandler) Login(c *gin.Context) {

}

func (tuh *twitterUserHandler) Unlogin(c *gin.Context) {

}

func (tuh *twitterUserHandler) CheckUserNumbers(c *gin.Context) {

}

func TwitterUserRoute(tuh TwitterUserHandler, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")

	api.Use(middlewares.RequestTraceMiddleWare())

	api.POST("/admin/users", tuh.Register)
	api.DELETE("/admin/users/:user", tuh.Unregister)
	api.PUT("/admin/users/:user/login", tuh.Login)
	api.PUT("/admin/users/:user/unlogin", tuh.Unlogin)
	api.GET("/admin/users/:user", tuh.CheckUserInfo)
	api.GET("/admin/users/", tuh.CheckUserNumbers)
}
