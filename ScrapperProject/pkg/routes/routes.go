// Package classification of Product API
//
// Documentation for twitter scrapper api
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package routes

import (
	"log"
	"net/http"

	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/controllers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	// scrapper routes
	router.HandleFunc("/twitter", controllers.GetScrapedData).Methods("POST", "OPTIONS")
	// authentication Routes
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	// Facebook authentication Routes
	router.HandleFunc("/facebook/login", controllers.InitFacebookLogin)
	router.HandleFunc("/facebook/callback", controllers.HandleFacebookLogin)
	//google authentication routes
	router.HandleFunc("/google/login", controllers.GoogleLogin)
	router.HandleFunc("/google/callback", controllers.GoogleCallBack)

	ops := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(ops, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	log.Println("server started at port 8000")

}
