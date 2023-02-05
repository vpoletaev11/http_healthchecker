package check

import (
	"context"
	"fmt"
	"net/http"
)

func NewStatusCodeChecker(c *http.Client) *StatusCodeChecker {
	return &StatusCodeChecker{client: c}
}

type StatusCodeChecker struct {
	client *http.Client
}

func (s *StatusCodeChecker) Check(ctx context.Context, url string, cfg map[string]interface{}) error {
	config, err := NewStatusCodeCheckerCfg(cfg)
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build request error: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	if resp.StatusCode != config.expectedStatusCode {
		return fmt.Errorf("status code check fail, expected: %d, actual: %d", config.expectedStatusCode, resp.StatusCode)
	}

	return nil
}

type statusCodeCheckerCfg struct {
	expectedStatusCode int
}

func NewStatusCodeCheckerCfg(cfg map[string]interface{}) (statusCodeCheckerCfg, error) {
	expectedStatusCodeInterface, ok := cfg["expected_status_code"]
	if !ok {
		return statusCodeCheckerCfg{}, fmt.Errorf("expected_status_code is undefined")
	}
	expectedStatusCode, ok := expectedStatusCodeInterface.(int)
	if !ok {
		return statusCodeCheckerCfg{}, fmt.Errorf("incorrect type of expected_status_code, expected int")
	}

	return statusCodeCheckerCfg{
		expectedStatusCode: expectedStatusCode,
	}, nil
}
