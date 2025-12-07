package applic

import (
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func showActions[T models.Model](form func(*App, T, *dialog.CustomDialog), delete func(*App, T, *dialog.CustomDialog), a *App, model T) {
	content := container.NewVBox(
		widget.NewLabelWithStyle("Действия", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
	)

	actionsDialog := dialog.NewCustom("Выберите действие", "Закрыть", content, a.window)
	editBtn := widget.NewButton("Редактировать", func() {
		form(a, model, actionsDialog)
	})

	deleteBtn := widget.NewButton("Удалить", func() {
		dialog.ShowCustomConfirm(
			"Точно удалить?",
			"УДАЛИТЬ",
			"Отмена",
			widget.NewLabel("Это действие нельзя отменить"),
			func(confirm bool) {
				if confirm {
					delete(a, model, actionsDialog)
				}
			},
			a.window,
		)
	})
	content.Add(editBtn)
	content.Add(deleteBtn)
	actionsDialog.Show()
}

func deleteSubject(a *App, subject models.Subject, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteSubject(subject)
	if err != nil {
		a.showError(err)
	} else {
		a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
		actionsDialog.Dismiss()
	}
}

func getSubjectEditForm(a *App, subject models.Subject, actionsDialog *dialog.CustomDialog) {
	name := widget.NewEntry()

	dialog.ShowForm(
		"Редактирование",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Название", Widget: name},
		},
		func(confirm bool) {
			subjectNew := models.Subject{subject.Id, name.Text}
			err := a.db.UpdateSubject(subject, subjectNew)
			if err != nil {
				a.showError(err)
			} else {
				a.tabs.Items[subjectTabIndex] = a.createSubjectsTab()
				actionsDialog.Dismiss()
			}
		},
		a.window,
	)
}
