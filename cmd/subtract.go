package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var subtractcmd = &cobra.Command{
	Use: "sub <num> <num>",
	Short: "Subtract second number from the first number",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		num1, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		num2, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		result := num1 - num2
		fmt.Printf("%.2f\n", result)
	},
}