package main

import (
	"golang-batch6-group3/server"
	"golang-batch6-group3/server/config"
	"net/http"
)

func main() {
	run()
}

func run() {
	router := http.NewServeMux()
	port := ":6000"

	db := config.ConnectDB()

	server.StartServer(router, port, db)

}
