package app

import (
	"net/http"

	authorsV1 "github.com/chadsmith12/pacer/internal/authors/v1"
	"github.com/chadsmith12/pacer/internal/components"
	"github.com/chadsmith12/pacer/internal/health"
	"github.com/chadsmith12/pacer/internal/results"
	"github.com/chadsmith12/pacer/pkgs/pulse"
)

func (a *App) loadEndpoints() {
    group := a.pulse.Group("/api")

    group.Get("/health", health.Health)

    authorsHandlers := authorsV1.NewHandlers(a.db, a.pulse.Logger())
    authorsHandlers.AuthorRoutes(group)

    a.pulse.Get("/", func(req *http.Request) pulse.PuleHttpWriter {
	helloComp := components.Hello("Chad")

	return results.TemplResult(helloComp)
    })
}
