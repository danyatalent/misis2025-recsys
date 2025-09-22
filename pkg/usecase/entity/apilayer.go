package entity

import (
	"encoding/json"
	"time"
)

// ApilayerResponse
//
// Структура для API Apilayer
type ApilayerResponse struct {
	OverallSentimentAnalysisResult

	Score      int64   // неясно, что значит
	Confidence float64 // неясно, что значит
}

func (r *ApilayerResponse) GetType() string         { return r.Type }
func (r *ApilayerResponse) GetTime() time.Duration  { return r.Time }
func (r *ApilayerResponse) SetTime(t time.Duration) { r.Time = t }
func (r *ApilayerResponse) GetProvider() string     { return "Apilayer" }
func (r *ApilayerResponse) JSONString() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(b)
}
