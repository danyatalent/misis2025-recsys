package presentation

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// nolint: mnd
// makeCSVResultView
//
// экран с результатами обработки CSV файла
func makeCSVResultView(ctx context.Context, g *GUI, csvPath string) fyne.CanvasObject {
	file, err := os.Open(csvPath)
	if err != nil {
		dialog.ShowError(err, g.window)

		return widget.NewLabel("Ошибка открытия CSV")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		dialog.ShowError(err, g.window)

		return widget.NewLabel("Ошибка чтения CSV")
	}

	// структуры для подсчёта результатов
	type APIMetrics struct {
		Total, Correct int
		Duration       time.Duration
	}

	// создание мапы под метрики
	metrics := make(map[string]*APIMetrics)
	metrics[twinwordName] = &APIMetrics{}
	metrics[apilayerName] = &APIMetrics{}

	// обработка CSV
	for i, row := range records {
		if i == 0 {
			continue // заголовок
		}

		text := row[0]
		expected := row[1]

		resultsCh := g.uc.AnalyzeStream(ctx, text)
		for res := range resultsCh {
			m := metrics[res.Result.GetProvider()]

			m.Total++
			if res.Result.GetType() == expected {
				m.Correct++
			}

			m.Duration += res.Result.GetTime()
		}
	}

	// создание виджетов для отображения метрик
	content := container.NewVBox()

	for apiName, m := range metrics {
		precision := 0.0
		if m.Total > 0 {
			precision = float64(m.Correct) / float64(m.Total)
		}

		label := widget.NewLabel(fmt.Sprintf("%s\nОбщее время: %s\nPrecision: %.2f",
			apiName, m.Duration.String(), precision))
		content.Add(label)
		content.Add(widget.NewSeparator())
	}

	scroll := container.NewScroll(content)
	scroll.SetMinSize(fyne.NewSize(500, 400))

	return scroll
}
