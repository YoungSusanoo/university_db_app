package applic

import (
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	subjectCols = 2
	studentCols = 5
	teacherCols = 4
	groupsCols  = 2
	marksCols   = 6
)

func (a *App) createStudentsTab() *container.TabItem {
	students, err := a.db.GetStudents()
	if err != nil {
		a.showError(err)
		return container.NewTabItem("Студенты", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := createStudentsTable(students)

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

	table := createTeachersTable(teachers)

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

	table := createGroupsTable(groups)

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

	table := createSubjectsTable(subjects)

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
	marks, err := a.db.GetMarks()
	if err != nil {
		a.showError(err)
		return container.NewTabItem("Оценки", widget.NewLabel("Не удалось загрузить данные"))
	}

	table := createMarksTable(marks)

	var topPanel fyne.CanvasObject
	if a.user.IsAdmin {
		topPanel = addAdminTools(
			showMarkEditForm,
			showMarkNewForm,
			deleteMark,
			table,
			a,
			marks,
		)
	}

	return container.NewTabItem("Оценки", container.NewBorder(topPanel, nil, nil, nil, table))
}

func (a *App) createStatsTab() *container.TabItem {
	return container.NewTabItem("Статистика", createStatsUtilPanel(a))
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
