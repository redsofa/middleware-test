package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redsofa/middleware-test/consts"
	"github.com/redsofa/middleware-test/logger"
	"net/http"
)

type PlacesHandler struct {
}

func (h *PlacesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In PlacesHandler ... ")

	//Get the connection that was set in the dbconn middleware
	db, ok := context.GetOk(r, consts.DB_KEY)

	if !ok {
		errJson := `{"Error":"Internal Server Error"}`
		logger.Error.Println("DBConnector middleware error. DB_KEY not set")
		http.Error(w, errJson, http.StatusInternalServerError)
		return
	}

	//Query the database to get the places...
	rows, err := db.(*sql.DB).Query("select *  from place")

	if err != nil {
		errJson := `{"Error":"Internal Server Error"}`
		logger.Error.Println("DB query issue. ", err)
		http.Error(w, errJson, http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	type Place struct {
		Name string
	}

	type Places []Place

	type WResponse struct {
		Places
	}

	resp := WResponse{}

	//Stuff places inside a response object
	for rows.Next() {
		var name string
		err = rows.Scan(&name)

		if err != nil {
			errJson := `{"Error":"Internal Server Error"}`
			logger.Error.Println("DB query issue. ", err)
			http.Error(w, errJson, http.StatusInternalServerError)
			return
		}

		res := Place{
			Name: name,
		}

		resp.Places = append(resp.Places, res)
	}

	//Create JSON out of our response object
	mResp, err := json.Marshal(resp)
	if err != nil {
		msg := "Response Marshal() error."
		json := fmt.Sprintf(`{"Error":"%s"}`, msg)
		logger.Error.Println(msg, err)
		http.Error(w, json, http.StatusInternalServerError)
		return
	}

	//Create json with all results...
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(mResp))
}
