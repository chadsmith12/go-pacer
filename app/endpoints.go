package app

import (
	authorsV1 "github.com/chadsmith12/pacer/app/authors/v1"
	"github.com/chadsmith12/pacer/app/health"
)

func (a *App) loadEndpoints() {
    a.pulse.UseStaticFiles()

    group := a.pulse.Group("/api")

    group.Get("/health", health.Health)

    authorsHandlers := authorsV1.NewHandlers(a.db, a.pulse.Logger())
    authorsHandlers.AuthorRoutes(group)
}
