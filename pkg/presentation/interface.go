package presentation

import (
	"context"

	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
)

type UseCase interface {
	AnalyzeStream(ctx context.Context, text string) <-chan entity.AnalyzeResult
}
