package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:          "notion2html",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("missing command line arguments")
	},
}

func main() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}
