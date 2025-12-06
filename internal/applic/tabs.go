package applic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createStudentsTab() fyne.Widget {
	return widget.NewLabel("Студенты")
}

func (a *App) createTeachersTab() fyne.Widget {
	return widget.NewLabel("Преподы")
}

func (a *App) createGroupsTab() fyne.Widget {
	return widget.NewLabel("Группы")
}

func (a *App) createSubjectsTab() fyne.Widget {
	return widget.NewLabel("Предметы")
}

func (a *App) createMarksTab() fyne.Widget {
	return widget.NewLabel("Оценки")
}
