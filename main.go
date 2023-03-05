package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"

	"github.com/flosch/pongo2/v6"
	gh "github.com/gorilla/handlers"
	"github.com/softilium/mb4/api"
	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/pages"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	router.HandleFunc("/ticker", pages.Ticker)
	router.HandleFunc("/industry", pages.Industry)
	router.HandleFunc("/strategies", pages.Strategies)
	router.HandleFunc("/strategy", pages.Strategy)
	router.HandleFunc("/report", pages.Report)

	router.HandleFunc("/api/users/login", api.UsersLogin)
	router.HandleFunc("/api/users/register", api.UsersRegister)
	router.HandleFunc("/api/users/logout", api.UsersLogout)
	router.HandleFunc("/api/users/start-invest-accounts-flow", api.UsersStartInvestAccountsFlow)
	router.HandleFunc("/api/invest-accounts", api.InvestAccounts)
	router.HandleFunc("/api/industries", api.Industries)
	router.HandleFunc("/api/emitents", api.Emitents)
	router.HandleFunc("/api/tickers", api.Tickers)
	router.HandleFunc("/api/quotes", api.Quotes)
	router.HandleFunc("/api/cube/reload", api.Cube)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
	if !config.C.Debug {
		server.ReadTimeout = 5 * time.Second
		server.WriteTimeout = 10 * time.Second
		server.IdleTimeout = 15 * time.Second
	}

	if !config.C.Debug {
		server.Handler = gh.RecoveryHandler()(server.Handler)
	}
	server.Handler = gh.LoggingHandler(os.Stdout, server.Handler)

	return server

}

var localFormatter = message.NewPrinter(language.Russian)

func filterCur0(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r := localFormatter.Sprintf("%.0f\n", in.Float())
	return pongo2.AsSafeValue(r), nil
}

func filterCur1(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r := localFormatter.Sprintf("%.1f\n", in.Float())
	return pongo2.AsSafeValue(r), nil
}

func main() {

	err := pongo2.RegisterFilter("cur0", filterCur0)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = pongo2.RegisterFilter("cur1", filterCur1)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if config.C.Debug {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err.Error())
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	listenAddr := config.C.ListenAddr

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	log.Println("Loading market cube...")

	clRoutine := func() {
		err := cube.Market.LoadCube()
		if err != nil {
			log.Fatalf("Could not load market cube: %v\n", err.Error())
		}
		log.Println("Market cube loaded.")
	}
	go clRoutine()

	server := initServer(listenAddr)
	go gracefullShutdown(server, quit, done)

	log.Println("Server is ready to handle requests at", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v\n", listenAddr, err.Error())
	}

	<-done

	if config.C.Debug {
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal(err.Error())
		}
		_ = pprof.WriteHeapProfile(f)
		f.Close()
	}

	log.Println("Server stopped")

}
