package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type URLConfig struct {
	URL          string
	Checks       []string
	MinChecksCnt int
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
	URL          string   `json:"url"`
	Checks       []string `json:"checks"`
	MinChecksCnt int      `json:"min_check_cnt"`
}

func (u urlConfigJSON) ToURLConfig() URLConfig {
	return URLConfig{
		URL:          u.URL,
		Checks:       u.Checks,
		MinChecksCnt: u.MinChecksCnt,
	}
}
