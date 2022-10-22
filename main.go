package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main()  {
	a := app.New()

	w := a.NewWindow("Windows Form Title")

	w.Resize(fyne.NewSize(400,400))
	w.SetContent(widget.NewLabel("Form Oluşturuldu!"))

	w.ShowAndRun()
}