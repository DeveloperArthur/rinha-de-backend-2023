package main

import (
	"golang-first-api-rest/database"
	"golang-first-api-rest/queue"
	"golang-first-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	go queue.Consumer()
	routes.HandleRequests()
}
