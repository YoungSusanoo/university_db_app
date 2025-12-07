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
