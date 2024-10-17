package results

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/chadsmith12/pacer/pkgs/pulse"
)

type TempResultWriter struct {
    component templ.Component
}

func (t TempResultWriter) Write(w http.ResponseWriter, req *http.Request) {
    t.component.Render(req.Context(), w)
}

func TemplResult(component templ.Component) pulse.PuleHttpWriter {
    return TempResultWriter{ component: component }
}
