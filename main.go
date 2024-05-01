package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebsvt/ATNL001/handlers"
	"github.com/sebsvt/ATNL001/repositories"
	"github.com/sebsvt/ATNL001/services"
)

func main() {
	initTimeZone()
	// configs.InitEnvConfigs()

	db := initDB()

	userRepo := repositories.NewUserRepositoryDB(db)
	authSrv := services.NewAuthService()
	userSrv := services.NewUserService(userRepo, authSrv)
	userHandler := handlers.NewUserHandler(userSrv)

	app := fiber.New()
	api := app.Group("/api")

	userAPI := api.Group("/users")
	userAPI.Get("/:id", userHandler.GetUserFromID)
	userAPI.Post("/", userHandler.CreateNewUserAccount)

	app.Listen(":8000")

}

// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
// 	viper.AutomaticEnv()
// 	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDB() *sqlx.DB {
	dsn := "postgres://postgres:example@localhost:5432/my_db?sslmode=disable"
	// dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// 	configs.EnvConfigs.DATABASE_USERNAME,
	// 	configs.EnvConfigs.DATABASE_PASSWORD,
	// 	configs.EnvConfigs.DATABASE_HOST,
	// 	configs.EnvConfigs.DATABASE_PORT,
	// 	configs.EnvConfigs.DATABASE_NAME,
	// )
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
