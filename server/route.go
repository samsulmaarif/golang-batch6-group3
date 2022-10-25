package server

import (
	"golang-batch6-group3/server/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
	user   *controller.UserHandler
}

type GinRouter struct {
	router     *gin.Engine
	user       *controller.UserHandler
	middleware *Middleware
}

func NewRouter(router *httprouter.Router, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}
func NewRouterGin(router *gin.Engine, user *controller.UserHandler, middleware *Middleware) *GinRouter {
	return &GinRouter{
		router:     router,
		user:       user,
		middleware: middleware,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/users", r.user.GetUsers)
	r.router.POST("/users/register", r.user.Register)
	r.router.POST("/users/login", r.user.Login)

	log.Println("server running at port", port)
	http.ListenAndServe(port, r.router)
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	users := r.router.Group("/users")
	users.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.user.GinGetUsers, []string{"admin"}))
	users.POST("/register", r.user.GinRegister)
	users.POST("/login", r.user.GinLogin)

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
