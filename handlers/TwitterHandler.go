package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/mytwitter/handlers/services"
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

}

func (th *twitterHander) PostNumbers(c *gin.Context) {

}

func (th *twitterHander) ReadPost(c *gin.Context) {

}

func TwitterRoute(th TwitterHander, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")

	api.Use(middlewares.RequestTraceMiddleWare())

	api.POST("/twitter/posts", th.Post)
	api.GET("/twitter/posts/:name", th.ReadPost)
	api.GET("/twitter/posts/:name/:index", th.ReadPost)
}
