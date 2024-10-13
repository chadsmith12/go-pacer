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

type EndpointHandler func(req *http.Request) PuleHttpWriter
type MiddlewareFunc func(EndpointHandler) EndpointHandler
type Option func(*PulseApp)

type PulseApp struct {
    server *http.Server
    router *PulseRouter
    logger *slog.Logger
    addr string
}

func Pulse(options ...Option) *PulseApp {
    pulseApp := &PulseApp{ addr: ":4500" }
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    router := NewRouter(pulseApp)
    server := &http.Server { Handler: router,}

    pulseApp.logger = logger
    pulseApp.server = server
    pulseApp.router = router

    for _, option := range options {
        option(pulseApp)
    }
    pulseApp.server.Addr = pulseApp.addr

    return pulseApp
}

func WithAddr(addr string) Option {
    return func(pa *PulseApp) {
        pa.addr = addr
    }
}

func WithLogger(logger *slog.Logger) Option {
    return func(pa *PulseApp) {
        pa.logger = logger
    }
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

func (p *PulseApp) Post(pattern string, endpoint EndpointHandler) {
    p.router.Post(pattern, endpoint)
}

func (p *PulseApp) Group(prefix string) *Group {
    return p.router.Group(prefix)
}
