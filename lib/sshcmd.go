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
	RootCmd.AddCommand(sshCMD)
}

var sshCMD = &cobra.Command{
	Use:   "ssh",
	Short: "ssh connect",
	Run: func(cmd *cobra.Command, args []string) {
		server := mustFlag("server", "string", cmd).(string)
		user := mustFlag("user", "string", cmd).(string)

		var session *ssh.Session // 默认是nil
		connCount := 0
		for connCount < 3 {
			fmt.Println("enter your password")
			pwd, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			session, err = SSHConnect(user, string(pwd), server, 22)
			if err != nil {
				fmt.Println("error password,try again")
				connCount++
			} else {
				break
			}
		}
		if session == nil {
			log.Fatal("ssh failed")
		}

		//设置 stdin等输入输出
		session.Stdin = os.Stdin
		session.Stderr = os.Stderr
		session.Stdout = os.Stdout

		//shell操作代码
		{
			reqErr := session.RequestPty("vt220", 0, 0, ShellModes)
			if reqErr != nil {
				log.Fatal(reqErr)
			}
			defer session.Close()
			shellerr := session.Shell()
			if shellerr != nil {
				log.Fatal(shellerr)
			}
			session.Wait()
		}

	},
}
