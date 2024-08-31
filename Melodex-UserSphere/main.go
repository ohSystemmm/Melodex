package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetText("Hello World!").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	flex := tview.NewFlex().
		AddItem(textView, 0, 1, true)

	app.SetRoot(flex, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
