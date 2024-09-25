package mk8s

import (
	"github.com/k8s-crafts/mk8s/pkg/mk8s/out"
	"github.com/k8s-crafts/mk8s/pkg/mk8s/version"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mk8s",
	Long:  `Print the version of mk8s`,
	Run: func(cmd *cobra.Command, args []string) {
		out.PrintLn(version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
