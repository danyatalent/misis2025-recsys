package apilayer

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danyatalent/misis2025-recsys/lab01/pkg/usecase/entity"
)

const (
	HeaderKey = "apikey"
	Endpoint  = "/sentiment/analysis"
)

type Config struct {
	BaseURL string
	Key     string
	Timeout time.Duration
}

type Client struct {
	client  *http.Client
	BaseURL string
	Key     string
}

func New(cfg Config) (*Client, error) {
	client := &http.Client{
		Timeout: cfg.Timeout,
	}

	return &Client{
		client:  client,
		BaseURL: cfg.BaseURL,
		Key:     cfg.Key,
	}, nil
}

// SentimentAnalysis
//
// Функция запроса к API Apilayer
func (c *Client) SentimentAnalysis(ctx context.Context, textToAnalyze string) (entity.SentimentResult, error) {
	path, err := url.JoinPath(c.BaseURL, Endpoint)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(textToAnalyze)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set(HeaderKey, c.Key)

	slog.Debug("apilayer request started")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result AnalysisResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	slog.Debug("apilayer response", slog.Any("result", result))

	return ResponseToEntity(result), nil
}
