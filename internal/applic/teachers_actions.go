package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteTeacher(a *App, teacher models.Teacher, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteTeacher(teacher)
	if err != nil {
		a.showError(err)
	} else {
		a.refreshTabs()
		actionsDialog.Dismiss()
	}
}

func getNewTeacher(a *App, teacher *models.Teacher, callback func()) {
	firstname := widget.NewEntry()
	lastname := widget.NewEntry()
	fathername := widget.NewEntry()
	dialog.ShowForm(
		"Новый преподаватель",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Имя", Widget: firstname},
			{Text: "Фамилия", Widget: lastname},
			{Text: "Отчество", Widget: fathername},
		},
		func(confirm bool) {
			if confirm {
				*teacher = models.Teacher{teacher.Id, firstname.Text, lastname.Text, fathername.Text}
				callback()
			}
		},
		a.window,
	)
}

func showTeacherNewForm(a *App) {
	var teacherNew models.Teacher
	getNewTeacher(a, &teacherNew, func() {
		err := a.db.InsertTeacher(teacherNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[teacherTabIndex] = a.createTeachersTab()
		}
	})
}

func showTeacherEditForm(a *App, teacher models.Teacher, actionsDialog *dialog.CustomDialog) {
	var teacherNew models.Teacher
	getNewTeacher(a, &teacherNew, func() {
		err := a.db.UpdateTeacher(teacher, teacherNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[teacherTabIndex] = a.createTeachersTab()
			actionsDialog.Dismiss()
		}
	})
}

func createTeachersTable(teachers []models.Teacher) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(teachers) + 1, teacherCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				headers := []string{"Id", "Имя", "Фамилия", "Отчество"}
				label.SetText(headers[tci.Col])
				label.TextStyle.Bold = true
			} else {
				teach := teachers[tci.Row-1]
				switch tci.Col {
				case 0:
					label.SetText(strconv.FormatInt(teach.Id, 10))
				case 1:
					label.SetText(teach.FirstName)
				case 2:
					label.SetText(teach.LastName)
				case 3:
					label.SetText(teach.FatherName)
				}
			}
		},
	)
	return table
}
