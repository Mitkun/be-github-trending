package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	log "backend-github-trending/log"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"github.com/labstack/echo/v4"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UseName:  "postgres",
		Password: "123456",
		DbName:   "github-trending",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
