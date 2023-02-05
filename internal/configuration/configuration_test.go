package configuration_test

import (
	"http_healthchecker/internal/configuration"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigurationSuccess(t *testing.T) {
	config, err := configuration.GetConfigFromFile("../../testdata/config.json")
	require.NoError(t, err)

	assert.Equal(t, []configuration.URLConfig{
		{
			URL:          "https://google.com",
			Checks:       []string{"status_code", "text"},
			MinChecksCnt: 0,
		},
		{
			URL:          "https://youtube.com",
			Checks:       []string{"status_code"},
			MinChecksCnt: 0,
		},
		{
			URL:          "https://aws.amazon.com",
			Checks:       []string{"status_code", "text"},
			MinChecksCnt: 0,
		},
	}, config)
}
