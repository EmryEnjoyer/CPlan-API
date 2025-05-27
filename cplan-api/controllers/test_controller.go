package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseFunction(context *gin.Context) {
	user_profile, _ := context.Get("user_profile")
	context.IndentedJSON(http.StatusOK, user_profile)
}

func PublicEndpoint(context *gin.Context) {
	context.String(http.StatusOK, "Public endpoint for you to land at")
}
