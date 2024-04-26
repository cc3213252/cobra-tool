package cmd

import (
	"github.com/spf13/cobra"
	"query/pkg"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "查询天气",
	Long:  `查询对应参数的天气`,
	Run: func(cmd *cobra.Command, args []string) {
		weather := pkg.Weather{}
		if len(args) == 0 {
			weather.GetWeather("")
		} else {
			weather.GetWeather(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
