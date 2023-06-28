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
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.Unregister start...")
		user := c.Param("user")
		if user == "" {
			return NewResponse(http.StatusBadRequest, failed, "must provide a user in url")
		}

		tx, err := tuh.twitterUserService.Unregister(c.Request.Context(), user)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterUserHandler.Unregister edned")
		return NewResponse(http.StatusOK, success, gin.H{"tx": tx})
	})

}

func (tuh *twitterUserHandler) CheckUserInfo(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.CheckUserInfo start...")
		user := c.Param("user")
		if user == "" {
			return NewResponse(http.StatusBadRequest, failed, "must provide a user in url")
		}
		name, isLogin, err := tuh.twitterUserService.CheckUserInfo(c.Request.Context(), user)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}

		logger.Infoln("TwitterUserHandler.CheckUserInfo ended")
		return NewResponse(http.StatusOK, success, gin.H{"name": name, "isLogin": isLogin})
	})
}

func (tuh *twitterUserHandler) Login(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.Login start...")
		var login dto.LoginDTO
		if err := c.ShouldBindJSON(&login); err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		tx, err := tuh.twitterUserService.Login(c.Request.Context(), login.Name, login.Password)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}

		logger.Infoln("TwitterUserHandler.Login ended")
		return NewResponse(http.StatusOK, success, gin.H{"tx": tx})
	})
}

func (tuh *twitterUserHandler) Unlogin(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.Unlogin start...")
		user := c.Param("user")
		if user == "" {
			return NewResponse(http.StatusBadRequest, failed, "must provide a user in url")
		}

		tx, err := tuh.twitterUserService.Unlogin(c.Request.Context(), user)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterUserHandler.Unlogin ended")
		return NewResponse(http.StatusOK, success, gin.H{"tx": tx})
	})
}

func (tuh *twitterUserHandler) CheckUserNumbers(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterUserHandler.CheckUserNumbers start...")
		cnt, err := tuh.twitterUserService.CheckUserNumbers(c.Request.Context())
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterUserHandler.CheckUserNumbers ended")
		return NewResponse(http.StatusOK, success, gin.H{"users": cnt})
	})
}

func TwitterUserRoute(tuh TwitterUserHandler, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")

	api.Use(middlewares.RequestTraceMiddleWare())

	api.POST("/admin/users", tuh.Register)
	api.DELETE("/admin/users/:user", tuh.Unregister)
	api.PUT("/admin/users/login", tuh.Login)
	api.PUT("/admin/users/:user/unlogin", tuh.Unlogin)
	api.GET("/admin/users/:user", tuh.CheckUserInfo)
	api.GET("/admin/users/", tuh.CheckUserNumbers)
}
