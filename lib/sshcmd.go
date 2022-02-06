package lib

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
)

func init() {
	sshCMD.Flags().StringP("server", "s", "", "set ssh server")
	sshCMD.Flags().StringP("user", "u", "", "set ssh username")
	//sshCMD.Flags().StringP("password", "p", "", "set ssh userpassword")

	RootCmd.AddCommand(sshCMD)
}

var sshCMD = &cobra.Command{
	Use:   "ssh",
	Short: "ssh connect",
	Run: func(cmd *cobra.Command, args []string) {
		server := mustFlag("server", "string", cmd).(string)
		user := mustFlag("user", "string", cmd).(string)
		//pwd:=mustFlag("password","string",cmd).(string)
		connectCount := 0
		var session *ssh.Session
		for connectCount < 3 {
			fmt.Println("enter your password")
			pwd, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			session, err = SSHConnect(user, string(pwd), server, 22)
			if err != nil {
				fmt.Println("password error")
				connectCount++
			} else {
				break
			}
		}
		if session == nil {
			log.Fatal("ssh failed")
		}

		//session,err:=SSHConnect(user,pwd,server,22)
		//if err != nil {
		//	log.Fatal(err)
		//}

		session.Stdin = os.Stdin
		session.Stderr = os.Stderr
		session.Stdout = os.Stdout
		err := session.RequestPty("vt220", 0, 0, ShellModes)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		err = session.Shell()
		if err != nil {
			log.Fatal(err)
		}
		session.Wait()

	},
}
