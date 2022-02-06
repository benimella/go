package lib

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
	"time"
)

func init()  {
	RootCmd.AddCommand(cpuCMD)
}

var cpuCMD = &cobra.Command{
	Use:   "cpu",
	Short: "Print the cpu percent ",
	Run: func(cmd *cobra.Command, args []string) {

		for {
			p, _ := cpu.Percent(time.Second, false)
			fmt.Printf("\rCPU:%.1f%%", p[0])
			time.Sleep(time.Second)
		}

	},
}
