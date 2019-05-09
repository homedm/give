package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command{
		Use: "give",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// User is your github account information
type User struct {
	Token string
	Name  string
}

// Repo is repository information
type Repo struct {
	Owner string
	URL   string
}

var (
	usr  User
	repo Repo
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

func getGitConfig(opt string) (out string, err error) {
	cmd := exec.Command("git", "config", opt)
	result, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Error: %s", cmd)
	}
	return strings.TrimSpace(string(result)), err
}

func checkAPILimit() error {
	return nil
}

func init() {
	// get user information
	var err error
	usr.Token, err = getGitUsrToken()
	if err != nil {
		fmt.Errorf("Error: can not get user token")
	}

	fmt.Println("initialization finished")
}
