package internal

import (
    "log"

    "github.com/go-git/go-git/v5"
)

func getRepository(path string) *git.Repository {
    repo, err := git.PlainOpen(path)
    if err != nil {
        log.Panicf("failed to open the repository %s", path)
    }
    return repo
}
