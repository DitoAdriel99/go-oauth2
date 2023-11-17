package router

import (
	"net/http"

	"github.com/DitoAdriel99/go-oauth2/handler"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/google/login", handler.Login)
	router.HandleFunc("/api/sessions/oauth/google", handler.CallBack)
	return router
}
