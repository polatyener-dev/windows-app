package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main()  {
	a := app.New()

	w := a.NewWindow("Windows Form Title")
	lbl := widget.NewLabel("Created Form!")
	w.Resize(fyne.NewSize(400,400))
	w.SetContent(lbl)

	w.ShowAndRun()
}