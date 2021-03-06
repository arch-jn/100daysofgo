package main

import (
	"cloud-native-go/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
