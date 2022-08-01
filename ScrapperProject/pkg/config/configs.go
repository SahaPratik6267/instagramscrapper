package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// Google OpenID client
)

var (
	dsn string = "host=host.docker.internal user=postgres password=postgres dbname=scraper port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db  *gorm.DB
)

func Connect() {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}

//configuration for facebook login
func GetFacebookOAuthConfig() *oauth2.Config {

	conf := &oauth2.Config{
		ClientID:     "3427411404146740",
		ClientSecret: "7d84efaa385f34e032d31ae5e6b84607",
		RedirectURL:  "http://localhost:8000/facebook/callback",
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"email"},
	}

	return conf

}

//configuration for google login
func GetgoogleOAuthConfig() *oauth2.Config {

	conf := &oauth2.Config{
		ClientID:     "237749859087-t14qs942vt4ucd0mis6e60nagkf1ussu.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8000/google/callback",
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	return conf

}
