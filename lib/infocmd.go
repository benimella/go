package lib

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"math"
	"os"
	"time"
)

func init() {
	RootCmd.AddCommand(infoCMD)
}

var infoCMD = &cobra.Command{
	Use:   "info",
	Short: "Print the info ",
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"项目", "数量", "百分比"})

		data := [][]string{}
		//CPU信息
		p_percent, _ := cpu.Percent(time.Second, false)
		p_count, _ := cpu.Counts(true)
		data = append(data, []string{"CPU",
			fmt.Sprintf("%d核", p_count),
			fmt.Sprintf("%.1f%%", p_percent[0])})
		//内存信息
		m, _ := mem.VirtualMemory()
		data = append(data, []string{"内存",
			fmt.Sprintf("%dG", cast.ToInt(math.Ceil(cast.ToFloat64(m.Total/1024/1024/1024)))),
			fmt.Sprintf("%.1f%%", m.UsedPercent)})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()

	},
}
