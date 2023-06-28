package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/mytwitter/handlers/dto"
	"github.com/hi-supergirl/mytwitter/handlers/services"
	"github.com/hi-supergirl/mytwitter/helper"
	"github.com/hi-supergirl/mytwitter/logging"
	"github.com/hi-supergirl/mytwitter/middlewares"
	"go.uber.org/zap"
)

type TwitterHander interface {
	// post message by a twitter user. The user who attempts to post message must have logined
	Post(c *gin.Context)
	// calculate how many posts under one twitter user
	PostNumbers(c *gin.Context)
	// read the content of post under one twitter user
	ReadPost(c *gin.Context)
}

func NewTwitterHander(logger *zap.Logger, twitterService services.TwitterService) TwitterHander {
	return &twitterHander{twitterService: twitterService}
}

type twitterHander struct {
	twitterService services.TwitterService
}

func (th *twitterHander) Post(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterHander.Post start...")
		var post dto.PostDTO
		if err := c.ShouldBindJSON(&post); err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		tx, err := th.twitterService.Post(c.Request.Context(), post.Name, post.Message)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterHander.Post ended")
		return NewResponse(http.StatusOK, success, gin.H{"tx": tx})
	})
}

func (th *twitterHander) PostNumbers(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterHander.PostNumbers start...")
		user := c.Param("user")
		if user == "" {
			return NewResponse(http.StatusBadRequest, failed, "must provide user")
		}
		posts, err := th.twitterService.PostNumbers(c.Request.Context(), user)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterHander.PostNumbers ended")
		return NewResponse(http.StatusOK, success, gin.H{"posts": posts})
	})

}

func (th *twitterHander) ReadPost(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infoln("TwitterHander.ReadPost start...")
		user := c.Param("user")
		if user == "" {
			return NewResponse(http.StatusBadRequest, failed, "must provide user")
		}
		index := helper.StringToUint(c.Param("index"))
		name, message, timestamp, err := th.twitterService.ReadPost(c.Request.Context(), user, index)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err.Error())
		}
		logger.Infoln("TwitterHander.ReadPost ended")
		return NewResponse(http.StatusOK, success, gin.H{"name": name, "message": message, "timestamp": timestamp})
	})
}

func TwitterRoute(th TwitterHander, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")

	api.Use(middlewares.RequestTraceMiddleWare())

	api.POST("/twitter/posts", th.Post)
	api.GET("/twitter/posts/:user", th.PostNumbers)
	api.GET("/twitter/posts/:user/:index", th.ReadPost)
}
