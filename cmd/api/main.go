package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Nztorz/tiktok_lite/internal/app"
	"github.com/Nztorz/tiktok_lite/internal/routes"
)

func main() {
	cfg := app.LoadConfig()

	app, err := app.NewApplication(cfg)
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)

	fmt.Printf("port number: %s", cfg.PORT)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.PORT),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port %s", cfg.PORT)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
