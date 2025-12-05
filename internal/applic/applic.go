package applic

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"university_app/internal/models"
	"university_app/internal/storage"
)

type App struct {
	fyneApp fyne.App
	window  fyne.Window
	db      *storage.Storage
	user    *models.User
}

func NewApp() *App {
	a := app.New()
	w := a.NewWindow("Сесетевой Город")

	return &App{a, w, nil, nil}
}

func (a *App) Run() {
	a.showLoginScreen()
	a.window.ShowAndRun()
}

func (a *App) showLoginScreen() {
	login := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := &widget.Form{Items: []*widget.FormItem{
		{Text: "Логин", Widget: login},
		{Text: "Пароль", Widget: password},
	},
		OnSubmit: func() {
			a.authorize(login.Text, password.Text)
		},
	}
	a.window.SetContent(container.NewVBox(widget.NewLabel("Войдите в систему"), form))
	a.window.ShowAndRun()
}

func (a *App) authorize(login, password string) error {
	var err error
	a.db, err = storage.NewStorage(fmt.Sprintf("postgres://%s:%s@localhost:5432/Megatron?sslmode=disable", login, password))
	if err != nil {
		dialog.ShowError(err, a.window)
		return err
	}
	a.user, err = a.db.Authorize(login, password)
	if err != nil {
		dialog.ShowError(err, a.window)
	}
	return nil
}
