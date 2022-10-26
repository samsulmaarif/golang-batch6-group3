package server

import (
	"golang-batch6-group3/server/controller"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router      *gin.Engine
	user        *controller.UserHandler
	product     *controller.ProductHandler
	transaction *controller.TransactionHandler
	middleware  *Middleware
}

func NewRouterGin(router *gin.Engine, user *controller.UserHandler, product *controller.ProductHandler, transaction *controller.TransactionHandler, middleware *Middleware) *GinRouter {
	return &GinRouter{
		router:      router,
		user:        user,
		product:     product,
		transaction: transaction,
		middleware:  middleware,
	}
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	users := r.router.Group("/users")
	users.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUsers, []string{"admin"}))
	users.GET("/email/:Email", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUserByEmail, []string{"admin"}))
	users.PUT("/profile", r.middleware.Auth, r.middleware.CheckRole(r.user.GinUpdateUserProfile, []string{"admin", "member"}))

	auth := r.router.Group("/auth")
	auth.POST("/login", r.user.GinLogin)
	auth.POST("/register", r.user.GinRegister)

	transactions := r.router.Group("/transactions")
	transactions.GET("/all", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinGetTransactions, []string{"admin"}))
	transactions.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinGetMemberTransactions, []string{"member"}))
	transactions.POST("/add", r.transaction.GinAddTransaction)

	products := r.router.Group("/products")
	products.GET("/", r.product.GinGetProducts)
	products.POST("/add", r.middleware.Auth, r.middleware.CheckRole(r.product.GinAddProduct, []string{"admin"}))
	products.DELETE("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.product.GinDeleteProduct, []string{"admin"}))
	products.PUT("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.product.GinUpdateProduct, []string{"admin"}))
	products.GET("/id/:Id", r.product.GetProductById)

	r.router.Run(port)
}

// package server

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// )

// func StartServer(router *http.ServeMux, port string, db *sql.DB) {
// 	buildRoute(router, db)
// 	fmt.Println("Server running at ", port)

// 	http.ListenAndServe(port, router)
// }

// func buildRoute(router *http.ServeMux, db *sql.DB) {

// }
