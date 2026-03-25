package routes

import (
	"net/http"

	"github.com/Nztorz/tiktok_lite/internal/app"
)

func SetupRoutes(app *app.Application) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", app.HealthCheck)
	return mux
}
