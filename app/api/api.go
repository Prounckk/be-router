package api

import (
	"be-router/api/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func StartApiServer() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	RegisterRoutes(router)

	http.ListenAndServe(port, router)
	fmt.Println("The app is running. http://localhost" + port + "/ping")
}
