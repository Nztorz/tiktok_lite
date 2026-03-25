package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nztorz/tiktok_lite/internal/database"
	_ "github.com/lib/pq"
)

type Application struct {
	Config Config
	DB     *database.Queries
	Logger *log.Logger
}

func NewApplication(cfg Config) (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dbURL)
	if err != nil {
		return nil, err
	}

	dbQueries := database.New(db)

	return &Application{
		Config: cfg,
		DB:     dbQueries,
		Logger: logger,
	}, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
