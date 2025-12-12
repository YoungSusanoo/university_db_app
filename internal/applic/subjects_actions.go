package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteSubject(a *App, subject models.Subject, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteSubject(subject)
	if err != nil {
		a.showError(err)
	} else {
		a.refreshTabs()
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

func createSubjectsTable(subjects []models.Subject) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(subjects) + 1, subjectCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				headers := []string{"Id", "Имя"}
				label.SetText(headers[tci.Col])
				label.TextStyle.Bold = true
			} else {
				subj := subjects[tci.Row-1]
				switch tci.Col {
				case 0:
					label.SetText(strconv.FormatInt(subj.Id, 10))
				case 1:
					label.SetText(subj.Name)
				}
			}
		},
	)
	return table
}
