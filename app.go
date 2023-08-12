package main

import (
	"context"

	"github.com/veritymedia/goldfish/pkg/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Env struct {
	ctx    context.Context
	Models *Models
}

func NewEnv(models *Models) *Env {
	return &Env{Models: models}
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
	// return &App{Models: &env}

}

func (a *App) domready(ctx context.Context) {
	// TODO: check if first startup, if yes redirect to onboarding route.
	isFirstLaunch := utils.CheckIsFirstLaunch()

	runtime.EventsEmit(a.ctx, "checkIsFirstLaunch", isFirstLaunch)

}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.StartClipboardWatcher()

}

// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }
