package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/config"
	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/models"
	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var userInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"

type userResponse struct {
	models.User
	Token string `json:"token"`
}
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type googleuserData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

const googlestate string = "googlestate"

//swagger:route POST /login loginparam
// Returns acount info (follower,folloing,likes,postcount)
// responses:
// 200: IdResponse
func LoginUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "err.Error()", http.StatusBadRequest)
		return
	}

	userFromDB, _ := models.GetUserByEmail(creds.Email)
	if userFromDB == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))
	}
	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if userFromDB.Password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("password did not match"))
	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(10 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		email:  creds.Email,
		expiry: expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	log.Println("user logged in")

	var userRes userResponse

	userRes.Token = sessionToken
	userRes.Name = userFromDB.Name

	userinjson, _ := json.Marshal(userRes)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(userinjson))
}

//swagger:route POST /register registerparam
// Returns acount info (follower,folloing,likes,postcount)
// responses:
// 200: IdResponse
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := newUser.CreateUser()
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}

//refresh token funtion
func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[newSessionToken] = session{
		email:  userSession.email,
		expiry: expiresAt,
	}

	// Delete the older session token
	delete(sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

//check authentication status from google
func authenticated(token string) (bool, error) {
	// validates the token by sending a request at userInfoURL
	bearToken := "Bearer " + token
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}

	req.Header.Add("Authorization", bearToken)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("fail to get response: %v", err)
	}
	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

// InitFacebookLogin function will initiate the Facebook Login
func InitFacebookLogin(response http.ResponseWriter, request *http.Request) {
	var OAuth2Config = config.GetFacebookOAuthConfig()
	url := OAuth2Config.AuthCodeURL(googlestate)
	http.Redirect(response, request, url, http.StatusTemporaryRedirect)
}

// HandleFacebookLogin function will handle the Facebook Login Callback
func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	var state = r.FormValue("state")
	var code = r.FormValue("code")

	if state != googlestate {
		http.Redirect(w, r, "/?invalidlogin=true", http.StatusTemporaryRedirect)
	}

	var OAuth2Config = config.GetFacebookOAuthConfig()

	token, err := OAuth2Config.Exchange(oauth2.NoContext, code)

	if err != nil || token == nil {
		http.Redirect(w, r, "/?invalidlogin=true", http.StatusTemporaryRedirect)
	}

	fbUserDetails, fbUserDetailsError := GetUserInfoFromFacebook(token.AccessToken)

	if fbUserDetailsError != nil {
		http.Redirect(w, r, "/?invalidlogin=true", http.StatusTemporaryRedirect)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fbUserDetails.Name))

}

// GoogleLogin function will initiate the google Login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	googleconfig := config.GetgoogleOAuthConfig()
	url := googleconfig.AuthCodeURL(googlestate)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Googlecallback function will initiate the google Login
func GoogleCallBack(w http.ResponseWriter, r *http.Request) {
	log.Println("call back came from google")
	state := r.FormValue("state")
	if state != googlestate {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("state dont match"))
		return
	}
	log.Println("State matched")
	code, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("decoding error" + err.Error()))
		return
	}
	log.Println("parse code from url")

	token, err := config.GetgoogleOAuthConfig().Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("code exchange failed" + err.Error() + code.Query().Get("code")))
		return
	}
	log.Println("exchange code with token : " + token.AccessToken)

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User Data Fetch Failed"))
		return
	}

	log.Println("get user details from accesstoken")

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User Data Fetch Failed"))
		return
	}
	var userData models.User
	data := make(map[string]interface{})
	err = json.Unmarshal(content, &data)
	if err != nil {
		w.Write([]byte("json parse error"))
		return
	}

	if name, ok := data["name"].(string); ok {
		userData.Name = name
	}
	if email, ok := data["email"].(string); ok {
		userData.Email = email
	}

	if userfound, _ := models.GetUserByEmail(userData.Email); userfound == nil {
		userData.CreateUser()
	}

	log.Println("parsed userData into user struct")

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		email:  userData.Email,
		expiry: expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userData.Name))
	return
}
