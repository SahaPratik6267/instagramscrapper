package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/utils"
	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/scrapper"
)

type ProfileName struct {
	InputProfile string `json:"userName"`
	SessionToken string `json:"token"`
}

//swagger:route POST /twitter twitterProfile
// Returns acount info (follower,folloing,likes,postcount)
// responses:
// 200: ProfileResponse
func GetScrapedData(w http.ResponseWriter, r *http.Request) {
	log.Println("data scrapper activated")
	utils.EnableCors(&w)
	//check if session is valid.

	var pro ProfileName
	err := json.NewDecoder(r.Body).Decode(&pro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !checkSession(w, pro.SessionToken) {
		w.Write([]byte("unauthorized access"))
		return
	}
	log.Println(pro.SessionToken)
	log.Println(pro.InputProfile)
	acc, err := scrapper.New().GetProfile(pro.InputProfile)
	if err != nil {
		log.Println("profile can not be found")
	}

	res, err := json.Marshal(acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Write(res)

}

func checkSession(w http.ResponseWriter, sessionToken string) bool {
	// We can obtain the session token from the requests cookies, which come with every request

	// We then get the session from our session map
	userSession, exists := sessions[sessionToken]
	if !exists {
		log.Println("error2")
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		log.Println("error 3")
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true

}
