package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	subjectRows = 2
	studentRows = 5
	teacherRows = 4
	groupsRow   = 2
)

func (a *App) createStudentsTab() *container.TabItem {
	students, err := a.db.GetStudents()
	if err != nil {
		a.showError(err)
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

	var topPanel fyne.CanvasObject
	topPanel = nil
	if a.user.IsAdmin {
		topPanel = addAdminTools(
			showStudentEditForm,
			showStudentNewForm,
			deleteStudent,
			table,
			a,
			students,
		)
	}
	return container.NewTabItem("Студенты", container.NewBorder(topPanel, nil, nil, nil, table))
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

	var topPanel fyne.CanvasObject
	topPanel = nil
	if a.user.IsAdmin {
		topPanel = addAdminTools(
			showTeacherEditForm,
			showTeacherNewForm,
			deleteTeacher,
			table,
			a,
			teachers,
		)
	}
	return container.NewTabItem("Преподаватели", container.NewBorder(topPanel, nil, nil, nil, table))
}

func (a *App) createGroupsTab() *container.TabItem {
	groups, err := a.db.GetGroups()
	if err != nil {
		return container.NewTabItem("Группы", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(groups) + 1, groupsRow
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
				group := groups[tci.Row-1]
				switch tci.Col {
				case 0:
					label.SetText(strconv.FormatInt(group.Id, 10))
				case 1:
					label.SetText(group.Name)
				}
			}
		},
	)

	var topPanel fyne.CanvasObject
	topPanel = nil
	if a.user.IsAdmin {
		topPanel = addAdminTools(
			showGroupEditForm,
			showGroupNewForm,
			deleteGroup,
			table,
			a,
			groups,
		)
	}

	return container.NewTabItem("Группы", container.NewBorder(topPanel, nil, nil, nil, table))
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

	var topPanel fyne.CanvasObject
	topPanel = nil
	if a.user.IsAdmin {
		topPanel = addAdminTools(
			showSubjectEditForm,
			showSubjectNewForm,
			deleteSubject,
			table,
			a,
			subjects,
		)
	}

	return container.NewTabItem("Предметы", container.NewBorder(topPanel, nil, nil, nil, table))
}

func (a *App) createMarksTab() *container.TabItem {
	return container.NewTabItem("Оценки", widget.NewLabel("Оценки"))
}

func addAdminTools[T models.Model](
	edit func(*App, T, *dialog.CustomDialog),
	newForm func(*App),
	delete func(*App, T, *dialog.CustomDialog),
	table *widget.Table,
	a *App,
	objects []T,
) fyne.CanvasObject {
	table.OnSelected = func(id widget.TableCellID) {
		showActions(edit, delete, a, objects[id.Row-1])
	}
	btn := widget.NewButton("Добавить", func() {
		newForm(a)
	})
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btn)
}
