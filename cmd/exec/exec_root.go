package exec

import (
	"github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/core"
	"github.com/Useurmind/gitbatch/pkg/output"
	"github.com/spf13/cobra"
)

type ExecFlags struct {
	ProjectFile string
	TagFilters []string
	Shell string
	Command string
	Script string
	ScriptArgs []string
}

var execFlags ExecFlags = ExecFlags{
	ScriptArgs: []string{},
}

var execRoot = &cobra.Command{
	Use:   "exec",
	Aliases: []string{"e"},
	Short: "Execute batch operations.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    // Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		globConfig, err := config.Get()
		output.CheckErrf(err, "get global config")

		shellConfig, err := globConfig.GetShellOrDefault(execFlags.Shell)
		output.CheckErrf(err, "get shell config")

		projectConfig, err := config.GetProjectConfig(execFlags.ProjectFile)
		output.CheckErrf(err, "get project config")

		tagFilters := config.NewTagFilters(execFlags.TagFilters)

		err = core.ExecuteShell(shellConfig, projectConfig, tagFilters, execFlags.Command, execFlags.Script)
		output.CheckErrf(err, "execute shell")
	},
}

func Init(parent *cobra.Command) {
	execRoot.PersistentFlags().StringVarP(&execFlags.Shell, "shell", "t", "", "The name of the shell in the config file that should be used.")
	execRoot.PersistentFlags().StringVarP(&execFlags.Command, "command", "c", "", "The shell command that should be executed.")
	execRoot.PersistentFlags().StringVarP(&execFlags.Script, "script", "s", "", "The script file that should be executed.")
	execRoot.PersistentFlags().StringVarP(&execFlags.ProjectFile, "project", "p", "", "The file of the project config that should be used as the source of the git repos.")
	execRoot.PersistentFlags().StringSliceVarP(&execFlags.TagFilters, "tag-filter", "f", nil, "Regex filters that are used to select a subset of repos from the project, each filter should have the form <tagName>=<regex>.")
	// execRoot.PersistentFlags().StringArrayVarP(&execFlags.ScriptArgs, "scriptArgs", "", "Arguments to hand over to the script.")

	execRoot.MarkFlagRequired("project")

	parent.AddCommand(execRoot)
}