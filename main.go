package main

import (
	"cloudcost-demo/dao"
	"cloudcost-demo/router"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	routersInit := router.SetupRouter()

	// srv := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        routersInit,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	routersInit.Run()

}
