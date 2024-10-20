package app

import (
	"net/http"

	authorsV1 "github.com/chadsmith12/pacer/app/authors/v1"
	"github.com/chadsmith12/pacer/app/health"
	"github.com/chadsmith12/pacer/app/results"
	"github.com/chadsmith12/pacer/app/views"
	"github.com/chadsmith12/pacer/pkgs/pulse"
)

func (a *App) loadEndpoints() {
    a.pulse.UseStaticFiles()

    a.pulse.Get("/home", func(req *http.Request) pulse.PuleHttpWriter {
	helloComp := views.Hello("Chad")

	return results.TemplResult(helloComp)
    })


    group := a.pulse.Group("/api")

    group.Get("/health", health.Health)

    authorsHandlers := authorsV1.NewHandlers(a.db, a.pulse.Logger())
    authorsHandlers.AuthorRoutes(group)
}
