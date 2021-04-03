package main

import (
	"fmt"
	"github.com/aemengo/pack-visualize/component"
	"github.com/rivo/tview"
	"os"
	"time"
)

type Screen interface {
	Run() chan bool
}

func main() {
	var (
		app       = tview.NewApplication()
		loading   = component.NewLoading(app)
		dashBoard = component.NewDashboard(app)
	)

	go func() {
		<-loading.Run()

		time.Sleep(time.Second)
		<-dashBoard.Run()

		app.Stop()
	}()

	time.Sleep(time.Second)
	err := app.Run()
	expectNoError(err)
}

func expectNoError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
