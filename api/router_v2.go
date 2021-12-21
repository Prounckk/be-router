package api

import (
	"be-router/api/middleware"
	controller "be-router/controller/parking"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//RegisterRoutesV2 manage v2 routes for the app with support of distributed by geolocation parking spots
// example of the path is: v2/canada/qc/montreal/p344
func RegisterRoutesV2(r *mux.Router) {

	//Health check
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
	})

	// register v1 of the api sub-router
	v2 := r.PathPrefix("/v2").Subrouter()
	v2.HandleFunc("/", v2Info)

	// register geolocation for the parking location
	geo := v2.PathPrefix("/{country}/{region}/{city}").Subrouter()
	// register parking sub-router
	p := geo.PathPrefix("/parkings").Subrouter()
	p.HandleFunc("/{number}", controller.GetParkingByName).Methods("GET")

	u := geo.PathPrefix("/users").Subrouter()
	u.Use(middleware.AuthMiddleware)
	u.HandleFunc("/{id}/parking/{number}", notImplemented).Methods("PUT")
	u.HandleFunc("/{id}/parking/{number}", controller.GetParkingStatusByName).Methods("GET")
	u.HandleFunc("/{id}/parking/{number}/timer", controller.GetParkingTimerByName).Methods("GET")

	a := geo.PathPrefix("/administrators").Subrouter()
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

	i := geo.PathPrefix("/inspectors").Subrouter()
	u.Use(middleware.AuthMiddleware)
	i.HandleFunc("/{id}/parking/{number}", controller.GetParkingStatusByName).Methods("GET")

}

func v2Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is v2 of the api with support of geolocation.\nDocumentation is available at https://amazinglink.com"))
}
