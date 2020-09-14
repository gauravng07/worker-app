package main

import (
	"context"
	"hello-fresh-app/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var ctx context.Context
const port = "9000"

func main()  {
	server := &http.Server{
		Addr: ":" + port,
		Handler: createRouter(),
	}
	start(server)
}

func createRouter() *mux.Router  {
	r := mux.NewRouter()
	internal.Configure(r)
	return r
}

func start(server *http.Server)  {
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	graceFulStop(server)
}

func graceFulStop(server *http.Server)  {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<- stop
	log.Fatal("Shutting server down...")
	server.Shutdown(ctx)
}