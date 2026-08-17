package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use: "add [numbers...]",
	Short: "Add multiple numbers together",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		var result float64

		for _, arg := range args {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
			fmt.Printf("Invalid number: %s\n", arg)
			os.Exit(1)
			}

			result += num
		}

		
		fmt.Printf("%.2f\n", result)
	},
}