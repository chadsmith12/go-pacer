package app

import "github.com/chadsmith12/pacer/pkgs/pulse"

type App struct {
    pulse *pulse.PulseApp
}

func New() *App {
    pulseApp := pulse.Pulse(":4500")

    return &App{pulse: pulseApp}
}

func (a *App) Start() error {
    a.loadEndpoints()
    return a.pulse.Start()
}
