package pulse_test

import (
	"net"
	"testing"
	"time"

	"github.com/chadsmith12/pacer/pkgs/pulse"
	"github.com/stretchr/testify/assert"
)

func TestPulseStart(t *testing.T) {
    app := pulse.Pulse(":4500")
  
    go func() {
        err := app.Start()
        if err != nil {
            t.Errorf("Server failed to start: %v", err)
        }
    }()
    
    // wait a little bit for server
    time.Sleep(time.Millisecond * 400)

    conn, err := net.Dial("tcp", ":4500")
    assert.NoErrorf(t, err, "error connecting to :4500: %v", err)
    conn.Close()
}
