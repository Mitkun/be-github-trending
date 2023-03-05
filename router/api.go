package router

import (
	"backend-github-trending/handler"
	middleware "backend-github-trending/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {

	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUP)
	api.Echo.POST("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
