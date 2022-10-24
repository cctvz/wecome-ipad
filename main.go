package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wecome-ipad/router"
)

func main() {

	r := router.Router()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Printf("[api] start http server listening %s", "8080")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}

	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
