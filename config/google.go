package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
