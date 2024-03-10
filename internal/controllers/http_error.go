package controllers

import "github.com/gin-gonic/gin"

func NewError(c *gin.Context, statusCode int, err error) {
	httpError := HttpError{
		Code:  statusCode,
		Error: err.Error(),
	}

	c.JSON(statusCode, httpError)
}

type HttpError struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"400"`
}
