package main

import (
	"log"
	"net/http"

	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/routes"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8000", handler))

}
