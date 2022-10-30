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
	rajaongkir  *controller.RajaOngkirHandler
	middleware  *Middleware
}

func NewRouterGin(router *gin.Engine, user *controller.UserHandler, product *controller.ProductHandler, transaction *controller.TransactionHandler,
	rajaongkir *controller.RajaOngkirHandler, middleware *Middleware) *GinRouter {
	return &GinRouter{
		router:      router,
		user:        user,
		product:     product,
		transaction: transaction,
		rajaongkir:  rajaongkir,
		middleware:  middleware,
	}
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	users := r.router.Group("/users")
	users.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUsers, []string{"admin"}))
	users.GET("/email/:Email", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUserByEmail, []string{"admin"}))
	users.PUT("/profile", r.middleware.Auth, r.middleware.CheckRole(r.user.GinUpdateUserProfile, []string{"admin", "member"}))
	users.DELETE("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.user.GinDeleteUser, []string{"admin"}))
	users.GET("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUserById, []string{"admin"}))
	users.POST("/add", r.middleware.Auth, r.middleware.CheckRole(r.user.GinAddUser, []string{"admin"}))

	auth := r.router.Group("/auth")
	auth.POST("/login", r.user.GinLogin)
	auth.POST("/register", r.user.GinRegister)

	transactions := r.router.Group("/transactions")
	transactions.GET("/all", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinGetTransactions, []string{"admin"}))
	transactions.PUT("/package/:Id", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinPackageTransactions, []string{"admin", "kasir"}))
	transactions.PUT("/send/:Id", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinSendTransactions, []string{"admin", "kasir"}))
	transactions.PUT("/confirm/:Id", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinConfirmTransactions, []string{"member"}))
	transactions.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinGetMemberTransactions, []string{"member"}))
	transactions.POST("/add", r.middleware.Auth, r.middleware.CheckRole(r.transaction.GinAddTransaction, []string{"member"}))

	products := r.router.Group("/products")
	products.GET("/", r.product.GinGetProducts)
	products.POST("/add", r.middleware.Auth, r.middleware.CheckRole(r.product.GinAddProduct, []string{"admin"}))
	products.DELETE("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.product.GinDeleteProduct, []string{"admin"}))
	products.PUT("/id/:Id", r.middleware.Auth, r.middleware.CheckRole(r.product.GinUpdateProduct, []string{"admin"}))
	products.GET("/id/:Id", r.product.GetProductById)

	rajaongkirs := r.router.Group("/rajaongkirs")
	rajaongkirs.GET("/city/", r.rajaongkir.GetCityById)
	rajaongkirs.GET("/province/", r.rajaongkir.GetProvinceById)
	rajaongkirs.GET("/cost/", r.rajaongkir.GetCost)

	r.router.Run(port)
}
