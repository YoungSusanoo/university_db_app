package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteStudent(a *App, student models.Student, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteStudent(student)
	if err != nil {
		a.showError(err)
	} else {
		a.refreshTabs()
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
	group := widget.NewSelectEntry(models.GroupsToStrings(groups))
	dialog.ShowForm(
		"Новый студент",
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

func createStudentsTable(students []models.Student) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(students) + 1, studentCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				headers := []string{"Id", "Имя", "Фамилия", "Отчество", "Группа"}
				label.SetText(headers[tci.Col])
				label.TextStyle.Bold = true
			} else {
				stud := students[tci.Row-1]
				switch tci.Col {
				case 0:
					label.SetText(strconv.FormatInt(stud.Id, 10))
				case 1:
					label.SetText(stud.FirstName)
				case 2:
					label.SetText(stud.LastName)
				case 3:
					label.SetText(stud.FatherName)
				case 4:
					label.SetText(stud.Group)
				}
			}
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	return table
}
