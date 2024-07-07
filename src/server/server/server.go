package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/baldurstod/go-shop/src/server/server/handlers"
)

func StartServer(c *Config) {
	initHandlers()

	fmt.Printf("StartServer %s", c.Address)

	/*helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)*/

	http.ListenAndServe(":12345", nil)
}

//go:embed html/*
var staticContent embed.FS

func initHandlers() {

	fsys := fs.FS(staticContent)
	html, _ := fs.Sub(fsys, "html")

	//http.Handle("/", new(handlers.StaticHandler));
	http.Handle("/", &handlers.RecoveryHandler{Handler: http.FileServer(http.FS(html))})
	http.Handle("/orders/", &handlers.RecoveryHandler{Handler: new(handlers.OrdersHandler)})
}
