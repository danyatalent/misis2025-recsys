package analysis

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
)

// Analyze
//
// Главная функция, которая производит запрос к разным API и получает унифицированный результат
func (u *UseCase) Analyze(ctx context.Context, textToAnalyze string) ([]entity.SentimentResult, error) {
	results := make([]entity.SentimentResult, len(u.apis))
	errs := make(chan error, len(u.apis))

	var wg sync.WaitGroup
	wg.Add(len(u.apis))

	for i, api := range u.apis {
		go func(i int, api API) {
			defer wg.Done()

			start := time.Now()

			resp, err := api.SentimentAnalysis(ctx, textToAnalyze)
			if err != nil {
				errs <- err

				return
			}

			resp.SetTime(time.Since(start))
			results[i] = resp
		}(i, api)
	}

	wg.Wait()
	close(errs)

	if len(errs) > 0 {
		// вернём первую ошибку (или можно собрать все)
		return nil, <-errs
	}

	return results, nil
}

// AnalyzeStream
//
// Главная функция, которая производит запрос к разным API и получает унифицированный результат
// В отличие от Analyze возвращает канал с результатами. Запросы идут асинхронно
func (u *UseCase) AnalyzeStream(ctx context.Context, text string) <-chan entity.AnalyzeResult {
	out := make(chan entity.AnalyzeResult, len(u.apis))

	var wg sync.WaitGroup
	wg.Add(len(u.apis))

	for _, api := range u.apis {
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
