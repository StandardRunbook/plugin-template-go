package config_test

import (
	"github.com/StandardRunbook/plugin-template-go/app/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigArguments(t *testing.T) {
	t.Parallel()
	cfg, err := config.LoadConfigFromEnv(testPeriodicTrigger)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	expectedConfig := &config.Config{
		Name:           "PeriodicTrigger",
		Version:        "v1.0.0",
		ExpectedOutput: "run.go run.sh run_test.go",
		ScriptArguments: []string{
			"--periodic-trigger",
		},
	}
	require.Equal(t, *expectedConfig, *cfg)
}

const testPeriodicTrigger = `
name: PeriodicTrigger
version: v1.0.0
expected_output: run.go run.sh run_test.go
arguments:
 - --periodic-trigger
`
