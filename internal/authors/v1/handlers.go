package v1

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/chadsmith12/pacer/internal/results"
	"github.com/chadsmith12/pacer/pkgs/db"
	"github.com/chadsmith12/pacer/pkgs/pulse"
	"github.com/jackc/pgx/v5/pgxpool"
)

type handlers struct {
    repo *db.Queries
    logger *slog.Logger
}

func NewHandlers(pool *pgxpool.Pool, logger *slog.Logger) *handlers {
    return &handlers{repo: db.New(pool), logger: logger}
}

func (h *handlers) AuthorRoutes(group *pulse.Group) {
    group.Get("/authors", h.ListAuthors)
    group.Post("/authors", h.CreateAuthor)
}

func (h *handlers) ListAuthors(req *http.Request) pulse.PuleHttpWriter {
	authors, err := h.repo.ListAuthors(context.Background())
	if err != nil {
	    return pulse.InternalErrorJson(err)
	}
    
    h.logger.LogAttrs(context.TODO(), slog.LevelInfo, "listed authors")

	return results.List(authors)
}

func (h *handlers) CreateAuthor(req *http.Request) pulse.PuleHttpWriter {
	var author db.CreateAuthorParams
	err := pulse.Json(req.Body, &author)

	if err != nil {
	    return pulse.InternalErrorJson(err)
	}

	created, err := h.repo.CreateAuthor(context.Background(), author)
	if err != nil {
	    return pulse.InternalErrorJson(err)
	}

	return pulse.JsonResult(created)
}

