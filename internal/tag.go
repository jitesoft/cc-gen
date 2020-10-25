package internal

import (
    "fmt"
    "log"

    "github.com/blang/semver/v4"
    "github.com/go-git/go-git/v5/plumbing"
)

type Tag struct {
    Name   string
    Commit string
}

func GetLastTag(repoPath string) (*Tag, error) {
    repo := getRepository(repoPath)
    head, err := repo.Head()
    if err != nil {
        return nil, fmt.Errorf("failed to find head")
    }

    allTags, err := repo.Tags()
    if err != nil || allTags == nil {
        log.Printf("failed to fetch tags")
        return nil, err
    }

    var tag *plumbing.Reference = nil
    _ = allTags.ForEach(func(reference *plumbing.Reference) error {
        // If the tag hash is actually the current head, we just skip it.
        if reference.Hash().String() != head.Hash().String() {
            if tag == nil {
                tag = reference
            } else {
                v1, err := semver.Make(tag.Name().Short())
                v2, err2 := semver.Make(reference.Name().Short())

                if err != nil || err2 != nil {
                    log.Printf("error while fetching latest versions. A version could not be parsed (%s - %s)", tag.Name().Short(), reference.Name().Short())
                } else {
                    if v2.Compare(v1) > 0 {
                        tag = reference
                    }
                }
            }
        }
        return nil
    })

    if tag == nil {
        return nil, fmt.Errorf("failed to find a tag")
    }

    return &Tag{
        Name:   tag.Name().String(),
        Commit: tag.Hash().String(),
    }, nil
}

func GetTag(repoPath string, name string) (*Tag, error) {
    repo := getRepository(repoPath)

    tag, err := repo.Tag(name)
    if err != nil {
        return nil, err
    }

    if tag == nil {
        log.Printf("tag %s does not exist in repository %s", name, repoPath)
    }

    // Try fetch annotated tag.
    cHash := ""
    if t, err := repo.TagObject(tag.Hash()); err == nil {
        commit, err := t.Commit()
        if err == nil {
            cHash = commit.Hash.String()
        }
    }

    if cHash == "" {
        cHash = tag.Hash().String()
    }

    return &Tag{
        Name:   tag.Name().Short(),
        Commit: cHash,
    }, nil
}
