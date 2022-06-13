package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gh "github.com/gorilla/handlers"
	"github.com/softilium/mb4/api"
	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/pages"
)

func gracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {

	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)

}

func initServer(listenAddr string) *http.Server {

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets/"))
	router.Handle("/assets/", http.StripPrefix("/assets", fs))

	router.HandleFunc("/", pages.Index)
	router.HandleFunc("/login", pages.Login)

	router.HandleFunc("/api/login", api.ApiLogin)
	router.HandleFunc("/api/register", api.ApiRegister)
	//router.HandleFunc("/auth/logout", api.Auth_Logout)
	//router.HandleFunc("/auth/refresh", api.Auth_RefreshToken)
	//router.HandleFunc("/auth/me", api.Auth_Me)
	//router.HandleFunc("/auth/reset-password", api.Auth_ResetPassword)
	//router.HandleFunc("/auth/update-user", api.Auth_UpdateUser)

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server.Handler = gh.RecoveryHandler()(server.Handler)
	server.Handler = gh.LoggingHandler(os.Stdout, server.Handler)

	return server

}

func main() {

	listenAddr := config.C.ListenAddr

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := initServer(listenAddr)
	go gracefullShutdown(server, quit, done)

	log.Println("Server is ready to handle requests at", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	log.Println("Server stopped")

}
