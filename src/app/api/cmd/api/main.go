package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"app/api/component"
	"app/api/middleware"
	"app/api/pkg/logger"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	port := "8081"

	// Create the logger.
	//l := logger.New(log.New(os.Stderr, "", log.LstdFlags))
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	core := component.Services(l)

	// Load the routes.
	component.LoadRoutes(core)

	l.Printf("Server started.")
	err := http.ListenAndServe(":"+port, middleware.Wrap(core.Router, l, core.Token.Secret()))
	if err != nil {
		l.Printf(err.Error())
	}
}