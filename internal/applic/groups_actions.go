package applic

import (
	"strconv"
	"university_app/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func deleteGroup(a *App, group models.Group, actionsDialog *dialog.CustomDialog) {
	err := a.db.DeleteGroup(group)
	if err != nil {
		a.showError(err)
	} else {
		a.refreshTabs()
		actionsDialog.Dismiss()
	}
}

func getNewGroup(a *App, group *models.Group, callback func()) {
	name := widget.NewEntry()
	dialog.ShowForm(
		"Новая группа",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Название", Widget: name},
		},
		func(confirm bool) {
			if confirm {
				*group = models.Group{Id: group.Id, Name: name.Text}
				callback()
			}
		},
		a.window,
	)
}

func showGroupNewForm(a *App) {
	var groupNew models.Group
	getNewGroup(a, &groupNew, func() {
		err := a.db.InsertGroup(groupNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[groupTabIndex] = a.createGroupsTab()
		}
	})
}

func showGroupEditForm(a *App, group models.Group, actionsDialog *dialog.CustomDialog) {
	var groupNew models.Group
	getNewGroup(a, &groupNew, func() {
		err := a.db.UpdateGroup(group, groupNew)
		if err != nil {
			a.showError(err)
		} else {
			a.tabs.Items[groupTabIndex] = a.createGroupsTab()
			actionsDialog.Dismiss()
		}
	})
}

func createGroupsTable(groups []models.Group) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(groups) + 1, groupsCols
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
	return table
}
