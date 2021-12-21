package google

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	conf       *oauth2.Config
	ctx        context.Context
	cookiename = "be-router"
)

//User is representation of Google Authentication response
type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Hd            string `json:"hd"`
}

// setCookie is a function that set cookie with expiration time of 2 Hours to keep access to the service secured,
// but at the same time not to force user re-login very often
// it saves base64encoded email as value
func setCookie(w http.ResponseWriter, email string) {
	cookieValue := b64.StdEncoding.EncodeToString([]byte(email))
	expiration := time.Now().Add(2 * time.Hour)
	cookie := http.Cookie{Name: cookiename, Value: cookieValue, Expires: expiration, Path: "/", Secure: false, MaxAge: 90000}
	http.SetCookie(w, &cookie)
}

// AuthenticateUser redirect user to Google Authentication page
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	// this is config file with setting for Google OAuth2 application
	// see the setting at ls-marketing-dev
	conf = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:    google.Endpoint,
		RedirectURL: os.Getenv("OAUTH_REDIRECT_URL"),
	}
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	authCodeURL := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, authCodeURL, 302)
}

//CallBackHandler is a function that we call after successful authentication with Google
func CallBackHandler(w http.ResponseWriter, r *http.Request) {

	//get params from the URL
	queryParts, err := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect
	code := queryParts["code"][0]
	if code == "" {
		fmt.Println("Url Param 'code' is missing")
		io.WriteString(w, "Error: could not find 'code' URL parameter\n") //nolint:errcheck
		return
	}

	// Exchange will do the handshake to retrieve the initial access token.
	UserInfo, err := getUserInfo(code)
	if err != nil {
		log.Println(err)
	}
	log.Println("User " + UserInfo.Name + " authenticated")
	setCookie(w, UserInfo.Email)

}

// getUserInfo return user's information based on the scope specified in the AuthenticateUser
func getUserInfo(code string) (User, error) {

	token, err := getAccessToken(code)
	if err != nil {
		return User{}, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return User{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return User{}, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	var user User
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//getAccessToken is a function to get and rotate AccessToken
func getAccessToken(code string) (string, error) {

	token, err := conf.Exchange(ctx, code)
	src := conf.TokenSource(ctx, token)
	oauth2.ReuseTokenSource(token, src)
	a, err := src.Token()
	if err != nil {
		return "", err

	}
	accessToken := a.AccessToken

	if accessToken != "" {
		return accessToken, nil
	}

	return "", errors.New("could not get Access Token from getAccessToken")

}
