package cc

import (
    "strings"

    "github.com/jitesoft/cc-gen/internal"
)

// FilterCommits filters all commits in the slice and returns a new slice
// with only commits which starts with a CC-prefix.
func FilterCommits(commits []*internal.Commit) []*internal.Commit {
    var ccList []*internal.Commit
    prefixes := getPrefixes()
    for _, c := range commits {
        if startsWithAny(prefixes, c.Message) {
            ccList = append(ccList, c)
        }
    }

    return ccList
}

func startsWithAny(arr []string, str string) bool {
    for _, a := range arr {
        if strings.HasPrefix(str, a) {
            return true
        }
    }
    return false
}
