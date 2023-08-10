package main

import (
	"golang.design/x/clipboard"
)

func InitClipboard() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}
