package commands

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
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
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: usr.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	// list all repository for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error: can not get list all repositories")
		return
	}
	fmt.Println(repos)
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
}
