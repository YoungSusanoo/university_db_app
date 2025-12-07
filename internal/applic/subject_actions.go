package applic

import (
	"university_app/internal/models"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteSubject(a *App, subject models.Subject, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteSubject(subject)
	if err != nil {
		a.showError(err)
	} else {
		a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
		actionsDialog.Dismiss()
	}
}

func getNewSubject(a *App, subject *models.Subject, callback func()) {
	name := widget.NewEntry()
	dialog.ShowForm(
		"Новый предмет",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Название", Widget: name},
		},
		func(confirm bool) {
			if confirm {
				*subject = models.Subject{Id: subject.Id, Name: name.Text}
				callback()
			}
		},
		a.window,
	)
}

func showSubjectNewForm(a *App) {
	var subjectNew models.Subject
	getNewSubject(a, &subjectNew, func() {
		err := a.db.InsertSubject(subjectNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
		}
	})
}

func showSubjectEditForm(a *App, subject models.Subject, actionsDialog *dialog.CustomDialog) {
	var subjectNew models.Subject
	getNewSubject(a, &subjectNew, func() {
		err := a.db.UpdateSubject(subject, subjectNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
			actionsDialog.Dismiss()
		}
	})
}
