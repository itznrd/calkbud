package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var dividecmd = &cobra.Command{
	Use: "div <num> <num>",
	Short: "divides the first number by the second number",
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
		if num2 == 0 {
			fmt.Println("can't divide by zero")
			return
		}

		result := num1 / num2
		fmt.Printf("%.2f\n", result)
	},
}