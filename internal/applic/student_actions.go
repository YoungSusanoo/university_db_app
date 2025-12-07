package applic

import (
	"university_app/internal/models"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteStudent(a *App, student models.Student, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteStudent(student)
	if err != nil {
		a.showError(err)
	} else {
		a.tabs.Items[studentTabIndex] = a.createStudentsTab()
		actionsDialog.Dismiss()
	}
}

func getNewStudent(a *App, student *models.Student, callback func()) {
	firstname := widget.NewEntry()
	lastname := widget.NewEntry()
	fathername := widget.NewEntry()
	groups, err := a.db.GetGroups()
	if err != nil {
		a.showError(err)
		return
	}
	group := widget.NewSelectEntry(models.GroupsToNames(groups))
	dialog.ShowForm(
		"Новый преподаватель",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Имя", Widget: firstname},
			{Text: "Фамилия", Widget: lastname},
			{Text: "Отчество", Widget: fathername},
			{Text: "Группа", Widget: group},
		},
		func(confirm bool) {
			if confirm {
				*student = models.Student{student.Id, firstname.Text, lastname.Text, fathername.Text, group.Text}
				callback()
			}
		},
		a.window,
	)
}

func showStudentNewForm(a *App) {
	var studentNew models.Student
	getNewStudent(a, &studentNew, func() {
		err := a.db.InsertStudent(studentNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[studentTabIndex] = a.createStudentsTab()
		}
	})
}

func showStudentEditForm(a *App, student models.Student, actionsDialog *dialog.CustomDialog) {
	var studentNew models.Student
	getNewStudent(a, &studentNew, func() {
		err := a.db.UpdateStudent(student, studentNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[studentTabIndex] = a.createStudentsTab()
			actionsDialog.Dismiss()
		}
	})
}
