package cmd

import (
    "io"
    "io/ioutil"
    "log"
    "os"
    "text/template"

    "github.com/spf13/cobra"

    "github.com/jitesoft/cc-gen/internal"
    "github.com/jitesoft/cc-gen/internal/cc"
)

var (
    from          string = ""
    createTag     string = ""
    commitUri     string = ""
    shouldPrepend bool   = true
    stdout        bool   = false
)

var generateCmd = &cobra.Command{
    Use:     "generate",
    Aliases: []string{"gen", "create"},
    Short:   "Generate a changelog.",
    Long:    "Generates a changelog in the defined project.",
    Example: "generate 1.2.3 --from=1.0.2 project/",
    Args:    cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        path := args[1]

        var fromCommit string
        if from != "" {
            fromTag, err := internal.GetTag(path, from)
            if err != nil {
                return err
            }
            fromCommit = fromTag.Commit
        } else {
            fromTag, err := internal.GetLastTag(path)
            if err != nil || fromTag == nil {
                cm, err := internal.GetInitialCommit(path)
                if err != nil {
                    return err
                }
                fromCommit = cm.Hash
            } else {
                fromCommit = fromTag.Commit
            }
        }

        commits, err := internal.GetCommits(path, fromCommit)

        if err != nil {
            return err
        }

        log.Printf("Count %d", len(commits))
        // Filter
        c := cc.FilterCommits(commits)

        final := cc.Order(cc.ExtractAll(c))

        var writer io.Writer
        content := []byte("")
        if stdout == true {
            writer = os.Stdout
        } else {
            // Stat file to check if it exist.
            stat, err := os.Stat("CHANGELOG.md")
            fileExists := false
            if stat != nil {
                fileExists = true
            }

            if fileExists && shouldPrepend == false {
                writer, err = os.OpenFile("CHANGELOG" + string(args[0]) + ".md", os.O_WRONLY|os.O_CREATE, 0600)
            } else {
                // Store the current content in memory.
                if fileExists {

                    content, err = ioutil.ReadFile("CHANGELOG.md")
                    if err != nil {
                        return err
                    }
                }

                writer, err = os.OpenFile("CHANGELOG.md", os.O_RDWR | os.O_CREATE, 0600)
            }
            if err != nil {
                return err
            }
        }

        outputTmpl.Execute(writer, templateData{
            Commits:   final,
            Tag:       args[0],
            CommitUri: commitUri,
            Extra:     string(content),
        })

        return nil
    },
}

type templateData struct {
    Commits   map[string][]*cc.ConventionalCommit
    Tag       string
    CommitUri string
    Extra     string
}

var outputTmpl *template.Template = template.Must(template.New("markdown").Parse(`# {{ .Tag }}
{{ range $key, $value := .Commits }}
{{- if $value }}
## {{ $key }}
{{ range $value }}
  * {{ if $.CommitUri }}[{{- .Hash }}]({{ $.CommitUri }}{{ .Hash }}){{ else }}[{{- .Hash }}]{{ end }} {{ .Header }} <small>By: {{ .Author }} @ {{ .Time }}</small>
{{- end }}
{{ end }}
{{- end }}
{{ .Extra }}
`))

func init() {
    generateCmd.Flags().StringVar(
        &from,
        "from",
        "",
        "Specific tag to use as start tag for changelog (defaults to latest none-pre release)",
    )
    generateCmd.Flags().StringVar(
        &createTag,
        "tag",
        "",
        "If the generator should also tag the new release with given value (default to empty string)",
    )
    generateCmd.Flags().BoolVar(
        &shouldPrepend,
        "prepend",
        true,
        "If true and a CHANGELOG.md file exists, the tool will prepend to the file, else it will create a new (CHANGELOG_VERSION.md) (default true)",
    )
    generateCmd.Flags().BoolVar(
        &stdout,
        "stdout",
        false,
        "Setting this to true will print the changelog to stdout instead of a file (default false)",
    )
    generateCmd.Flags().StringVarP(
        &commitUri,
        "commit-uri",
        "c",
        "",
        "Sets base URI to append commit sha to",
    )

    rootCmd.AddCommand(generateCmd)
}
