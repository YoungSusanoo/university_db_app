package applic

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	fyneApp fyne.App
	window  fyne.Window
	isAdmin bool
}

func NewApp() *App {
	a := app.New()
	w := a.NewWindow("Сесетевой Город")

	return &App{a, w, false}
}

func (a *App) ShowLoginScreen() {
	login := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := &widget.Form{Items: []*widget.FormItem{
		{Text: "Логин", Widget: login},
		{Text: "Пароль", Widget: password},
	},
		OnSubmit: func() {
			fmt.Println("AUF")
		},
	}
	a.window.SetContent(container.NewVBox(widget.NewLabel("Войдите в систему"), form))
	a.window.ShowAndRun()
}
