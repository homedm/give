package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type options struct {
	num int
}

var (
	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Option about GitHub issue viewer",
		Run:   issueCommand,
	}
	opt = &options{}
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
		if i > opt.num-1 {
			return nil
		}
		// Labelがissueについて無い場合
		fmt.Printf("#%d\tUpdated:%s\t%s\n", *issue.Number, (*issue.UpdatedAt).Format("2006/01/02/"), *issue.Title)
	}
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
	issueCmd.Flags().IntVarP(&opt.num, "num", "n", 10, "integer option")
}
