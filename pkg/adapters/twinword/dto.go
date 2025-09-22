package twinword

import (
	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
	"github.com/danyatalent/misis2025-recsys/pkg/utils"
)

func KeywordToEntity(k Keyword) entity.Keyword {
	return entity.Keyword{
		Word:  k.Word,
		Score: k.Score,
	}
}

func ResponseToEntity(response AnalysisResponse) *entity.TwinwordResponse {
	return &entity.TwinwordResponse{
		OverallSentimentAnalysisResult: entity.OverallSentimentAnalysisResult{
			Type: response.Type,
		},
		Score:    response.Score,
		Ratio:    response.Ratio,
		Keywords: utils.ConvertArray(response.Keywords, KeywordToEntity),
	}
}
