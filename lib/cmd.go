package lib

import (

	"fmt"
	"github.com/spf13/cobra"

	"os"
)

var RootCmd = &cobra.Command{
	Use:   "jt-devops",
	Short: "程序员在囧途运维小工具",
	Long: `程序员在囧途运维小工具`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

}
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
