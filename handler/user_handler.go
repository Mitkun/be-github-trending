package handler

import (
	"backend-github-trending/log"
	"backend-github-trending/model"
	req "backend-github-trending/model/req"
	"backend-github-trending/repository"
	"backend-github-trending/security"
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Mitkun",
		"email": "mitkun@gmail.com",
	})
}
func (u *UserHandler) HandleSignUP(c echo.Context) error {
	var (
		err  error
		req  req.ReqSignUp
		user model.User
		hash string
		role string
	)
	//req := req2.ReqSignUp{}
	if err = c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validator := validator.New()
	err = validator.Struct(req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash = security.HashAndSalt([]byte(req.Password))
	role = model.MEMBER.String()
	userId, err := uuid.NewUUID()

	user = model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Password = ""

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       user,
	})
}
