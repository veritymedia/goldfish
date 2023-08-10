package main

import (
	"context"
	"fmt"

	"github.com/f1bonacc1/glippy"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ClipboardMessage struct {
	Status  string `json:"status"`
	Text    string `json:"text"`
	Message string `json:"message"`
}

func (a *App) StartClipboardWatcher() {
	ch := glippy.Watch(context.Background())

	// list := []string{}
	for data := range ch {
		fmt.Println(data)
		if data != "" {
			a.EmmitClipboardText(&ClipboardMessage{Status: "ok", Text: data, Message: ""})
		} else {
			a.EmmitClipboardText(&ClipboardMessage{Status: "not-text", Message: "You copied something, but it wasnt text."})
		}

	}

}

func (a *App) EmmitClipboardText(message *ClipboardMessage) {
	runtime.EventsEmit(a.ctx, "newTextData", *message)
}
