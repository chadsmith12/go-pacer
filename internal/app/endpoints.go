package app

import (
	"context"
	"net/http"

	"github.com/chadsmith12/pacer/internal/health"
	"github.com/chadsmith12/pacer/internal/results"
	"github.com/chadsmith12/pacer/pkgs/db"
	"github.com/chadsmith12/pacer/pkgs/pulse"
)

func (a *App) loadEndpoints() {
    group := a.pulse.Group("/api")

    group.Get("/health", health.Health)

    group.Get("/authors", func(req *http.Request) pulse.PuleHttpWriter {
	repo := db.New(a.db)
   
	authors, err := repo.ListAuthors(context.Background())
	if err != nil {
	    return pulse.InternalErrorJson(err)
	}

	return results.List(authors)
    })

    group.Post("/authors", func(req *http.Request) pulse.PuleHttpWriter {
	var author db.CreateAuthorParams
	err := pulse.Json(req.Body, &author)

	if err != nil {
	    return pulse.InternalErrorJson(err)
	}

	repo := db.New(a.db)
	created, err := repo.CreateAuthor(context.Background(), author)
	if err != nil {
	    return pulse.InternalErrorJson(err)
	}

	return pulse.JsonResult(created)
    })
}
