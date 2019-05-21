package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

type options struct {
	num   int
	show  int
	close int
	add   string
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
	owner, err := getGitRepoOwner()
	if err != nil {
		return err
	}

	repoName, err := getGitRepoName()
	if err != nil {
		return err
	}

	if opt.add != "" {
		makeIssue(owner, repoName)
		return
	}
	if opt.close != 0 {
		closeIssue(owner, repoName)
		return
	}
	printIssue(owner, repoName)

	return nil
}

func openTextEditor() (string, error) {
	tmpFile, err := ioutil.TempFile("", "tmp")
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	if err != nil {
		return "", err
	}
	editor, err := getGitEditor()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(editor, tmpFile.Name())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// wait to finish text editor
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func makeIssue(owner string, repoName string) error {
	// input body
	body, err := openTextEditor()
	if err != nil {
		return err
	}

	pushIssue := &github.IssueRequest{
		Title: &opt.add,
		Body:  &body,
	}

	_, _, err = client.Issues.Create(ctx, owner, repoName, pushIssue)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	return nil
}

func closeIssue(owner string, repoName string) error {
	state := "close"
	request := &github.IssueRequest{
		State: &state,
	}
	client.Issues.Edit(ctx, owner, repoName, opt.close, request)
	return nil
}

func printIssue(owner string, repoName string) error {
	// get issue list
	issues, _, err := client.Issues.ListByRepo(ctx, owner, repoName, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error: can not get list all repositories")
	}

	// not set all option, default behave
	if opt.show == 0 {
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
		return nil, fmt.Errorf("Error: #%d is not found", opt.show)
	}

	var issue *github.Issue
	if issue, err = f1(issues, opt.show); err != nil {
		return err
	}
	fmt.Printf("#%d\tUpdated:%s\t%s\n", *issue.Number, (*issue.UpdatedAt).Format("2006/01/02/"), *issue.Title)
	fmt.Printf("Labels: ")
	for _, label := range issue.Labels {
		fmt.Printf("%s ", *label.Name)
	}
	fmt.Printf("\nIssue URL: %s\n%s\n", *issue.HTMLURL, *issue.Body)
	return nil
}

func init() {
	RootCmd.AddCommand(issueCmd)
	issueCmd.Flags().IntVarP(&opt.num, "num", "n", 10, "integer option, the number to display")
	issueCmd.Flags().IntVarP(&opt.show, "show", "s", 0, "integer option, issue number to show")
	issueCmd.Flags().StringVarP(&opt.add, "add", "a", "", "string option, issue title")
	issueCmd.Flags().IntVarP(&opt.close, "close", "c", 0, "interger option, issue number to delete")
}
