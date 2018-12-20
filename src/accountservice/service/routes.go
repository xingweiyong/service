package service

import "net/http"

//Define a route type,a single route includes readable name,HTTP method,pattern
// and the function that will execute when the route is called.

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		//func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type","application/json;charset=UTF-8")
		//	w.Write([]byte("{\"result\":\"OK\"}"))
		//},
		GetAccount,
	},
	Route{
		"HealCheck",
		"GET",
		"/health",
		HealCheck,
	},
}
