package apilayer

import (
	entity2 "github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
)

func ResponseToEntity(response AnalysisResponse) *entity2.ApilayerResponse {
	return &entity2.ApilayerResponse{
		OverallSentimentAnalysisResult: entity2.OverallSentimentAnalysisResult{
			Type: response.Sentiment,
		},
		Score:      response.Score,
		Confidence: response.Confidence,
	}
}
