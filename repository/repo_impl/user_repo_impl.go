package repo_impl

import (
	"backend-github-trending/banana"
	"backend-github-trending/db"
	"backend-github-trending/log"
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(c context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(c, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) CheckLogin(c context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var (
		user model.User
		err  error
	)
	err = u.sql.Db.GetContext(c, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (u *UserRepoImpl) SelectUserById(c context.Context, userId string) (model.User, error) {
	var (
		user model.User
		err  error
	)
	err = u.sql.Db.GetContext(c, &user, "SELECT * FROM users WHERE user_id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}

func (u UserRepoImpl) UpdateUser(c context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			full_name  = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(c, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserNotUpdated
	}

	return user, nil
}

func (g GithubRepoImpl) SelectAllBookmarks(c context.Context, userId string) ([]model.GithubRepo, error) {
	repos := []model.GithubRepo{}

	err := g.sql.Db.SelectContext(c, &repos,
		`SELECT 
					repos.name, repos.description, repos.url, 
					repos.color, repos.lang, repos.fork, repos.stars, 
					repos.stars_today, repos.build_by, true as bookmarked
				FROM bookmarks 
				INNER JOIN repos
				ON bookmarks.user_id=$1 AND repos.name = bookmarks.repo_name`, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return repos, banana.BookmarkNotFound
		}
		log.Error(err.Error())
	}

	return repos, nil
}
