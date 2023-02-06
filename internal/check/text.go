package check

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewTextChecker(c *http.Client) *TextChecker {
	return &TextChecker{client: c}
}

type TextChecker struct {
	client *http.Client
}

func (t *TextChecker) Check(ctx context.Context, url string, cfg map[string]interface{}) error {
	config, err := NewTextCheckerCfg(cfg)
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build request error: %w", err)
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body error: %w", err)
	}
	defer resp.Body.Close()

	if !strings.Contains(string(body), config.expectedText) {
		return CheckFailed
	}

	return nil
}

type textCheckerCfg struct {
	expectedText string
}

func NewTextCheckerCfg(cfg map[string]interface{}) (textCheckerCfg, error) {
	expectedTextInterface, ok := cfg["expected_text"]
	if !ok {
		return textCheckerCfg{}, fmt.Errorf("expected_text is undefined")
	}
	expectedText, ok := expectedTextInterface.(string)
	if !ok {
		return textCheckerCfg{}, fmt.Errorf("incorrect type of expected_text, expected int")
	}

	return textCheckerCfg{
		expectedText: expectedText,
	}, nil
}
