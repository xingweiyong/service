package service

import (
	"../dbclient"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter,r *http.Request)  {
	//Read the accountId path parameter from mux map
	var accountId = mux.Vars(r)["accountId"]

	//Read the accouont struct BoitDB
	account,err := DBClient.QueryAccount(accountId)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		return
	}

	//If found marshal into JSON,write headers and content
	data,_ :=json.Marshal(account)
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Content-Length",strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}