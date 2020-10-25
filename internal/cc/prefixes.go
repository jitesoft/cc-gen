package cc

var prefixes map[string]string = map[string]string{
    "feat":     "Features",
    "fix":      "Bugfixes",
    "misc":     "Misc",
    "style":    "Styling",
    "docs":     "Documentation",
    "ci":       "CI/CD",
    "refactor": "Refactoring",
    "test":     "Testing",
}

func SetPrefixes(pref map[string]string) {
    prefixes = pref
}

func Order (commits []*ConventionalCommit) map[string][]*ConventionalCommit {
    out := map[string][]*ConventionalCommit{}
    for _, name := range getPrefixes() {
        out[getTypeName(name)] = []*ConventionalCommit{}
    }

    for _, c := range commits {
        out[c.NamedType] = append(out[c.NamedType], c)
    }
    return out
}

func getPrefixes() []string {
    var result []string
    for name, _ := range prefixes {
        result = append(result, name)
    }

    return result
}

func getTypeName(t string) string {
    return prefixes[t]
}
