package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClains{},
		SigningKey: security.SECRET_KEY,
	}

	return middleware.JWTWithConfig(config)
}
