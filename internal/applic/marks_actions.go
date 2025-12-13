package applic

import (
	"fmt"
	"strconv"
	"strings"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
		"Новая оценка",
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

func showMarkNewForm(a *App) {
	var markNew models.Mark
	getNewMark(a, &markNew, func() {
		err := a.db.InsertMark(markNew)
		if err != nil {
			a.showError(fmt.Errorf("некорректное значение оценки"))
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

func createMarksTable(marks []models.Mark) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(marks) + 1, marksCols
		},
		func() fyne.CanvasObject {
			return container.NewGridWithColumns(3, widget.NewLabel("t"), widget.NewLabel("t"), widget.NewLabel("t"))
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			cell := co.(*fyne.Container)
			if tci.Row == 0 {
				cell.RemoveAll()
				headers := []string{"Id", "Преподаватель", "Студент", "Группа", "Предмет", "Оценка"}
				cell.Add(widget.NewLabel(headers[tci.Col]))
			} else {
				switch tci.Col {
				case 0:
					cell.RemoveAll()
					cell.Add(widget.NewLabel(strconv.FormatInt(marks[tci.Row-1].Id, 10)))
				case 1:
					cell.Objects[0].(*widget.Label).SetText(marks[tci.Row-1].Teach.FirstName)
					cell.Objects[1].(*widget.Label).SetText(marks[tci.Row-1].Teach.LastName)
					cell.Objects[2].(*widget.Label).SetText(marks[tci.Row-1].Teach.FatherName)
				case 2:
					cell.Objects[0].(*widget.Label).SetText(marks[tci.Row-1].Stud.FirstName)
					cell.Objects[1].(*widget.Label).SetText(marks[tci.Row-1].Stud.LastName)
					cell.Objects[2].(*widget.Label).SetText(marks[tci.Row-1].Stud.FatherName)
				case 3:
					cell.RemoveAll()
					cell.Add(widget.NewLabel(marks[tci.Row-1].Stud.Group))
				case 4:
					cell.RemoveAll()
					cell.Add(widget.NewLabel(marks[tci.Row-1].Subj.Name))
				case 5:
					cell.RemoveAll()
					cell.Add(widget.NewLabel(strconv.FormatInt(int64(marks[tci.Row-1].Value), 10)))
				}
			}
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 400)
	table.SetColumnWidth(2, 400)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 200)
	return table
}
