package main

import (
	"fmt"
	"gds/lib"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
	"time"
)

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

func main() {
	//s1 := "xxxxx1"
	//n, err := fmt.Fprintln(os.Stdout, s1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(n)

	lib.RootCmd.AddCommand(cpuCMD)
	lib.Execute()

}
