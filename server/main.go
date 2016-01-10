package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/redsofa/middleware-test/configuration"
	"github.com/redsofa/middleware-test/consts"
	"github.com/redsofa/middleware-test/handlers"
	"github.com/redsofa/middleware-test/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//Setup the logger
	logger.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//Read configuration information from server_config.json
	config.ReadServerConf()

	//The port our server listens on
	listenPort := config.ServerConf.Port

	logger.Info.Printf("Sever Starting - Listing on port %d - (Version - %s)", listenPort, consts.SERVER_VERSION)

	//Create router
	router := mux.NewRouter()

	//Setup our routes
	router.Handle("/people",
		handlers.Execute(&handlers.PeopleHandler{},
			handlers.AuthHandlerAdapter(),
			handlers.DBConnAdapter())).Methods("GET")

	router.Handle("/places",
		handlers.Execute(&handlers.PlacesHandler{},
			handlers.AuthHandlerAdapter(),
			handlers.DBConnAdapter())).Methods("GET")

	//Listen for connections and serve content
	logger.Info.Println(http.ListenAndServe(":"+strconv.Itoa(listenPort),
		context.ClearHandler(logger.HttpLog(router))))
}
