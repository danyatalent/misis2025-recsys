package presentation

import (
	"context"

	"github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
)

type UseCase interface {
	AnalyzeStream(ctx context.Context, text string) <-chan entity.AnalyzeResult
}
