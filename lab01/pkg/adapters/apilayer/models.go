package apilayer

type AnalysisResponse struct {
	Sentiment   string  `json:"sentiment"`
	Score       int64   `json:"score"`
	Confidence  float64 `json:"confidence"`
	Language    string  `json:"language"`
	ContentType string  `json:"content_type"`
}
