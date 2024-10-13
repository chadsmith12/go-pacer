package health

import (
	"net/http"

	"github.com/chadsmith12/pacer/pkgs/pulse"
)

func Health(req *http.Request) pulse.PuleHttpWriter {
    return pulse.OkResult()
}
