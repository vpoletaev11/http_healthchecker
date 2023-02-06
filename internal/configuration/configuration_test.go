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

	assert.Equal(t, []configuration.URLConfig(
		[]configuration.URLConfig{
			{
				URL: "https://google.com",
				Checks: []configuration.Check{
					{
						Name: "status_code",
						Params: map[string]interface{}{
							"expected_status_code": float64(200),
						},
					}, {
						Name: "text",
						Params: map[string]interface{}{
							"expected_text": "ok",
						},
					},
				},
				MinChecksCnt: 2,
			},
			{
				URL: "https://youtube.com",
				Checks: []configuration.Check{
					{
						Name: "status_code",
						Params: map[string]interface{}{
							"expected_status_code": float64(200),
						},
					},
				},
				MinChecksCnt: 1,
			}, {
				URL: "https://aws.amazon.com",
				Checks: []configuration.Check{
					{
						Name: "status_code",
						Params: map[string]interface{}{
							"expected_status_code": float64(200),
						},
					}, {
						Name: "text",
						Params: map[string]interface{}{
							"expected_text": "ok",
						},
					},
				},
				MinChecksCnt: 1,
			},
		},
	), config)
}
