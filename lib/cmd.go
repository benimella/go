package lib

import (

	"fmt"
	"github.com/spf13/cobra"

	"os"
)

var RootCmd = &cobra.Command{
	Use:   "devops",
	Short: "运维小工具",
	Long: `运维小工具`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
