package core

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/output"
)

func ExecuteShell(config *config.ShellExecConfig, projectConfig *config.ProjectConfig, tagFilters *config.TagFilters, command string, script string) error {
	args := config.Args
	if command != "" && script != "" {
		return fmt.Errorf("command and script were set for shell execution, but only one may be set")
	}

	if command == "" && script == "" {
		return fmt.Errorf("neither command nor script were set for shell execution, but one must be set")
	}

	if command != "" {
		if config.CmdArg == "" {
			return fmt.Errorf("shell %s does not support command execution, no cmdArg was set", config.Name)
		}
		args = append(args, config.CmdArg, command)
	}
	if script != "" {
		if config.ScriptArg == "" {
			return fmt.Errorf("shell %s does not support script execution, no scriptArg was set", config.Name)
		}

		args = append(args, config.ScriptArg, script)
	}

	for i, gitRepo := range projectConfig.GetGitRepos() {
		indexInfo := fmt.Sprintf("(%d/%d)", i+1, len(projectConfig.GitRepos))

		match, err := tagFilters.IsMatch(gitRepo)
		if err != nil {
			return err
		}
		if !match { 
			// output.WriteSeparator()
			// output.Writeln("Skipping '%s', no match for tag filter %s", gitRepo.GetName(), indexInfo)
			// output.WriteSeparator()
			continue
		}

		output.WriteSeparator()
		output.Writeln("Executing command in '%s' %s", gitRepo.GetName(), indexInfo)
		output.WriteSeparator()
		err = execCmd(gitRepo.Path, config.Bin, args...)
		if err != nil {
			return err
		}
		output.Writeln("")
	}

	return nil
}

func execCmd(workDir string, bin string, args ...string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(pwd)
	os.Chdir(workDir)

	cmd := exec.Command(bin, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Stderr = cmd.Stdout

	scanner := bufio.NewScanner(stdout)

	cmd.Start()

	for scanner.Scan() {
		out := scanner.Text()
		output.Writeln(out)
	}

	return cmd.Wait()
}