package app

import (
	"go.uber.org/fx"
)

type App struct {
	fx *fx.App
}

func NewApp() *App {
	f := getFx()

	return &App{
		fx: f,
	}
}

func (a *App) Start() {
	a.fx.Run()
}
