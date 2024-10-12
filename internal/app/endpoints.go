package app

import "github.com/chadsmith12/pacer/internal/health"

func (a *App) loadEndpoints() {
    group := a.pulse.Group("/api")

    group.Get("/health", health.Health)
}
