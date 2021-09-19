package gitwrapper

import (
	"time"
)

type Commit struct {
	Hash    string
	Message string
	Time    time.Time
	Author  string
	Email   string
}

func GetCommitRangeFrom(commits []*Commit, from string) []*Commit {
	for i, c := range commits {
		if c.Hash == from {
			// We can sub-slice here as we won't modify the commit either way.
			return commits[i:]
		}
	}
	return commits
}

func GetCommitRangeTo(commits []*Commit, from string) []*Commit {
	for i, c := range commits {
		if c.Hash == from {
			// We can sub-slice here as we won't modify the commit either way.
			return commits[:i+1]
		}
	}
	return commits
}

func GetCommitRangeBetween(commits []*Commit, from string, to string) []*Commit {
	j, k := 0, 0

	for i, c := range commits {
		if c.Hash == from {
			j = i
		}
		if c.Hash == to {
			k = i + 1
			break
		}
	}

	return commits[j:k]
}
