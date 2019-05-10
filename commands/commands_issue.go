package commands

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

type options struct {
	num int
	all int
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

	// not set all option, default behave
	if opt.all == 0 {
		for i, issue := range issues {
			if i > opt.num-1 {
				return nil
			}
			// Labelがissueについて無い場合
			fmt.Printf("#%d\tUpdated:%s\t%s\n", *issue.Number, (*issue.UpdatedAt).Format("2006/01/02/"), *issue.Title)
		}
		return nil
	}

	f1 := func(issues []*github.Issue, number int) (*github.Issue, error) {
		for _, issue := range issues {
			if *issue.Number == number {
				return issue, nil
			}
		}
		return nil, fmt.Errorf("Error: #%d is not found", opt.all)
	}

	var issue *github.Issue
	if issue, err = f1(issues, opt.all); err != nil {
		return err
	}
	fmt.Printf("#%d\tUpdated:%s\t%s\n", *issue.Number, (*issue.UpdatedAt).Format("2006/01/02/"), *issue.Title)
	fmt.Printf("Labels: ")
	for _, label := range issue.Labels {
		fmt.Printf("%s ", *label.Name)
	}
	fmt.Printf("\n")
	fmt.Printf("Issue URL: %s\n", *issue.HTMLURL)
	fmt.Printf("%s\n", *issue.Body)
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
	issueCmd.Flags().IntVarP(&opt.num, "num", "n", 10, "integer option, the number to display")
	issueCmd.Flags().IntVarP(&opt.all, "all", "a", 0, "integer option, issue number")
}
