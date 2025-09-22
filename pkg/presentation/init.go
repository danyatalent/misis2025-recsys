package presentation

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type GUI struct {
	app    fyne.App
	window fyne.Window
	uc     UseCase
}

func New(uc UseCase) (*GUI, error) {
	a := app.New()
	w := a.NewWindow("Sentiment Analyzer")
	w.Resize(fyne.NewSize(defaultWindowWidth, defaultWindowHeight))

	return &GUI{app: a, window: w, uc: uc}, nil
}

func (g *GUI) Run(ctx context.Context, cancel context.CancelFunc) {
	g.window.SetOnClosed(func() {
		cancel()
	})
	g.window.SetContent(makeMainView(ctx, g)) // стартовый экран
	g.window.ShowAndRun()
}

func (g *GUI) Window() fyne.Window {
	return g.window
}
