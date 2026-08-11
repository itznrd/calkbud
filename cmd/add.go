package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use: "add <num> <num>",
	Short: "add two numbers",
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

		result := num1 + num2
		fmt.Printf("%.2f\n", result)
	},
}