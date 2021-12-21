package api

import (
	"be-router/api/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApiServer() {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	RegisterRoutes(router)

	http.ListenAndServe(":8080", router)
	fmt.Println("The app is running. http://localhost:8080/ping")
}
