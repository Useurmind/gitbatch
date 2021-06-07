package config

import (
	"fmt"
	"os"
	"path/filepath"

	// "github.com/Useurmind/gitbatch/pkg/output"
	"github.com/Useurmind/gitbatch/pkg/output"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)


type GlobalConfig struct {
	DefaultShell string `yaml:"defaultShell"`
	Shells []*ShellExecConfig `yaml:"shells"`
}

func (c *GlobalConfig) GetShellOrDefault(shellName string) (*ShellExecConfig, error) {
	if shellName == "" {
		output.Writeln("Shell not specified, using default shell %s", c.DefaultShell)
		shellName = c.DefaultShell
	}

	for _, shellConfig := range c.Shells {
		if shellConfig.Name == shellName {
			return shellConfig, nil
		}
	}

	return nil, fmt.Errorf("Could not find shell %s in config", shellName)
}

type ShellExecConfig struct {
	Name string `yaml:"name"` 

	// Path to the binary or only name if on path.
	Bin string `yaml:"bin"`

	Args []string `yaml:"args"`

	// The arg that preceedes a text command to execute.
	CmdArg string `yaml:"cmdArg"`

	// The arg that preceedes a script file to execute.
	ScriptArg string `yaml:"scriptArg"`
}

func GetDefaultConfig() *GlobalConfig {
	return &GlobalConfig{
		DefaultShell: "pwsh",
		Shells: []*ShellExecConfig{
			{
				Name: "pwsh",
				Bin: "pwsh",
				Args: []string{"-NonInteractive", "-NoProfile"},
				CmdArg: "-Command",
				ScriptArg: "-File",
			},
			{
				Name: "wsl",
				Bin: "wsl",
				Args: []string{},
				CmdArg: "--",
				ScriptArg: "",
			},
		},
	}
}

func Init() error {
	location, err := GetLocation()
	if err != nil {
		return err
	}

	_, err = os.Stat(location)
	if !os.IsNotExist(err) {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(location), 0777)
	if err != nil {
		return err
	}

	// output.Writeln("Initialize global config in %s", location)

	data, err := yaml.Marshal(GetDefaultConfig())
	if err != nil {
		return err
	}

	err = os.WriteFile(location, data, 0777)
	if err != nil {
		return err
	}

	return nil
}

func GetLocation() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.ToSlash(fmt.Sprintf("%s/.gitbatch/config", homeDir)), nil;
}

func Get() (*GlobalConfig, error) {
	location, err := GetLocation()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	config := &GlobalConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}