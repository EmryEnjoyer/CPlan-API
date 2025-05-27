package main

import (
	"encoding/gob"
	"log"
	"os"

	"cplan-api/auth"
	"cplan-api/auth/middleware"
	"cplan-api/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	authenticator, auth_err := auth.New()
	if auth_err != nil {
		log.Fatal("ERROR: Failed to initialize Authenticator")
		return
	}
	router := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	router.Use(sessions.Sessions("auth-session", store))

	gob.Register(map[string]interface{}{})
	router.GET("/", middleware.IsAuthenticated(authenticator), middleware.GetUserProfile, controllers.BaseFunction)
	router.GET("/auth/login", auth.LoginHandler(authenticator))
	router.GET("/auth/logout", auth.LogoutHandler)
	router.GET("/auth/logout_callback", auth.LogoutCallbackHandler(store))
	router.GET("/callback", auth.AuthenticationCallbackHandler(authenticator))
	router.GET("/public", controllers.PublicEndpoint)
	router.Run("localhost:8080")
}
