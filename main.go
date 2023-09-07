package main

import (
	userHttp "go-crud/adapter/http/user"
	userRepository "go-crud/adapter/repository/user"
	"go-crud/domain/user"
	"go-crud/infra/gormdb"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := gormdb.SetupGormDB()
	if err != nil {
		panic(err)
	}

	userRepo := userRepository.NewGormUserRepository(db)
	userService := user.New(userRepo)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	group := r.Group("/api")

	userHttp.ConfigureRoutes(group, userService)

	r.Run() // listen and serve on 0.0.0.0:8080
}
