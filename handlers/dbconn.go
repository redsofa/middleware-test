package handlers

import (
	"database/sql"
	"github.com/gorilla/context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redsofa/middleware-test/consts"
	"github.com/redsofa/middleware-test/logger"
	"net/http"
)

type DBConn struct {
	next http.Handler
}

func (h *DBConn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Setting up the database connection")
	_, ok := context.GetOk(r, consts.DB_KEY)

	if ok {
		errJson := `{"Error":"Internal Server Error"}`
		logger.Error.Println("DBConnector middleware error. DB_KEY already set.")
		http.Error(w, errJson, http.StatusInternalServerError)
		return
	}

	//https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
	db, err := sql.Open("sqlite3", "./middleware-test.db")

	if err != nil {
		errJson := `{"Error":"Internal Server Error"}`
		logger.Error.Println("DBConnector database connection problem. Check Logs.")
		http.Error(w, errJson, http.StatusInternalServerError)
		return
	}

	context.Set(r, consts.DB_KEY, db)

	//Close database connection once middleware chain is complete
	defer db.Close()

	if h.next != nil {
		h.next.ServeHTTP(w, r)
	}
}

func DBConnAdapter() Adapter {
	//The adapter type is a function that
	//takes an http.Handler interface and returns
	//an http.Handler inteface
	adapter := func(h http.Handler) http.Handler {
		//The DBConn type implements the
		//http.Handler interface
		return &DBConn{h} //h being the next handler to call ServeHTTP on
	}
	//Return the adapter
	return adapter
}
