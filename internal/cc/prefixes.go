package cc

var prefixes map[string]string = map[string]string{
    "fix":      "Bugfixes",
    "feat":     "Features",
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
