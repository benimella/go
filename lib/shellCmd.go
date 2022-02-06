package lib

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	shellCMD.Flags().StringP("name", "n", "", "set remote config name")
	shellCMD.Flags().StringP("command", "c", "", "set shell script")
	RootCmd.AddCommand(shellCMD)
}

var shellCMD = &cobra.Command{
	Use:   "shell",
	Short: "shell exec",
	Run: func(cmd *cobra.Command, args []string) {
		remoteName := mustFlag("name", "string", cmd).(string)
		command := mustFlag("command", "string", cmd).(string)
		remote := SysConfig.GetRemote(remoteName)
		if remote == nil {
			log.Fatal("no such remote")
		}
		session, err := SSHConnect(remote.User, remote.Pwd, remote.Host, 22)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		session.Stderr = os.Stderr
		session.Stdout = os.Stdout
		err = session.Run(command)
		if err != nil {
			log.Fatal(err)
		}
	},
}
