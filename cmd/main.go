package main

import (
	"golang-shop/initializers"
	"golang-shop/internal/router"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_shop_counter",
	Help: "Counting the total number of requets being handled",
})

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	db, err := initializers.ConnectDatabase()

	if err != nil {
		log.Println("Failed to connect database")
	}

	router := router.ApiRouter(db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
}
