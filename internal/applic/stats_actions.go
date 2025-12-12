package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	groupSelectedPos = 0
)

func createStatsUtilPanel(a *App) *fyne.Container {
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()

	var valueToFind string
	filterValue := widget.NewSelect([]string{}, func(s string) {
		valueToFind = s
	})
	options := []string{"Группа", "Студент", "Преподаватель"}
	filterType := widget.NewSelect(options, func(s string) {
		switch s {
		case "Группа":
			filterValue.Options, _ = a.db.GetGroupsNoYearSlice()
		case "Преподаватель":
			teachers, err := a.db.GetTeachers()
			if err != nil {
				a.showError(err)
			}
			filterValue.Options = models.TeachersToStrings(teachers)
		case "Студент":
			students, err := a.db.GetStudents()
			if err != nil {
				a.showError(err)
			}
			filterValue.Options = models.StudentsToStrings(students)
		}
	})

	gridColumn1 := container.NewVBox(widget.NewLabel("Начало"), widget.NewLabel("Конец"), widget.NewLabel("Среднее"))
	gridColumn2 := container.NewVBox(startEntry, endEntry)
	content := container.NewVBox(
		widget.NewLabelWithStyle("Рассчитать средний", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.New(layout.NewGridLayout(2), gridColumn1, gridColumn2),
		container.NewHBox(filterType, filterValue),
		btn,
	)
	return content
}

func createCountTableButton(a *App) {
	btn := widget.NewButton("Рассчитать таблицу", func() {
		startInt, err := strconv.Atoi(startEntry.Text)
		if err != nil {
			a.showError(err)
		}
		endInt, err := strconv.Atoi(endEntry.Text)
		if err != nil {
			a.showError(err)
		}

		switch filterType.SelectedIndex() {
		case groupSelectedPos:
			avgNum, err := a.db.GetAvgGroupRange(startInt, endInt, valueToFind)
			if err != nil {
				a.showError(err)
				return
			}
		}
	})
}
