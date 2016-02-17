package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/github"
)

var (
	Account string
	DstDir  string
	Action  string
)

func init() {
	flag.StringVar(&Account, "acc", "docker", "github account")
	flag.StringVar(&DstDir, "dst", "E:/golib/src/github.com/", "github.com code directory")
	flag.StringVar(&Action, "action", "pull", "clone/pull")
	flag.Parse()
}

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(Account, opt)
	if err != nil {
		fmt.Printf("error : %s\n", err)
		return
	}
	for _, repo := range repos {
		cloneUrl := fmt.Sprintf("https://github.com/%s/%s", Account, *repo.Name)
		var stdout, stderr bytes.Buffer
		var cmd *exec.Cmd
		switch Action {
		case "clone":
			cmd = exec.Command("git", "clone", cloneUrl)
			cmd.Dir = filepath.Join(DstDir, Account)
		case "pull":
			workdir := filepath.Join(DstDir, Account, *repo.Name)
			if _, err := os.Stat(workdir); os.IsNotExist(err) {
				// path/to/whatever does not exist
				continue
			}
			cmd = exec.Command("git", "pull")
			cmd.Dir = workdir
		}
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Start(); err != nil {
			fmt.Printf("git %s %s/%s failed, %s\n", Action, Account, *repo.Name, err)
			continue
		}
		if err := cmd.Wait(); err != nil {
			fmt.Printf("git %s %s/%s failed, %s\n", Action, Account, *repo.Name, err)
			continue
		}
		fmt.Printf("%s/%s stdout:%s, stderr:%s", Account, *repo.Name, stdout.String(), stderr.String())
	}
}
