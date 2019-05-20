package commands

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command{
		Use:   "give",
		Short: "GitHub Viewer command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

var (
	ctx    = context.Background()
	client *github.Client

	pattern = regexp.MustCompile(`^(?:(?:ssh://)?git@github\.com(?::|/)|https://github\.com/)([^/]+)/([^/]+?)(?:\.git)?$`)
)

// Run runs command.
func Run() {
	RootCmd.Execute()
}

// Exit finishes a running action.
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func getGitUsrName() (name string, err error) {
	return getGitConfig("user.name")
}

func getGitUsrToken() (token string, err error) {
	return getGitConfig("give.token")
}

func getGitRepoOwner() (owner string, err error) {
	url, err := getGitRepoURL()
	if err != nil {
		return "", nil
	}

	matches := pattern.FindStringSubmatch(url)
	if len(matches) != 3 {
		return "", fmt.Errorf("can not parse remote.origin.url")
	}

	owner = matches[1]
	return
}

func getGitRepoName() (name string, err error) {
	url, err := getGitRepoURL()
	if err != nil {
		return "", nil
	}

	matches := pattern.FindStringSubmatch(url)
	if len(matches) != 3 {
		err = fmt.Errorf("can not parse remote.origin.url")
		return
	}

	name = matches[2]
	return
}

func getGitRepoURL() (url string, err error) {
	return getGitConfig("remote.origin.url")
}

func getGitEditor() (editor string, err error) {
	return getGitConfig("core.editor")
}

func getGitConfig(opt string) (out string, err error) {
	cmd := exec.Command("git", "config", opt)
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func checkAPILimit(usr string) error {
	_, _, err := client.Repositories.List(ctx, usr, nil)
	if _, ok := err.(*github.RateLimitError); ok {
		return fmt.Errorf("hit rate limit: %s", err)
	}
	return nil
}

func init() {
	token, err := getGitUsrToken()
	if err != nil {
		return
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}
