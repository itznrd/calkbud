package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/spf13/cobra"
)

var precision int = 2

var evalcmd = &cobra.Command{
	Use:                "eval \"[expression]\"",
	Short:              "Evaluate a mathematical expression",
	DisableFlagParsing: true, // Manual parsing prevents pflag from dropping negative math operators (e.g. "-7-9")

	Run: func(cmd *cobra.Command, args []string) {

		// Reset precision default per command execution
		precision = 2

		var exprTokens []string

		// Check flags manually
		for i := 0; i < len(args); i++ {
			arg := args[i]

			if arg == "-h" || arg == "--help" {
				cmd.Help()
				return
			}

			// Registered Flag: Precision (-p or --precision)
			if (arg == "-p" || arg == "--precision") && i+1 < len(args) {
				val, err := strconv.Atoi(args[i+1])
				if err == nil {
					precision = val
					i++
					continue
				}
			}

			// Registered Flag: Inline Precision (-p=val or --precision=val)
			if strings.HasPrefix(arg, "-p=") || strings.HasPrefix(arg, "--precision=") {
				parts := strings.SplitN(arg, "=", 2)
				val, err := strconv.Atoi(parts[1])
				if err == nil {
					precision = val
					continue
				}
			}

			// Not a registered flag -> pass to math expression
			exprTokens = append(exprTokens, arg)
		}

		// Check 1: Missing expression (handles `eval` or `eval -p 4` with no math payload)
		if len(exprTokens) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		// Check 2: Unquoted spaces split the expression into multiple arguments
		if len(exprTokens) > 1 {
			fmt.Println("Error: expression should be inside \"\"")
			fmt.Println("type eval -h or eval --help for help ")
			os.Exit(1)
		}

		rawExpr := exprTokens[0]

		// Normalize brackets [] and {} to ()
		replacer := strings.NewReplacer("[", "(", "]", ")", "{", "(", "}", ")")
		cleanExpr := replacer.Replace(rawExpr)

		// Safely inject '*' for implicit multiplication edge cases
		reBeforeOpen := regexp.MustCompile(`(\d|\))\s*\(`)
		cleanExpr = reBeforeOpen.ReplaceAllString(cleanExpr, "${1}*(")

		reAfterClose := regexp.MustCompile(`\)\s*(\d|\()`)
		cleanExpr = reAfterClose.ReplaceAllString(cleanExpr, ")*${1}")

		// Validate and evaluate expression
		expression, err := govaluate.NewEvaluableExpression(cleanExpr)
		if err != nil {
			fmt.Println("Error: invalid expression")
			fmt.Println("type eval -h or eval --help for help ")
			os.Exit(1)
		}

		result, err := expression.Evaluate(make(map[string]interface{}))
		if err != nil {
    		fmt.Println("Error: invalid expression")
    		fmt.Println("type eval -h or eval --help for help ")
    		os.Exit(1)
		}

		formatStr := fmt.Sprintf("%%.%df\n", precision)
		fmt.Printf(formatStr, result)
	},
}

func init() {
	// Register the flag with Cobra so it appears in `--help` output
	evalcmd.Flags().IntVarP(&precision, "precision", "p", 2, "Number of decimal places in output")
}