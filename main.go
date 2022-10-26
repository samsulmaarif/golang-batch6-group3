package main

import (
	"golang-batch6-group3/adaptor"
	"golang-batch6-group3/db"
	"golang-batch6-group3/server"
	"golang-batch6-group3/server/controller"
	"golang-batch6-group3/server/repository/gorm_postgres"
	"golang-batch6-group3/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	// router := http.NewServeMux()
	port := ":6000"

	// db := config.ConnectDB()
	db, err := db.ConnectGormDB()
	if err != nil {
		panic(err)
	}
	typicodeAdaptor := adaptor.NewTypicodeAdaptor("https://jsonplaceholder.typicode.com/posts")

	userRepo := gorm_postgres.NewUserRepoGormPostgres(db)
	userSvc := service.NewUserServices(userRepo, typicodeAdaptor)
	userHandler := controller.NewUserHandler(userSvc)
	// server.StartServer(router, port, db)

	productRepo := gorm_postgres.NewProductRepoGormPostgres(db)
	productSvc := service.NewProductServices(productRepo, typicodeAdaptor)
	productHandler := controller.NewProductHandler(productSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)
	app := server.NewRouterGin(router, userHandler, productHandler, middleware)
	app.Start(port)
}
