package config

import (
	"github.com/spf13/cobra"
)

var configRoot = &cobra.Command{
	Use:   "config",
	Short: "Perform operations against the config file of gitbatch.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    // Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func Init(parent *cobra.Command) {
	parent.AddCommand(configRoot)

	configRoot.AddCommand(configPrint)
	configRoot.AddCommand(configPath)
}