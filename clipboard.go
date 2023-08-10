package main

import (
	"golang.design/x/clipboard"
)

func (a *App) InitClipboard() string {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	return "From backend"
}
