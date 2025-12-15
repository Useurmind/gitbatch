package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Useurmind/gitbatch/cmd/config"
	"github.com/Useurmind/gitbatch/cmd/project"
	"github.com/Useurmind/gitbatch/cmd/exec"
	pconfig "github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/output"
)

var rootCmd = &cobra.Command{
	Use:   "gitbatch",
	Short: "A cli to perform batch operations against multiple git repositories.",
	Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func Init() {
	config.Init(rootCmd)
	exec.Init(rootCmd)
	project.Init(rootCmd)
}

func Execute() error {
	err := pconfig.Init()
	output.CheckErrf(err, "init config")

	return rootCmd.Execute()
}
