package main

import (
	"fmt"
	"log"
	"net/http"
)

const webServerPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting server on port %s\n", webServerPort)
	// กำหนดค่าให้กับ http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webServerPort),
		Handler: app.routes(),
	}

	// Start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
