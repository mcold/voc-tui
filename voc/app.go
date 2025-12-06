package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type applicationType struct {
	pages *tview.Pages
}

var app *tview.Application

func (application *applicationType) init() {
	app = tview.NewApplication()

	application.pages = tview.NewPages()
	pageMain.build()

	application.registerGlobalShortcuts()

	if err := app.SetRoot(application.pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func (application *applicationType) registerGlobalShortcuts() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC:
			application.ConfirmQuit()
		default:
			return event
		}
		return nil
	})
}

func (application *applicationType) ConfirmQuit() {
	pageConfirm.show("Are you sure you want to exit?", application.Quit)
}

func (application *applicationType) Quit() {
	if zoteroDB.DB != nil {
		zoteroDB.DB.Close()
	}
	app.Stop()
}
