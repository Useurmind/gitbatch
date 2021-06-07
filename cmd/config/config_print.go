package config

import (
	"os"

	"github.com/spf13/cobra"

	pconfig "github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/output"
)

var configPrint = &cobra.Command{
	Use:   "print",
	Short: "Print the config file.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    // Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := pconfig.GetLocation()
		output.CheckErrf(err, "get global config path")

		content, err := os.ReadFile(filePath)
		output.CheckErrf(err, "read config file %s", filePath)

		output.Writef(string(content))
	},
}

var configPath = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
    // Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := pconfig.GetLocation()
		output.CheckErrf(err, "get global config path")

		output.Writeln(filePath)
	},
}