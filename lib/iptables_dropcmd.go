package lib

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	iptablesDropCMD.Flags().IntP("port", "p", 0, "set drop-port")
}

var iptablesDropCMD = &cobra.Command{
	Use:   "drop",
	Short: "drop port",
	Run: func(cmd *cobra.Command, args []string) {
		remoteName := mustFlag("name", "string", cmd).(string)
		port := mustFlag("port", "int", cmd).(int)
		remote := SysConfig.GetRemote(remoteName)
		if remote == nil {
			log.Fatal("no such remote")
		}
		session, err := SSHConnect(remote.User, remote.Pwd, remote.Host, 22)
		if err != nil {
			log.Fatal(err)
		}

		err = session.Run(fmt.Sprintf("iptables  -A INPUT -p tcp --dport %d -j DROP", port))
		if err != nil {
			log.Fatal(err)
		}
		iptablesCMD.Run(cmd, args)
	},
}
