package analysis

import (
	"context"

	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
)

type API interface {
	SentimentAnalysis(ctx context.Context, textToAnalyze string) (entity.SentimentResult, error)
}
