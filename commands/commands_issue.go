package commands

import (
	"github.com/spf13/cobra"
)

var (
	issueCmd = &cobra.Command{
		Use: "issue",
		Run: issueCommand,
	}
)

func issueCommand(cmd *cobra.Command, args []string) {
	if err := issueAction(); err != nil {
		Exit(err, 1)
	}
}

func issueAction() (err error) {
	// execution body
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
}
