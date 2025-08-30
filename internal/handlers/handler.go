package handlers

import "github.com/gin-gonic/gin"

type UserHandlers interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	CreateWallet(c *gin.Context)
	GetWallet(c *gin.Context)
}