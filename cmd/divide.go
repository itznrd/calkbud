package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var dividecmd = &cobra.Command{
	Use: "div [numbers...]",
	Short: "Divide the first number by consecutive numbers",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		firstNum, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", args[0])
			os.Exit(1)
			}
		
		result := firstNum

		for _, arg := range args[1:] {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Invalid number: %s\n", arg)
				os.Exit(1)
			}
			if num == 0 {
				fmt.Println("Error: division by zero is not allowed")
				os.Exit(1)
			}

			result /= num	
		}

		fmt.Printf("%.2f\n", result)
	},
}