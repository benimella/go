package lib

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	scpCMD.Flags().StringP("name", "n", "", "set remote config name")
	scpCMD.Flags().StringP("source", "s", "", "set source path")
	scpCMD.Flags().StringP("dest", "d", "", "set remote path")

	RootCmd.AddCommand(scpCMD)

	//xxx scp -n xxx  -s 源文件路径 -d 目标路径
}

var scpCMD = &cobra.Command{
	Use:   "scp",
	Short: "scp ",
	Run: func(cmd *cobra.Command, args []string) {
		remoteName := mustFlag("name", "string", cmd).(string)
		localPath := mustFlag("source", "string", cmd).(string)
		remotePath := mustFlag("dest", "string", cmd).(string)
		fmt.Println(remoteName, localPath, remotePath)
		remote := SysConfig.GetRemote(remoteName)
		if remote == nil {
			log.Fatal("no such remote")
		}
		session, err := SSHConnect(remote.User, remote.Pwd, remote.Host, 22)
		if err != nil {
			log.Fatal(err)
		}
		// 连上远程主机了
		defer session.Close()
		in, err := session.StdinPipe()
		if err != nil {
			log.Fatal("StdinPipe", err)
		}

		fileInfo, err := os.Open(localPath)
		if err != nil {
			log.Fatal("Open", err)
		}
		fileBytes, err := ioutil.ReadAll(fileInfo)
		if err != nil {
			log.Fatal("ReadAll", err)
		}
		defer fileInfo.Close()
		readerBuffer := bytes.NewBuffer(fileBytes)

		err = session.Start(fmt.Sprintf("/usr/bin/scp %s %s ", "-t", remotePath))
		_, err = fmt.Fprintln(in, "C0644", int64(len(fileBytes)), fileInfo.Name())
		if err != nil {
			log.Fatal("Fprintln", err)
		}
		n, err := io.Copy(in, readerBuffer)
		if err != nil {
			log.Fatal("Copy", err)
		}
		_, err = fmt.Fprint(in, "\x00")
		if err != nil {
			log.Fatal("\\x00", err)
		}

		err = in.Close()
		if err != nil {
			log.Fatal(err)
		}

		err = session.Wait()
		if err != nil {
			log.Fatal("Wait", err)
		}

		fmt.Println("传输了", n, "个字节")
	},
}
