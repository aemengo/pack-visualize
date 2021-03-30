package main

import (
	"fmt"
	"github.com/aemengo/pack-visualize/component"
	_ "github.com/fatih/color"
	"github.com/rivo/tview"
	"os"
)

type Screen interface {
	Run() chan bool
}

func main() {
	app := tview.NewApplication()
	loading := component.NewLoading(app)

	go func() {
		<-loading.Run()

		app.Stop()
	}()

	err := app.Run()
	expectNoError(err)
}

func expectNoError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
