package main

import (
	"context"

	"github.com/f1bonacc1/glippy"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ClipboardMessage struct {
	Status  string `json:"status"`
	Text    string `json:"text"`
	Message string `json:"message"`
}

func (a *App) startClipboardWatcher() {
	ch := glippy.Watch(context.Background())

	// TODO: Show the window from which the text came from
	// https://github.com/jezek/xgb/blob/master/examples/get-active-window/main.go

	// list := []string{}
	for data := range ch {
		a.Models.ClipboardItem.Set(data)
		a.notifyClipboardUpdate()

		// if data != "" {
		// 	a.EmmitClipboardText(&ClipboardMessage{Status: "ok", Text: data, Message: ""})
		// } else {
		// 	a.EmmitClipboardText(&ClipboardMessage{Status: "not-text", Message: "You copied something, but it wasnt text."})
		// }

	}

}

func (a *App) notifyClipboardUpdate() {
	runtime.EventsEmit(a.ctx, "clipboard-update")
}
