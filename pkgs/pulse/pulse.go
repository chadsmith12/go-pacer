package pulse

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
    GET = http.MethodGet
    HEAD = http.MethodHead
    POST = http.MethodPost
    PUT = http.MethodPut
    DELETE = http.MethodDelete
    CONNECT = http.MethodConnect
    OPTIONS = http.MethodOptions
    PATCH = http.MethodPatch
)

type EndpointHandler = func(req *http.Request) PuleHttpWriter
type MiddlewareFunc = func(EndpointHandler) EndpointHandler

type PulseApp struct {
    server *http.Server
    router *PulseRouter
    logger *slog.Logger
    addr string
}

func Pulse(addr string) *PulseApp {
    pulseApp := &PulseApp{ addr: addr }
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    router := NewRouter(pulseApp)
    server := &http.Server {
        Addr: addr,
        Handler: router,
    }

    pulseApp.logger = logger
    pulseApp.server = server
    pulseApp.router = router

    return pulseApp
}

func (p *PulseApp) Start() error {
    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    done := make(chan struct{})
    go func() {
        err := p.server.ListenAndServe()
        if err != nil && !errors.Is(err, http.ErrServerClosed) {
            p.logger.Error("Failed to start server and listen", slog.Any("error", err))
        }
        close(done)
    }()

    p.logger.Info("server listening on ", slog.String("addr", p.addr))

    select {
    case <-done:
        break;
    case <-ctx.Done():
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
        p.server.Shutdown(ctx)
        cancel()
    }

    return nil
}

func (p *PulseApp) Get(pattern string, endpoint EndpointHandler) {
    p.router.Get(pattern, endpoint)
}
