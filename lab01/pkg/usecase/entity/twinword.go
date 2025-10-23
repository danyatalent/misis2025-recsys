package entity

import (
	"encoding/json"
	"time"
)

// TwinwordResponse
//
// Структура для API Twinword
type TwinwordResponse struct {
	OverallSentimentAnalysisResult

	Score    float64 // https://www.twinword.com/blog/interpreting-the-score-and-ratio-of-sentiment/
	Ratio    float64 // https://www.twinword.com/blog/interpreting-the-score-and-ratio-of-sentiment
	Keywords []Keyword
}

type Keyword struct {
	Word  string
	Score float64
}

func (r *TwinwordResponse) GetType() string         { return r.Type }
func (r *TwinwordResponse) GetTime() time.Duration  { return r.Time }
func (r *TwinwordResponse) SetTime(t time.Duration) { r.Time = t }
func (r *TwinwordResponse) GetProvider() string     { return "Twinword" }
func (r *TwinwordResponse) JSONString() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(b)
}
