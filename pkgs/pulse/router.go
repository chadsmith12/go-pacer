package pulse

import (
	"fmt"
	"net/http"
)

type PulseRouter struct {
    mux *http.ServeMux
    pulse *PulseApp
}

func NewRouter(p *PulseApp) *PulseRouter {
    return &PulseRouter{
        mux: http.NewServeMux(),
        pulse: p,
    }
}

func (router *PulseRouter) Get(pattern string, endpoint EndpointHandler, middleware ...MiddlewareFunc) {
    router.apply(GET, pattern, endpoint, middleware...)
}

func (router *PulseRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    router.mux.ServeHTTP(w, req)
}

func (router *PulseRouter) apply(method, pattern string, endpoint EndpointHandler, middleware ...MiddlewareFunc) {
    prefix := fmt.Sprintf("%s %s", method, pattern)
    router.mux.Handle(prefix, handle(endpoint))
}

func handle(endpoint EndpointHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        response := endpoint(r)
        response.Write(w, r)
    }
}
