package  cmd

import (
	"github.com/spf13/cobra"
)


func Execute() {
	var rootcmd = &cobra.Command{
		Use: "calkbud",
	}

	rootcmd.AddCommand(addcmd)
	rootcmd.AddCommand(subtractcmd)

	rootcmd.Execute()
}