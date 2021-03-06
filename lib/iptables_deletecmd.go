package lib

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
	"strconv"
)

type lineRange struct {
	start int
	end   int
}

func parseRange(l string) *lineRange {
	ret := &lineRange{}
	reg := regexp.MustCompile("^(?P<start>\\d*)-?(?P<end>\\d*)$")
	if match := reg.FindStringSubmatch(l); len(match) > 0 {
		for i, name := range reg.SubexpNames() {
			if i != 0 && match[i] != "" {
				num, _ := strconv.Atoi(match[i])
				if name == "start" {
					ret.start = num
				}
				if name == "end" {
					ret.end = num
				}

			}
		}
	}
	if ret.start < 0 {
		ret.start = 0
	}
	if ret.end < ret.start {
		ret.end = ret.start
	}
	return ret
}

func init() {
	iptablesDeleteCMD.Flags().StringP("line", "l", "", "delete by line-number")
	iptablesDeleteCMD.Flags().StringP("chain", "c", "INPUT", "set the chain")
	// go run main.go iptables -n h1 -t filter delete -c INPUT -l "2-3"
}

// 根据行号删除
var iptablesDeleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "delete rule by line-number",
	Run: func(cmd *cobra.Command, args []string) {
		remoteName := mustFlag("name", "string", cmd).(string)
		line := mustFlag("line", "string", cmd).(string)
		chain := mustFlag("chain", "string", cmd).(string)
		session := getSession(remoteName)
		session.Stdout = os.Stdout

		in, _ := session.StdinPipe()
		err := session.Shell()
		if err != nil {
			log.Fatal("ccc", err)
		}
		lineRange := parseRange(line)
		fmt.Println(lineRange)
		for i := lineRange.start; i <= lineRange.end; i++ {
			cmd := fmt.Sprintf("iptables  -D %s %d"+
				"", chain, lineRange.start)
			//fmt.Println(cmd)
			in.Write([]byte(cmd + " \n"))
		}

		iptablesCMD.Run(cmd, args)

	},
}
