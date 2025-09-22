package presentation

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// makeMainView
//
// создание главного окна приложения
func makeMainView(ctx context.Context, g *GUI) fyne.CanvasObject {
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Введите текст для анализа...")
	input.SetMinRowsVisible(defaultMinRow)
	input.Wrapping = fyne.TextWrapWord

	// скролл для поля
	inputScroll := container.NewScroll(input)
	inputScroll.SetMinSize(fyne.NewSize(400, 200)) //nolint: mnd

	analyzeBtn := widget.NewButton("Проанализировать", func() {
		text := input.Text
		if text == "" {
			return
		}
		// переход на экран с результатами
		g.window.SetContent(makeResultView(ctx, g, text))
	})

	// Вертикальный бокс с полем и кнопкой
	content := container.NewVBox(
		inputScroll,
		analyzeBtn,
	)

	// Центрируем контент
	centered := container.NewCenter(content)

	// Чтобы занимало всё окно
	return container.NewMax(centered)
}
