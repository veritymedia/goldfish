package main

import (
	"context"

	"github.com/f1bonacc1/glippy"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) startClipboardWatcher() {
	ch := glippy.Watch(context.Background())

	// TODO: Show the window from which the text came from
	// https://github.com/jezek/xgb/blob/master/examples/get-active-window/main.go

	// list := []string{}
	for copiedText := range ch {
		a.Models.ClipboardItem.Set(copiedText)
		a.notifyClipboardUpdate()

	}

}

func (a *App) notifyClipboardUpdate() {
	runtime.EventsEmit(a.ctx, "clipboard-update")
}
