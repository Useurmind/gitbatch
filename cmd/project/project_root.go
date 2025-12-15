package project

import (
	"github.com/spf13/cobra"
)

var projectRoot = &cobra.Command{
	Use:   "project",
	Short: "Perform operations against project configs.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    // Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(projectRoot)

	InitExtend(projectRoot)
}