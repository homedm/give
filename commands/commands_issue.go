package commands

import (
	"fmt"

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
	// list repository issue for the authenticated user
	owner, err := getGitRepoOwner()
	if err != nil {
		return err
	}

	repoName, err := getGitRepoName()
	if err != nil {
		return err
	}

	issues, _, err := client.Issues.ListByRepo(ctx, owner, repoName, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error: can not get list all repositories")
	}

	for i, issue := range issues {
		fmt.Printf("%d\t%s\t%s\n", *issue.Number, *issue.Labels[i].Name, *issue.Title)
	}
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
}
