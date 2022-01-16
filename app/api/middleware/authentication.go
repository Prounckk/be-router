package middleware

import (
	"be-router/controller/auth/google"
	controller "be-router/controller/user"
	b64 "encoding/base64"
	"fmt"
	"net/http"
)

var cookieName = "be-router"

//AuthMiddleware manages user authentication and authorization
// Ex: there are 3 types of users that has different level of access
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		theCookie, err := r.Cookie(cookieName)

		if err != nil {
			google.AuthenticateUser(w, r)
			return
		}

		email, err := b64.StdEncoding.DecodeString(theCookie.Value)
		if err != nil {
			fmt.Println(err)
			return
		}
		ok := controller.GetUserByEmail(string(email))
		if !ok {
			w.WriteHeader(404)
			fmt.Println("user is not authorized, sorry")
		}

		//TODO: implement logic for different type of users, please!
		next.ServeHTTP(w, r)
	})
}
