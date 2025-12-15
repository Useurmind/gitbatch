package project

import (
	"github.com/spf13/cobra"

	"github.com/Useurmind/gitbatch/pkg/config"
	"github.com/Useurmind/gitbatch/pkg/output"
)

type ProjectExtendFlags struct {
	ProjectFile string
	ReposPath   string
}

var projectExtendFlags ProjectExtendFlags

var projectExtend = &cobra.Command{
	Use:   "extend",
	Short: "Extend a project config.",
	// Long: `With gitbatch you can perform batch operations agains multiple git repositories at once.
	// Visit https://github.com/Useurmind/gitbatch to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.ExtendProjectFromSubdirOf(projectExtendFlags.ReposPath, projectExtendFlags.ProjectFile)
		output.CheckErrf(err, "extend project config")
	},
}

func InitExtend(parent *cobra.Command) {
	projectExtend.Flags().StringVarP(&projectExtendFlags.ProjectFile, "project", "p", "", "The project to extend.")
	projectExtend.Flags().StringVarP(&projectExtendFlags.ReposPath, "repos-path", "r", "", "The path under which git repos are located.")

	projectExtend.MarkFlagRequired("project")
	projectExtend.MarkFlagRequired("repos-path")

	parent.AddCommand(projectExtend)
}
