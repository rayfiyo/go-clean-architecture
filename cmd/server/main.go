package main

import (
	"app/internal/adapter/http"
	"app/internal/platform"
	"fmt"
	"log"
)

func main() {
	app, err := platform.Build()
	if err != nil {
		log.Fatalf("bootstrap error: %v", err)
	}
	r := http.NewRouter(app.Handler)
	addr := ":" + app.Config.Port
	fmt.Printf("listening on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
