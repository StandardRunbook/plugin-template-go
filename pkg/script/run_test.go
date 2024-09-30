package script

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplate_Init(t *testing.T) {
	cfg := map[string]string{
		"Name":    "TestPlugin",
		"Version": "1.0.0",
		"Hobbies": "running, swimming, biking",
	}
	tmpl := &Template{}
	err := tmpl.Init(cfg)
	require.Nil(t, err, "expected no error")
	require.Equal(t, "TestPlugin", tmpl.name, "script name should be initialized")
	require.Equal(t, "1.0.0", tmpl.version, "script version should be initialized")
}

func TestTemplate_Name(t *testing.T) {
	tmpl := &Template{name: "TestPlugin"}
	name, err := tmpl.Name()
	require.Nil(t, err, "expected no error")
	require.Equal(t, "TestPlugin", name, "script name should match")
}

func TestTemplate_Name_Empty(t *testing.T) {
	tmpl := &Template{name: ""}
	_, err := tmpl.Name()
	require.NotNil(t, err, "expected error for empty name")
}

func TestTemplate_Version(t *testing.T) {
	tmpl := &Template{version: "1.0.0"}
	version, err := tmpl.Version()
	require.Nil(t, err, "expected no error")
	require.Equal(t, "1.0.0", version, "script version should match")
}

func TestTemplate_Version_Empty(t *testing.T) {
	tmpl := &Template{version: ""}
	_, err := tmpl.Version()
	require.NotNil(t, err, "expected error for empty version")
}

func TestTemplate_Run(t *testing.T) {
	tmpl := &Template{
		name:      "TestPlugin",
		arguments: []string{},
	}

	// Create a mock script
	scriptContent := []byte("#!/bin/bash\necho 'Hello, World!'")
	runScript = scriptContent

	err := tmpl.Run()
	require.Nil(t, err, "expected no error running script")
	require.Contains(t, tmpl.output, "Hello, World!", "expected script output to be 'Hello, World!'")
}

func TestTemplate_Run_Failure(t *testing.T) {
	tmpl := &Template{
		name:      "TestPlugin",
		arguments: []string{},
	}

	// Create an invalid script
	runScript = []byte("#!/bin/bash\nexit 1")

	err := tmpl.Run()
	require.NotNil(t, err, "expected error running failing script")
}

func TestTemplate_ParseOutput_Success(t *testing.T) {
	tmpl := &Template{
		output:         "Hello, World!",
		expectedOutput: "Hello",
	}
	result, err := tmpl.ParseOutput()
	require.Nil(t, err, "expected no error")
	require.Equal(t, "success", result, "expected success when output matches expected")
}

func TestTemplate_ParseOutput_Failure(t *testing.T) {
	tmpl := &Template{
		output:         "Goodbye, World!",
		expectedOutput: "Hello",
	}
	result, err := tmpl.ParseOutput()
	require.Nil(t, err, "expected no error")
	require.Equal(t, "failure", result, "expected failure when output does not match expected")
}
