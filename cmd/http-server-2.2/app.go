package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type App struct {
	server *http.Server
	ctx    context.Context
	wg     sync.WaitGroup
}

func NewApp() *App {
	return &App{
		ctx: context.Background(),
	}
}

func (a *App) Initialize() error {
	a.ctx = context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/healthz", healthz)

	a.server = &http.Server{
		Addr:    Listener,
		Handler: mux,
	}
	return nil
}

func (a *App) RunLoop() {
	log.Println("http server startup...")
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		err := a.server.ListenAndServe()
		if err == http.ErrServerClosed {
			log.Println("http server got shutdown signal")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (a *App) Destroy() {
	log.Println("http server going to shutdown")
	a.server.Shutdown(a.ctx)
	a.wg.Wait()
	log.Println("http server has been shutdown")
}
