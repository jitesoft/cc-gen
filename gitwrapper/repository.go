package gitwrapper

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var (
	repo *git.Repository
)

func GetCurrentBranch(path string) (*Branch, error) {
	if repo == nil {
		repository, err := git.PlainOpen(path)

		if err != nil {
			return nil, fmt.Errorf("failed to open repository: %s", err)
		}
		repo = repository
	}

	headRef, err := repo.Head()

	if err != nil {
		return nil, err
	}

	cIter, err := repo.Log(&git.LogOptions{
		From: headRef.Hash(),
		All:  false,
	})

	branch := new(Branch)

	_ = cIter.ForEach(func(commit *object.Commit) error {
		branch.Commits = append(branch.Commits, &Commit{
			Hash:    commit.Hash.String(),
			Message: commit.Message,
			Time:    commit.Author.When,
			Author:  commit.Author.Name,
			Email:   commit.Author.Email,
		})

		return nil
	})

	return branch, nil
}

// GetBranch uses the internal go-git library to check out a branch
// and add all the commits in said branch as Commit objects.
func GetBranch(branchName string, path string) (*Branch, error) {
	if repo == nil {
		repository, err := git.PlainOpen(path)

		if err != nil {
			return nil, fmt.Errorf("failed to open repository: %s", err)
		}
		repo = repository
	}

	tree, err := repo.Worktree()

	if err != nil {
		return nil, err
	}

	err = tree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branchName),
	})

	if err != nil {
		return nil, err
	}

	return GetCurrentBranch(path)
}

// GetTags parses repository tags and returns a list of Tag objects.
func GetTags(path string) ([]*Tag, error) {
	if repo == nil {
		repository, err := git.PlainOpen(path)

		if err != nil {
			return nil, fmt.Errorf("failed to open repository: %s", err)
		}
		repo = repository
	}

	var result []*Tag

	tIter, _ := repo.TagObjects()
	_ = tIter.ForEach(func(tag *object.Tag) error {
		result = append(result, &Tag{
			Name:    tag.Name,
			Hash:    tag.Target.String(),
			TagHash: tag.Hash.String(),
			Tagger:  tag.Tagger.Name,
			Time:    tag.Tagger.When,
			Email:   tag.Tagger.Email,
			Message: tag.Message,
		})
		return nil
	})

	sortTagsBy(func(t1 *Tag, t2 *Tag) bool {
		return t1.Time.After(t2.Time)
	}).Sort(result)

	return result, nil
}
