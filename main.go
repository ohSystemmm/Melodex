package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Username", "", 20, nil, nil).
		AddInputField("Client ID", "", 20, nil, nil).
		AddInputField("Client Secret", "", 20, nil, nil).
		AddInputField("Redirect URI", "", 20, nil, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Welcome to Melodex!").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
