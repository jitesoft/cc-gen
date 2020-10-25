package cmd

import (
    "log"
    "os"

    "github.com/spf13/cobra"

    "github.com/jitesoft/cc-gen/internal"
    "github.com/jitesoft/cc-gen/internal/cc"
)

var (
    from         string = ""
    createTag    string = ""
    shouldAppend bool   = true
    stdout       bool   = false
)

var generateCmd = &cobra.Command{
    Use:     "generate",
    Aliases: []string{"gen", "create"},
    Short:   "Generate a changelog.",
    Long:    "Generates a changelog in the defined project.",
    Example: "generate --from=1.0.2 project/",
    Args:    cobra.MaximumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        var path string
        path, err := os.Getwd()
        if err != nil {
                return err
        }

        if len(args) == 1  {
            path = args[0]
        }

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
        commits = cc.FilterCommits(commits)

        for _, c := range commits {
            log.Printf("Msg: %s", cc.Extract(c).Header)
        }


        return nil
    },
}


func init() {
    generateCmd.Flags().StringVar(&from, "from", "", "Specific tag to use as start tag for changelog (defaults to latest none-pre release)")
    generateCmd.Flags().StringVar(&createTag, "tag", "", "If the generator should also tag the new release (default to empty string)")
    generateCmd.Flags().BoolVar(&shouldAppend, "append", true, "If true and a CHANGELOG.md file exists, the tool will append the file, else it will create a new (CHANGELOG_VERSION.md) (default true)")
    generateCmd.Flags().BoolVar(&stdout, "stdout", false, "Setting this to true will print the changelog to stdout instead of a file (default false)")

    rootCmd.AddCommand(generateCmd)
}
