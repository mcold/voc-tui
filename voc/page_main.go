package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type pageMainType struct {
	*tview.Pages
	*tview.Flex
}

var pageMain pageMainType

func (pageMain *pageMainType) build() {

	connectDB()
	pageMain.Pages = tview.NewPages()

	pageMain.Pages.SetBackgroundColor(tcell.ColorBlack)

	pageVoc.build()
	pageVoc.show()

	pageMain.Flex = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(pageMain.Pages, 0, 1, true)

	pageMain.Flex.SetBackgroundColor(tcell.ColorBlack)

	app.SetFocus(pageVoc.lVoc)

	application.pages.AddPage("main", pageMain.Flex, true, true)
}

func connectDB() {
	err := zoteroDB.Connect()
	check(err)
}
