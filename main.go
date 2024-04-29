package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebsvt/ATNL001/handlers"
	"github.com/sebsvt/ATNL001/repositories"
	"github.com/sebsvt/ATNL001/services"
)

func main() {

	db := initDB()

	userRepo := repositories.NewUserRepositoryDB(db)
	authSrv := services.NewAuthService()
	userSrv := services.NewUserService(userRepo, authSrv)
	userHandler := handlers.NewUserHandler(userSrv)

	app := fiber.New()
	api := app.Group("/api")

	userAPI := api.Group("/users")
	userAPI.Get("/:id", userHandler.GetUserFromID)

	app.Listen(":8000")

}

func initDB() *sqlx.DB {
	dsn := "postgres://postgres:example@localhost:5432/my_db?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
