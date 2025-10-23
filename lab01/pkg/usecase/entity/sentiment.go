package entity

import "time"

// OverallSentimentAnalysisResult
//
// Общая структура для результатов
type OverallSentimentAnalysisResult struct {
	Type string
	Time time.Duration
}

// SentimentResult
//
// Общий интерфейс результатов, для получения всей необходимой информации
type SentimentResult interface {
	GetType() string
	GetTime() time.Duration
	SetTime(duration time.Duration)
	GetProvider() string // чтобы знать из какого API пришёл результат
	JSONString() string
}

// AnalyzeResult
//
// структура для результатов и возможной ошибкой
type AnalyzeResult struct {
	Result SentimentResult
	Error  error
}
