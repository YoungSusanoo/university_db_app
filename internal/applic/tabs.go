package applic

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	subjectRows = 2
	studentRows = 5
	teacherRows = 4
)

func (a *App) createStudentsTab() *container.TabItem {
	students, err := a.db.GetStudents()
	if err != nil {
		return container.NewTabItem("Студенты", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(students) + 1, studentRows
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
	return container.NewTabItem("Студенты", container.NewBorder(nil, nil, nil, nil, table))
}

func (a *App) createTeachersTab() *container.TabItem {
	teachers, err := a.db.GetTeachers()
	if err != nil {
		return container.NewTabItem("Преподаватели", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(teachers) + 1, teacherRows
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
	return container.NewTabItem("Преподаватели", container.NewBorder(nil, nil, nil, nil, table))
}

func (a *App) createGroupsTab() *container.TabItem {
	return container.NewTabItem("Группы", widget.NewLabel("Группы"))
}

func (a *App) createSubjectsTab() *container.TabItem {
	subjects, err := a.db.GetSubjects()
	if err != nil {
		return container.NewTabItem("Предметы", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(subjects) + 1, subjectRows
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

	if a.user.IsAdmin {
		table.OnSelected = func(id widget.TableCellID) {
			showActions(getSubjectEditForm, deleteSubject, a, subjects[id.Row-1])
		}
	}

	return container.NewTabItem("Предметы", container.NewBorder(nil, nil, nil, nil, table))
}

func (a *App) createMarksTab() *container.TabItem {
	return container.NewTabItem("Оценки", widget.NewLabel("Оценки"))
}
