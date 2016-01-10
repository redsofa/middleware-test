package handlers

import (
	"github.com/redsofa/middleware-test/configuration"
	"github.com/redsofa/middleware-test/logger"
	"net/http"
)

type AuthHandler struct {
	next http.Handler
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In AuthHandler ... ")

	logger.Info.Println("Authentication check....")

	//An authentication handler stub. authSuccess could be set by validating
	//a token sent in request header. User sends a valid token, authSuccess is set to true

	//Right now it's set in a config file.
	//Try setting AuthPass to false in the server_config.json file and restarting the server
	//to see the authentication faliure
	authSuccess := config.ServerConf.AuthPass

	if !authSuccess {
		errJson := `{"Error":"Access Denied"}`
		logger.Error.Println("Access Denied")
		http.Error(w, errJson, http.StatusForbidden)
		return
	}

	logger.Info.Println("Authentication check passed.")

	if h.next != nil {
		h.next.ServeHTTP(w, r)
	}
}

func AuthHandlerAdapter() Adapter {
	//The adapter type is a function that
	//takes an http.Handler interface and returns
	//an http.Handler inteface
	adapter := func(h http.Handler) http.Handler {
		//The AuthHandler type implements the
		//http.Handler interface
		return &AuthHandler{next: h}
	}
	//Return the adapter
	return adapter
}
