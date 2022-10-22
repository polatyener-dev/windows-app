package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main()  {
	a := app.New()

	w := a.NewWindow("Windows Form Title")
	lbl := widget.NewLabel("Created Form!")
	
	btn := widget.NewButton("Button 1", func() {
		lbl.SetText("Button Clicked")
	})

	btn.Resize(fyne.NewSize(20,100))
	w.Resize(fyne.NewSize(400,400))

	w.SetContent(container.NewVBox(lbl,btn))
	w.ShowAndRun()
}