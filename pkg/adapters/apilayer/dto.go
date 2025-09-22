package apilayer

import "github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"

func ResponseToEntity(response AnalysisResponse) *entity.ApilayerResponse {
	return &entity.ApilayerResponse{
		OverallSentimentAnalysisResult: entity.OverallSentimentAnalysisResult{
			Type: response.Sentiment,
		},
		Score:      response.Score,
		Confidence: response.Confidence,
	}
}
