package analysis

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
)

// AnalyzeStream
//
// Главная функция, которая производит запрос к разным API и получает унифицированный результат
// В отличие от Analyze возвращает канал с результатами. Запросы идут асинхронно
func (u *UseCase) AnalyzeStream(ctx context.Context, text string) <-chan entity.AnalyzeResult {
	out := make(chan entity.AnalyzeResult, len(u.apis))

	var wg sync.WaitGroup
	wg.Add(len(u.apis))

	for _, api := range u.apis {
		// асинхронно вызываются API
		// гарантируется корректное завершение с помощью WaitGroup
		go func(api API) {
			defer wg.Done()

			start := time.Now()

			resp, err := api.SentimentAnalysis(ctx, text)
			if err != nil {
				slog.Error("AnalyzeStream err:", slog.Any("err", err))

				out <- entity.AnalyzeResult{Error: err}

				return
			}

			resp.SetTime(time.Since(start))

			out <- entity.AnalyzeResult{Result: resp}
		}(api)
	}

	// Закроем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
