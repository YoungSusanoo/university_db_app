package main

import (
	"db_course/application"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")

	tabs := application.GetLoginScreen()
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}
