package script

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"

	pluginInterface "github.com/StandardRunbook/plugin-interface/shared"
)

//go:embed run.sh
var runScript []byte

// Template is a placeholder - please change to be unique to your script name
type Template struct {
	name           string   `yaml:"Name"`
	version        string   `yaml:"Version"`
	arguments      []string `yaml:"Arguments"`
	output         string   `yaml:"Output"`
	expectedOutput string   `yaml:"ExpectedOutput"`
}

func (t *Template) Init(cfg map[string]string) error {
	t.name = cfg["Name"]
	t.version = cfg["Version"]
	t.arguments = strings.Split(cfg["Arguments"], ",")
	return nil
}

func (t *Template) Name() (string, error) {
	if strings.EqualFold(t.name, "") {
		return "", fmt.Errorf("script name is empty")
	}
	return t.name, nil
}

func (t *Template) Version() (string, error) {
	if strings.EqualFold(t.version, "") {
		return "", fmt.Errorf("script version is empty")
	}
	return t.version, nil
}

func (t *Template) Run() error {
	// Step 1: Create a temporary file
	tmpFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Ensure the file is removed after execution

	// Step 2: Write the embedded script to the temporary file
	pluginName, err := t.Name()
	if err != nil {
		return err
	}
	_, err = tmpFile.Write(runScript)
	if err != nil {
		return fmt.Errorf("failed to write '%s' script to temporary file: %w", pluginName, err)
	}

	// Step 3: Close the file to flush writes and prepare it for execution
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	// Step 4: Set the appropriate permissions to make the script executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		return fmt.Errorf("failed to set executable permissions on file: %w", err)
	}

	// Step 5: Execute the script
	cmd := exec.Command(tmpFile.Name(), t.arguments...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing script: %w", err)
	}
	t.output = string(output)
	return nil
}

func (t *Template) ParseOutput() (string, error) {
	if strings.Contains(t.output, t.expectedOutput) {
		return "success", nil
	}
	return "failure", nil
}

func NewPluginTemplate() pluginInterface.IPlugin {
	return &Template{}
}
