package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mtssrs/car_dealer_API/adapter/postgres"
	"github.com/mtssrs/car_dealer_API/di"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	carService := di.ConfigCarDi(conn)

	router := mux.NewRouter()
	router.Handle("/car", http.HandlerFunc(carService.Create)).Methods("POST")
	router.Handle("/car", http.HandlerFunc(carService.Fetch)).Queries(
		"page", "{page}",
		"descending", "{descending}",
		"sort", "{sort}",
		"itemsPerPage", "{itemsPerPage}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
