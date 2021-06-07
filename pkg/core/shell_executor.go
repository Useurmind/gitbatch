package core

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/output"
)

func ExecuteShell(config *config.ShellExecConfig, command string, script string) error {
	args := config.Args
	if command != "" && script != "" {
		return fmt.Errorf("Command and script were set for shell execution, but only one may be set.")
	}

	if command == "" && script == "" {
		return fmt.Errorf("Neither command nor script were set for shell execution, but one must be set.")
	}

	if command != "" {
		if config.CmdArg == "" {
			return fmt.Errorf("Shell %s does not support command execution, no cmdArg was set", config.Name)
		}
		args = append(args, config.CmdArg, command)
	}
	if script != "" {
		if config.ScriptArg == "" {
			return fmt.Errorf("Shell %s does not support script execution, no scriptArg was set", config.Name)
		}

		args = append(args, config.ScriptArg, script)
	}

	cmd := exec.Command(config.Bin, args...)

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
