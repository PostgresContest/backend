package app

import (
	"backend/internal/config"
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
	cfgReader := config.NewReader()
	err := cfgReader.Read()
	if err != nil {
		panic(err)
	}

}
