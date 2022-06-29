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
	router.HandleFunc("/profile", pages.Profile)
	router.HandleFunc("/invest-accounts", pages.InvestAccounts)
	router.HandleFunc("/invest-edit-account", pages.InvestAccount)

	router.HandleFunc("/api/users/login", api.UsersLogin)
	router.HandleFunc("/api/users/register", api.UsersRegister)
	router.HandleFunc("/api/users/logout", api.UsersLogout)
	router.HandleFunc("/api/users/start-invest-accounts-flow", api.UsersStartInvestAccountsFlow)
	router.HandleFunc("/api/invest-accounts", api.InvestAccounts)
	router.HandleFunc("/api/invest-account-valuations", api.InvestAccountValuations)
	router.HandleFunc("/api/industries", api.Industries)
	router.HandleFunc("/api/emitents", api.Emitents)
	router.HandleFunc("/api/tickers", api.Tickers)
	router.HandleFunc("/api/quotes", api.Quotes)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
	if config.C.Debug {
		server.ReadTimeout = 5 * time.Second
		server.WriteTimeout = 10 * time.Second
		server.IdleTimeout = 15 * time.Second
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

//TODO редактировать в тикерах параметры эмиссии (только админ)
//TODO Определять в квотах капитализацию и див.доходность+выплаты "на лету при загрузке квот"
//TODO Определять и записывать DSI по DivPayouts + mcap при загрузке из бэкэнд
//TODO Часовой пояс в настройках
