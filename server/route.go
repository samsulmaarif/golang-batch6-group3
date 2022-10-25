package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

func StartServer(router *http.ServeMux, port string, db *sql.DB) {
	buildRoute(router, db)
	fmt.Println("Server running at ", port)

	http.ListenAndServe(port, router)
}

func buildRoute(router *http.ServeMux, db *sql.DB) {

}
