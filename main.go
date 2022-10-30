package main

import (
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
	port := ":6000"

	db, err := db.ConnectGormDB()
	if err != nil {
		panic(err)
	}

	userRepo := gorm_postgres.NewUserRepoGormPostgres(db)
	userSvc := service.NewUserServices(userRepo)
	userHandler := controller.NewUserHandler(userSvc)

	productRepo := gorm_postgres.NewProductRepoGormPostgres(db)
	productSvc := service.NewProductServices(productRepo)
	productHandler := controller.NewProductHandler(productSvc)

	transactionRepo := gorm_postgres.NewTransactionRepoGormPostgres(db)
	transactionSvc := service.NewTransactionServices(transactionRepo, productRepo)
	transactionHandler := controller.NewTransactionHandler(transactionSvc)

	rajaOngkirRepo := gorm_postgres.NewRajaOngkirRepoGormPostgres(db)
	rajaOngkirSvc := service.NewRajaOngkirServices(rajaOngkirRepo)
	rajaOngkirHandler := controller.NewRajaOngkirHandler(rajaOngkirSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)
	app := server.NewRouterGin(router, userHandler, productHandler, transactionHandler, rajaOngkirHandler, middleware)
	app.Start(port)
}
