package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DitoAdriel99/go-oauth2/config"
)

func CallBack(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query()["state"][0]
	if state != "randomState" {
		fmt.Fprintln(w, "State dont match")
		return
	}

	code := r.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(w, "Code-Token exchange failed")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(w, "User data fetch failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(w, "Failed to read response body:", err)
			return
		}
		fmt.Println("Error response body:", string(body))
		fmt.Fprintln(w, "User data fetch failed. Status code:", resp.StatusCode)
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(w, "Failed to read user data:", err)
		return
	}

	fmt.Fprintln(w, string(userData))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomState")

	http.Redirect(w, r, url, http.StatusSeeOther)
}
