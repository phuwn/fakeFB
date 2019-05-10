package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/phuwn/fakeFB/pkg/routes"
	"github.com/phuwn/fakeFB/pkg/routes/mw"
)

var addr string = ":8080"

func main() {
	_, cleanup := Init()
	defer cleanup()
	r := routes.NewRoutes()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Printf("Transport HTTP,addr: %s", addr)
		errs <- http.ListenAndServe(addr, mw.WithCORS(r))
	}()

	log.Printf("exit %s", <-errs)
}
