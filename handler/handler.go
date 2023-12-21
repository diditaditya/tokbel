package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	Path     string
	Function gin.HandlerFunc
}
