package applic

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createStudentsTab() fyne.CanvasObject {
	return widget.NewLabel("Студенты")
}

func (a *App) createTeachersTab() fyne.CanvasObject {
	return widget.NewLabel("Преподы")
}

func (a *App) createGroupsTab() fyne.CanvasObject {
	return widget.NewLabel("Группы")
}

func (a *App) createSubjectsTab() fyne.CanvasObject {
	subjects, err := a.db.GetSubjects()
	if err != nil {
		return widget.NewLabel("Не удалось загрузить данные")
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(subjects) + 1, 2
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
	return container.NewBorder(nil, nil, nil, nil, table)
}

func (a *App) createMarksTab() fyne.CanvasObject {
	return widget.NewLabel("Оценки")
}
