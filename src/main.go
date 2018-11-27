package main

import (
	"./accountservice/service"
	"./accountservice/dbclient"
	"fmt"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoldClient()
	defer service.DBClient.Close()
	service.StartWebService("8088")
}

//creates instance and calls the OpenBoldDb and Seed funcs
// 启动时初始化一个BoltDb
func initializeBoldClient()  {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}