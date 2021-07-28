package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = NewRootCmd()
	cmd     *cobra.Command
)

func NewRootCmd() *cobra.Command {
	cmd = &cobra.Command{
		Use:     "tdd",
		Short:   "tdd tips",
		Long:    "provides tips for practising tdd",
		Version: "0.1v",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
				return nil
			}
			return nil
		},
	}
	return cmd
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
