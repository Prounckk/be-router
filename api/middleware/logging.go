package middleware

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// LoggingMiddleware use it for save requests to the logging system
// for now it just prints logs to the stdout but NewRelic, DataDog or others tools might be implemented here
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id != "" {
			// todo: log type of user, by parsing url and getting /user/ or /administrator / or etc
			log.Println("Special route with id " + id)
		}
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
