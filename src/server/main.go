package main

import "github.com/baldurstod/go-shop/src/server/server"

func main() {
	config := server.Config{Address: "12345"}
	server.StartServer(&config)
	/*order := model.Order{}
	order.Init()*/
}
