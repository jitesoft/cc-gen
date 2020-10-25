package internal

import (
    "log"
    "time"

    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/object"
)

// Commit is a single commit in the repository.
type Commit struct {
    Author  string
    Hash    string
    Message string
    Time    time.Time
}

var commits []*Commit

// GetInitialCommit fetches the first commit in the tree.
func GetInitialCommit(path string) (*Commit, error) {
    if err := fillUpCommits(path, true); err != nil {
        return nil, err
    }

    return commits[len(commits)-1], nil
}

func fillUpCommits(path string, force bool) error {

    if force {
        commits = []*Commit{}
    }

    if len(commits) == 0 {
        log.Print("no commits in memory, fetching all commits")

        repo := getRepository(path)
        headRef, _ := repo.Head()

        cIter, err := repo.Log(&git.LogOptions{
            From: headRef.Hash(),
            All:  true,
        })

        if err != nil {
            return err
        }

        // Store the commits, we will most likely use them soon again.
        _ = cIter.ForEach(func(commit *object.Commit) error {
            commits = append(commits, &Commit{
                Author:  commit.Author.Name,
                Time:    commit.Author.When,
                Hash:    commit.Hash.String(),
                Message: commit.Message,
            })
            return nil
        })
    } else {
        log.Print("commits already in memory")
    }


    return nil
}

// GetCommits returns a series of commits from a repository.
// The span will go from `HEAD` until `toHash` value (which should be a sha hash).
func GetCommits(repoPath string, toHash string) ([]*Commit, error) {
    if err := fillUpCommits(repoPath, false); err != nil {
        return nil, err
    }

    log.Printf("trying to find hash (%s)...", toHash)
    var out []*Commit
    for _, c := range commits {
        if c.Hash == toHash {
            log.Printf("Hash found %s.", c.Hash)
            break
        }

        out = append(out, c)
    }

    log.Printf("Total commits between current (head) and %s: %d", toHash, len(out))

    return out, nil
}
