package core

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/asciifaceman/incremental-barber/pkg/bundles"
)

// MainGameWindow is the central wrapper around primary gameplay
func (g *Game) MainGameWindow() {
	w := g.Application.NewWindow(fmt.Sprintf("Incremental Barber v%s", g.Version))
	w.Resize(g.Resolution)
	w.CenterOnScreen()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(bundles.AppBarberPng, nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.LogoutIcon(), func() { os.Exit(0) }),
	)

	currency := widget.NewLabel("Money: $0")

	playerProfile := container.NewVBox(
		widget.NewLabel("Shop Name"),
		currency,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Shop", widget.NewLabel("TODO: Barber shop")),
		container.NewTabItem("Employees", widget.NewLabel("TODO: Shop Employees")),
		container.NewTabItem("Employees", widget.NewLabel("TODO: Shop Employees")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewBorder(toolbar, nil, playerProfile, nil, tabs)
	w.SetContent(content)

	w.Show()
}

// ShopWindow ...
func (g *Game) ShopTab() {

}
