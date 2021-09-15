package conventional

import (
	"fmt"
	"github.com/jitesoft/cc-gen/gitwrapper"
	"regexp"
	"strings"
	"time"
)

// Commit is a struct which contains parsed commit messages following
// the conventional commit standard.
type Commit struct {
	Type string
	SubType string
	Header string
	Message string
	Hash string
	Author string
	Time time.Time
}

// Regexp does...
// First group: (\w+) - Type (feat:, refactor:, test: etc).
// Second group, optional (\w+) - SubType (feat(subtype):)
// Required `:\s`
// Third group (.*) - Header.
// 1 or more newlines.
// Fourth group (.*) - Body (?s) - as single line (so it captures multi-line)
var extractor = regexp.MustCompile(`^(\w+)[(]?(\w+)?[)]?[:]\s(.*)([\n]+(?s)(.*))?`)

// IsConventional tests if a commit is a conventional commit.
func IsConventional(c *gitwrapper.Commit) bool {
	return extractor.Match([]byte(c.Message))
}

// ParseConventional takes a commit and parses it into the ConventionalCommit
// structure, ready for usage.
func ParseConventional(c *gitwrapper.Commit) (*Commit, error) {
	if !IsConventional(c) {
		return nil, fmt.Errorf(
			"failed to parse commit, commit is not in the correct format",
		)
	}

	extracted := extractor.FindAllStringSubmatch(c.Message, -1)

	return &Commit{
		Type:    extracted[0][1],
		SubType: extracted[0][2],
		Header:  extracted[0][3],
		Message: strings.TrimSpace(extracted[0][4]),
		Hash:    c.Hash,
		Author:  c.Author,
		Time:    c.Time,
	}, nil
}

// GroupByType groups a list of conventional commits into a map
// in which the key is the Commit.Type and the value is the Commit.
func GroupByType(commits []*Commit) map[string][]*Commit {
	var out = map[string][]*Commit{}

	for _, c := range commits {
		if _, found := out[c.Type]; !found {
			out[c.Type] = []*Commit{}
		}

		out[c.Type] = append(out[c.Type], c)
	}

	return out
}
