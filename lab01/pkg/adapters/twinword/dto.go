package twinword

import (
	entity2 "github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
	"github.com/danyatalent/misis2025-recsys/lab01/pkg/utils"
)

func KeywordToEntity(k Keyword) entity2.Keyword {
	return entity2.Keyword{
		Word:  k.Word,
		Score: k.Score,
	}
}

func ResponseToEntity(response AnalysisResponse) *entity2.TwinwordResponse {
	return &entity2.TwinwordResponse{
		OverallSentimentAnalysisResult: entity2.OverallSentimentAnalysisResult{
			Type: response.Type,
		},
		Score:    response.Score,
		Ratio:    response.Ratio,
		Keywords: utils.ConvertArray(response.Keywords, KeywordToEntity),
	}
}
