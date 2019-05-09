package main

import (
	"fmt"
	"github.com/homedm/give/commands"
	"os"
	"os/exec"
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

func main() {
	commands.Run()
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
		fmt.Errorf("Error: %s", cmd)
	}
	return string(result), err
}

func init() {
	// get user information
}
