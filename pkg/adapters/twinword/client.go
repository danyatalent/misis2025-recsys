package twinword

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/danyatalent/misis2025-recsys/pkg/usecase/entity"
)

const (
	Endpoint  = "/sentiment/analyze/latest/"
	HeaderKey = "X-Twaip-Key"
	TextKey   = "text"
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
// Функция запроса к API Twinword
func (c *Client) SentimentAnalysis(ctx context.Context, textToAnalyze string) (entity.SentimentResult, error) {
	path, err := url.JoinPath(c.BaseURL, Endpoint)
	if err != nil {
		return nil, err
	}

	payload := url.Values{}
	payload.Set(TextKey, textToAnalyze)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(HeaderKey, c.Key)

	slog.Debug("twinword request started")

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

	if result.ResultCode != "200" {
		return nil, fmt.Errorf("unexpected result code: %s", result.ResultCode)
	}

	slog.Debug("twinword response", slog.Any("result", result))

	return ResponseToEntity(result), nil
}
