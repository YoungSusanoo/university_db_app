package applic

import (
	"fmt"
	"strconv"
	"strings"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func deleteMark(a *App, mark models.Mark, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteMark(mark)
	if err != nil {
		a.showError(err)
	} else {
		a.refreshTabs()
		actionsDialog.Dismiss()
	}
}

func getNewMark(a *App, mark *models.Mark, callback func()) {
	teachers, err := a.db.GetTeachers()
	if err != nil {
		a.showError(err)
		return
	}
	students, err := a.db.GetStudents()
	if err != nil {
		a.showError(err)
	}
	subjects, err := a.db.GetSubjects()
	if err != nil {
		a.showError(err)
	}

	teacher := widget.NewSelectEntry(models.TeachersToStrings(teachers))
	student := widget.NewSelectEntry(models.StudentsToStrings(students))
	subject := widget.NewSelectEntry(models.SubjectsToStrings(subjects))

	value := widget.NewEntry()
	dialog.ShowForm(
		"Новый преподаватель",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Преподаватель", Widget: teacher},
			{Text: "Студент", Widget: student},
			{Text: "Предмет", Widget: subject},
			{Text: "Оценка", Widget: value},
		},
		func(confirm bool) {
			if confirm {
				teachStrs := strings.Fields(teacher.Text)
				studStrs := strings.Fields(student.Text)
				valueInt, err := strconv.Atoi(value.Text)
				if err != nil {
					a.showError(err)
				}
				*mark = models.Mark{
					-1,
					teacherFromStrings(teachStrs),
					studentFromStrings(studStrs),
					models.Subject{
						-1,
						subject.Text,
					},
					valueInt,
				}
				callback()
			}
		},
		a.window,
	)
}

func showAvgDialog(a *App) {
	avgBind := binding.NewString()

	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()
	avgEntry := widget.NewLabelWithData(avgBind)
	avgEntry.TextStyle.Monospace = true

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

	btn := widget.NewButton("Рассчитать", func() {
		startInt, err := strconv.Atoi(startEntry.Text)
		if err != nil {
			a.showError(err)
		}
		endInt, err := strconv.Atoi(endEntry.Text)
		if err != nil {
			a.showError(err)
		}

		const (
			groupSelectedPos = 0
		)
		switch filterType.SelectedIndex() {
		case groupSelectedPos:
			avgNum, err := a.db.GetAvgGroupRange(startInt, endInt, valueToFind)
			if err != nil {
				a.showError(err)
				return
			}
			avgBind.Set(fmt.Sprintf("%f", avgNum))
		}
	})

	gridColumn1 := container.NewVBox(widget.NewLabel("Начало"), widget.NewLabel("Конец"), widget.NewLabel("Среднее"))
	gridColumn2 := container.NewVBox(startEntry, endEntry, avgEntry)
	content := container.NewVBox(
		widget.NewLabelWithStyle("Рассчитать средний", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.New(layout.NewGridLayout(2), gridColumn1, gridColumn2),
		container.NewHBox(filterType, filterValue),
		btn,
	)

	dialog.ShowCustom("Выберите действие", "Закрыть", content, a.window)
}

func showMarkNewForm(a *App) {
	var markNew models.Mark
	getNewMark(a, &markNew, func() {
		err := a.db.InsertMark(markNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[markTabIndex] = a.createMarksTab()
		}
	})
}

func showMarkEditForm(a *App, mark models.Mark, actionsDialog *dialog.CustomDialog) {
	var markNew models.Mark
	getNewMark(a, &markNew, func() {
		err := a.db.UpdateMark(mark, markNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[markTabIndex] = a.createMarksTab()
			actionsDialog.Dismiss()
		}
	})
}

func teacherFromStrings(strs []string) (t models.Teacher) {
	t = models.Teacher{}
	if len(strs) > 0 {
		t.FirstName = strs[0]
	}
	if len(strs) > 1 {
		t.LastName = strs[1]
	}
	if len(strs) > 2 {
		t.FatherName = strings.Join(strs[2:], " ")
	}
	return
}

func studentFromStrings(strs []string) (s models.Student) {
	s = models.Student{}
	if len(strs) > 1 {
		s.FirstName = strs[0]
		s.Group = strs[len(strs)-1]
	}
	if len(strs) > 2 {
		s.LastName = strs[1]
	}
	if len(strs) > 3 {
		s.FatherName = strings.Join(strs[2:len(strs)-1], " ")
	}
	return
}
