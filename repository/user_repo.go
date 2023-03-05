package repository

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"context"
)

type UserRepo interface {
	CheckLogin(c context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(c context.Context, user model.User) (model.User, error)
}
