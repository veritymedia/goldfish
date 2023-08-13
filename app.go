package main

import (
	"context"
	"fmt"

	"github.com/veritymedia/goldfish/pkg/models"
	"github.com/veritymedia/goldfish/pkg/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Env struct {
	ctx    context.Context
	Models Models
}

func (e *Env) GetClipboardItems() ([]models.ClipboardItem, error) {
	models, err := e.Models.ClipboardItem.All()
	if err != nil {
		fmt.Println("ERROR: (*Env)GetClipboardItems: could not fetch from DB. ")
		return nil, err
	}
	fmt.Println("SUCCESS: (*Env)GetClipboardItems:  models fetched from DB.", models)
	return models, nil
}

func NewEnv(models Models) *Env {
	return &Env{Models: models}
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) domready(ctx context.Context) {
	// TODO: check if first startup, if yes redirect to onboarding route.
	isFirstLaunch := utils.CheckIsFirstLaunch()

	runtime.EventsEmit(a.ctx, "checkIsFirstLaunch", isFirstLaunch)

}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
// func (a *App) startup(ctx context.Context) {
// 	a.ctx = ctx
// 	a.startClipboardWatcher()

// }

// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }
