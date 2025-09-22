package presentation

import (
	"context"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
)

// makeResultView
//
// создание окна с результатом
func makeResultView(ctx context.Context, g *GUI, text string) fyne.CanvasObject {
	// Заглушки для результатов
	apilayerBox := newResultBox("Apilayer")
	twinwordBox := newResultBox("Twinword")

	// запускаем асинхронно
	go func() {
		resultsCh := g.uc.AnalyzeStream(ctx, text)
		for res := range resultsCh {
			var box *ResultBox

			provider := ""
			if res.Result != nil {
				provider = res.Result.GetProvider()
			}

			switch provider {
			case "Apilayer":
				box = apilayerBox
			case "Twinword":
				box = twinwordBox
			}

			fyne.Do(func() {
				if res.Error != nil {
					box.ResultLabel.SetText("Ошибка: " + res.Error.Error())

					return
				}

				updateResultBox(box, res.Result)
			})
		}
	}()

	return container.New(layout.NewGridLayout(2), //nolint: mnd
		apilayerBox.Container,
		twinwordBox.Container,
	)
}

// newResultBox
//
// создание контейнеров для API
func newResultBox(title string) *ResultBox {
	rect := canvas.NewRectangle(color.White)
	rect.SetMinSize(fyne.NewSize(defaultRectWidth, defaultRectHeight))

	titleLabel := widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	resultLabel := widget.NewLabel("Ожидание ответа...")

	jsonView := widget.NewMultiLineEntry()
	jsonView.SetPlaceHolder("JSON появится здесь…")
	jsonView.Wrapping = fyne.TextWrapWord
	jsonView.Disable()

	jsonScroll := container.NewScroll(jsonView)
	jsonScroll.SetMinSize(fyne.NewSize(defaultRectWidth, defaultJSONHeight))

	// Верхний блок (название + результат)
	topBox := container.NewVBox(
		titleLabel,
		resultLabel,
	)

	// Border: сверху надписи, в центре JSON
	textBox := container.NewBorder(
		topBox, nil, nil, nil,
		jsonScroll,
	)

	// NewMax для того чтобы контейнеры занимали весь предоставленный объем
	box := container.NewMax(rect, textBox)

	return &ResultBox{
		Container:   box,
		Rect:        rect,
		ResultLabel: resultLabel,
		JSONView:    jsonView,
	}
}

// nolint: mnd
// updateResultBox
//
// функция для обновления результирующих контейнеров
func updateResultBox(box *ResultBox, result entity.SentimentResult) {
	switch result.GetType() {
	case "positive":
		box.Rect.FillColor = color.NRGBA{R: 100, G: 200, B: 100, A: 255}
	case "neutral":
		box.Rect.FillColor = color.NRGBA{R: 200, G: 200, B: 100, A: 255}
	case "negative":
		box.Rect.FillColor = color.NRGBA{R: 200, G: 100, B: 100, A: 255}
	default:
		box.Rect.FillColor = color.NRGBA{R: 180, G: 180, B: 180, A: 255}
	}

	box.Rect.Refresh()

	// Обновляем текст
	box.ResultLabel.SetText("Result: " + result.GetType() + "\nTime: " + result.GetTime().String())

	// Обновляем JSON
	box.JSONView.SetText(result.JSONString())
}
