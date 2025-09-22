package twinword

type AnalysisResponse struct {
	Type       string    `json:"type"`
	Score      float64   `json:"score"`
	Ratio      float64   `json:"ratio"`
	ResultCode string    `json:"result_code"`
	ResultMsg  string    `json:"result_msg"`
	Author     string    `json:"author"`
	Email      string    `json:"email"`
	Version    string    `json:"version"`
	Keywords   []Keyword `json:"keywords"`
}

type Keyword struct {
	Word  string  `json:"word"`
	Score float64 `json:"score"`
}
