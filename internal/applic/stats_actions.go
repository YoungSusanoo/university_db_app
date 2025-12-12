package applic

import (
	"fmt"
	"strconv"
	"strings"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	yearAvgCols        = 2
	groupSelectedPos   = 0
	studentSelectedPos = 1
	teacherSelectedPos = 2
)

func createStatsUtilPanel(a *App) *fyne.Container {
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()

	filterValue := widget.NewSelect([]string{}, func(s string) {})
	options := []string{"Группа", "Студент", "Преподаватель"}
	filterType := widget.NewSelect(options, func(s string) {
		switch s {
		case "Группа":
			groups, err := a.db.GetGroupsNoYear()
			if err != nil {
				a.showError(err)
			}
			filterValue.Options = models.GroupsNoYearToStrings(groups)
		case "Преподаватель":
			teachers, err := a.db.GetTeachers()
			if err != nil {
				a.showError(err)
			}
			filterValue.Options = models.TeachersToStrings(teachers)
		case "Студент":
			students, err := a.db.GetStudentsNoYearGroup()
			if err != nil {
				a.showError(err)
			}
			filterValue.Options = models.StudentsNoYearGroupsToString(students)
		}
	})

	gridColumn1 := container.NewVBox(widget.NewLabel("Начало"), widget.NewLabel("Конец"), widget.NewLabel("Среднее"))
	gridColumn2 := container.NewVBox(startEntry, endEntry)
	topPanel := container.NewVBox(
		widget.NewLabelWithStyle("Показать статистику", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.New(layout.NewGridLayout(2), gridColumn1, gridColumn2),
		container.NewHBox(filterType, filterValue),
	)
	content := container.NewAdaptiveGrid(1)
	topPanel.Add(createCountTableButton(a, startEntry, endEntry, filterValue, filterType, content))
	return container.NewBorder(topPanel, nil, nil, nil, content)
}

func createCountTableButton(a *App, start, end *widget.Entry, filter, fType *widget.Select, content *fyne.Container) *widget.Button {
	btn := widget.NewButton("Рассчитать таблицу", func() {
		startInt, err := strconv.Atoi(start.Text)
		if err != nil {
			a.showError(err)
		}
		endInt, err := strconv.Atoi(end.Text)
		if err != nil {
			a.showError(err)
		}

		var yearAvg []models.YearAverage
		switch fType.SelectedIndex() {
		case groupSelectedPos:
			yearAvg, err = a.db.GetAvgGroupRange(startInt, endInt, filter.Selected)
		case studentSelectedPos:
			student := studentNoYearGroupFromStrings(strings.Split(filter.Selected, " "))
			yearAvg, err = a.db.GetAvgStudentRange(startInt, endInt, student)
		case teacherSelectedPos:
			yearAvg, err = a.db.GetAvgGroupRange(startInt, endInt, filter.Selected)
		}
		if err != nil {
			a.showError(err)
			return
		}

		if len(content.Objects) > 0 {
			content.Objects[0] = createYearAvgTable(yearAvg)
		} else {
			content.Objects = append(content.Objects, createYearAvgTable(yearAvg))
		}
	})
	return btn
}

func createYearAvgTable(yearAvg []models.YearAverage) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(yearAvg) + 1, yearAvgCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				headers := []string{"Год", "Среднее"}
				label.SetText(headers[tci.Col])
				label.TextStyle.Bold = true
			} else {
				curr := yearAvg[tci.Row-1]
				switch tci.Col {
				case 0:
					label.SetText(strconv.FormatInt(curr.Year, 10))
				case 1:
					label.SetText(fmt.Sprintf("%f", curr.Avg))
				}
			}
		},
	)
	return table
}
