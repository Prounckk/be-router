package api

import (
	"be-router/api/middleware"
	google "be-router/controller/auth/google"
	controller "be-router/controller/parking"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router) {
	//Health check
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
	})
	//Google OAuth
	r.HandleFunc("/oauth/callback", google.CallBackHandler)

	// register v1 of the api sub-router
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/", v1Info)
	// register parking sub-router
	p := v1.PathPrefix("/parkings").Subrouter()
	p.HandleFunc("/{number}", controller.GetParkingByName).Methods("GET")

	u := v1.PathPrefix("/users").Subrouter()
	u.Use(middleware.AuthMiddleware)
	u.HandleFunc("/{id}/parking/{number}", notImplemented).Methods("PUT")
	u.HandleFunc("/{id}/parking/{number}", controller.GetParkingStatusByName).Methods("GET")
	u.HandleFunc("/{id}/parking/{number}/timer", controller.GetParkingTimerByName).Methods("GET")

	a := v1.PathPrefix("/administrators").Subrouter()
	u.Use(middleware.AuthMiddleware)
	a.HandleFunc("/{id}/parking/{number}", controller.GetParkingStatusByName).Methods("GET")
	a.HandleFunc("/{id}/parking/{number}", notImplemented).Methods("POST")
	a.HandleFunc("/{id}/parking/{number}", notImplemented).Methods("PUT")
	a.HandleFunc("/{id}/parking/{number}", notImplemented).Methods("DELETE")
	a.HandleFunc("/{id}/parkingGroups", notImplemented).Methods("GET")
	a.HandleFunc("/{id}/parkingGroups/{letter}", notImplemented).Methods("GET")
	a.HandleFunc("/{id}/parkingGroups/{letter}", notImplemented).Methods("POST")
	a.HandleFunc("/{id}/parkingGroups/{letter}", notImplemented).Methods("PUT")
	a.HandleFunc("/{id}/parkingGroups/{letter}", notImplemented).Methods("DELETE")

	i := v1.PathPrefix("/inspectors").Subrouter()
	u.Use(middleware.AuthMiddleware)
	i.HandleFunc("/{id}/parking/{number}", controller.GetParkingStatusByName).Methods("GET")

}

func v1Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is v1 of the api. Documentation is available at https://amazinglink.com"))
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented, stay tuned!"))
})
