package main

import (
	"./accountservice/dbclient"
	"./accountservice/service"
	"fmt"
)

//https://segmentfault.com/blog/microgo
//https://segmentfault.com/a/1190000015135650
var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoldClient()
	defer service.DBClient.Close()
	service.StartWebService("8088")
}

//creates instance and calls the OpenBoldDb and Seed funcs
// 启动时初始化一个BoltDb
func initializeBoldClient() {
	// service.BDClient 为dbclient 定义的接口类型实例；struct 类型的BoltClient实现了上述接口，dbclient.BoltClient{}是struct类型的实例
	// 接口类型的实例可以赋值为不同实现该接口的类型值，只要此类型实现接口即可
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
