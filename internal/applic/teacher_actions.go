package applic

import (
	"university_app/internal/models"

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
