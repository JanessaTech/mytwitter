package handlers

import "github.com/gin-gonic/gin"

type ResponseMessage string

const (
	success = ResponseMessage("success")
	failed  = ResponseMessage("failed")
)

type Response struct {
	StatusCode int             `json:"code"`
	Message    ResponseMessage `json:"message"`
	Details    interface{}     `json:"details"`
}

func NewResponse(statusCode int, message ResponseMessage, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Details:    data,
	}
}

func execute(c *gin.Context, f func(c *gin.Context) *Response) {
	ctx := c.Request.Context()

	doneChan := make(chan *Response)
	go func() {
		doneChan <- f(c)
	}()

	select {
	case <-ctx.Done():
		// no defined yet
	case res := <-doneChan:
		handleResponse(c, res)
	}
}

func handleResponse(c *gin.Context, r *Response) {
	c.JSON(r.StatusCode, r)
}
