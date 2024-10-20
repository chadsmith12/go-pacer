package app

import (
	"context"
	"log"

	"github.com/chadsmith12/pacer/pkgs/db"
	"github.com/chadsmith12/pacer/pkgs/pulse"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
    pulse *pulse.PulseApp
    db *pgxpool.Pool
}

func New() *App {
    pulseApp := pulse.Pulse()

    return &App{pulse: pulseApp}
}

func (a *App) Start() error {
    pool, err := db.Connect(context.Background())
    if err != nil {
	log.Fatalf("failed to start server because of database pool failure: %v", err)
    }
    a.db = pool

    a.loadEndpoints()
    return a.pulse.Start()
}
