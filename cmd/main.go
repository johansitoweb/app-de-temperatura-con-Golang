package main

import (
	"app-clima/internal/weather"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {

	var cmdTemperature = &cobra.Command{
		Use:   "temperature [string to print]",
		Short: "Shows the temperature of a location.",
		Long:  `Displays the temperature of a location by specifying its name as a command parameter.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(weather.GetTemperature(strings.Join(args, " ")))

		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdTemperature)
	rootCmd.Execute()
}
