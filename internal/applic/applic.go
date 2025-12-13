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

	url    string
	dbName string
	tabs   *container.AppTabs
}

const (
	studentTabIndex = 0
	teacherTabIndex = 1
	groupTabIndex   = 2
	subjectTabIndex = 3
	markTabIndex    = 4
	statsTabIndex   = 5
)

func NewApp(url string, dbName string) *App {
	a := app.New()
	w := a.NewWindow("Сесетевой Город")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(1000, 800))

	app := &App{a, w, nil, nil, url, dbName, nil}
	return app
}

func (a *App) Run() {
	a.showLoginScreen()
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
	a.window.Resize(fyne.NewSize(500, 500))
	a.window.ShowAndRun()
}

func (a *App) showMainScreen() *container.AppTabs {
	a.tabs = container.NewAppTabs(
		a.createStudentsTab(),
		a.createTeachersTab(),
		a.createGroupsTab(),
		a.createSubjectsTab(),
		a.createMarksTab(),
		a.createStatsTab(),
	)
	return a.tabs
}

func (a *App) showError(err error) {
	dialog.ShowError(err, a.window)
}

func (a *App) authorize(login, password string) {
	var err error
	a.db, err = storage.NewStorage(fmt.Sprintf("postgres://%s:%s@%s/%s", login, password, a.url, a.dbName))
	if err != nil {
		a.showError(fmt.Errorf("не удалось подключиться к базе"))
		return
	}
	a.user, err = a.db.Authorize(login, password)
	if err != nil {
		a.showError(fmt.Errorf("неверные логин или пароль"))
		return
	}
	a.window.SetContent(a.showMainScreen())
}

func (a *App) refreshTabs() {
	a.tabs.Items[groupTabIndex] = a.createGroupsTab()
	a.tabs.Items[markTabIndex] = a.createMarksTab()
	a.tabs.Items[studentTabIndex] = a.createStudentsTab()
	a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
	a.tabs.Items[teacherTabIndex] = a.createTeachersTab()
	a.tabs.Items[statsTabIndex] = a.createStatsTab()
}
