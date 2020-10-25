package cc

import (
    "regexp"
    "time"

    "github.com/jitesoft/cc-gen/internal"
)

// ConventionalCommit is a structure which contains the commit split up in its different parts.
type ConventionalCommit struct {
    Type      string
    NamedType string
    SubType   string
    Header    string
    Body      string
    Hash      string
    Author    string
    Time      string
}

var extractor = regexp.MustCompile(`^(\w+)[(]?(\w+)?[)]?[:]\s(.*)[\n](.*)`)

// Extract parses a commit and extracts the different parts into a new structure.
func Extract(commit *internal.Commit) ConventionalCommit {
    extracted := extractor.FindAllStringSubmatch(commit.Message, -1)

    return ConventionalCommit{
        Type:      extracted[0][1],
        NamedType: getTypeName(extracted[0][1]),
        SubType:   extracted[0][2],
        Header:    extracted[0][3],
        Body:      extracted[0][4],
        Hash:      commit.Hash,
        Author:    commit.Author,
        Time:      commit.Time.UTC().Format(time.RFC3339),
    }
}
