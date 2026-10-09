package main

import (
	"golang-auth/internal/httpserver"
	"golang-auth/internal/app"

	"time"
	"net/http"
	"log"
)

func main(){
	ctx := context.Background()
	
	a,err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Failed to close app: %v", err)
		}
		log.Println("App closed successfully")
	}
	router := httpserver.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}