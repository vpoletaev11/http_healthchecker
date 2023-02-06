package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type URLConfig struct {
	URL          string
	Checks       []Check
	MinChecksCnt int
}

type Check struct {
	Name   string
	Params map[string]interface{}
}

func GetConfigFromFile(path string) ([]URLConfig, error) {
	dataJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q error: %w", path, err)
	}

	configJSON := configJSON{}
	err = json.Unmarshal(dataJSON, &configJSON)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json config: %w", err)
	}

	config := make([]URLConfig, 0, len(configJSON.Urls))
	for _, cfgJSON := range configJSON.Urls {
		config = append(config, cfgJSON.ToURLConfig())
	}

	return config, nil
}

type configJSON struct {
	Urls []urlConfigJSON `json:"urls"`
}

type urlConfigJSON struct {
	URL          string      `json:"url"`
	Checks       []checkJSON `json:"checks"`
	MinChecksCnt int         `json:"min_checks_cnt"`
}

type checkJSON struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

func (u urlConfigJSON) ToURLConfig() URLConfig {
	checks := make([]Check, 0, len(u.Checks))
	for _, c := range u.Checks {
		checks = append(checks, Check{
			Name:   c.Name,
			Params: c.Params,
		})
	}

	return URLConfig{
		URL:          u.URL,
		Checks:       checks,
		MinChecksCnt: u.MinChecksCnt,
	}
}
