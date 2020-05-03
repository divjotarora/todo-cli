package main

import (
	"fmt"

	"github.com/divjotarora/todo-cli/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(fmt.Sprintf("error initializing application: %v", err))
	}

	app.Start()
}
