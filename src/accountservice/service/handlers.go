package service

import (
	"../dbclient"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	//Read the accountId path parameter from mux map
	var accountId = mux.Vars(r)["accountId"]

	//Read the accouont struct BoitDB
	account, err := DBClient.QueryAccount(accountId)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	//If found marshal into JSON,write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//check health

type healthCheckResponse struct {
	Status string `json:"status"`
}

func writeJsonRespnse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

func HealCheck(w http.ResponseWriter, r *http.Request) {
	dbUP := DBClient.Check()
	if dbUP {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJsonRespnse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "DATABASE UNACCESSABLE"})
		writeJsonRespnse(w, http.StatusOK, data)
	}
}
