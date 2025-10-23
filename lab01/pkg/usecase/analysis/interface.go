package analysis

import (
	"context"

	"github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
)

type API interface {
	SentimentAnalysis(ctx context.Context, textToAnalyze string) (entity.SentimentResult, error)
}
