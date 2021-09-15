package cmd

import (
	"github.com/jitesoft/cc-gen/gitwrapper"
	"github.com/jitesoft/cc-gen/gitwrapper/conventional"
	"github.com/jitesoft/cc-gen/internal/template"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	from          string = ""
	commitUri     string = ""
	shouldPrepend bool   = true
	stdout        bool   = false
	extTemplate   string = ""
)

var generateCmd = &cobra.Command{
	Use:     "cc-gen",
	Short:   "Generate a changelog.",
	Long:    "Generates a changelog in the defined project.",
	// Example: "cc-gen 1.2.3 --from=1.0.2 project/",
	Args:    cobra.ExactArgs(2),
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		branch, err := gitwrapper.GetCurrentBranch()

		if err != nil {
			log.Panic(err)
		}

		tags, err := gitwrapper.GetTags()

		if err != nil {
			log.Panic(err)
		}

		tag := tags[0]
		if branch.Commits[0].Hash == tags[0].Hash {
			tag = tags[1]
		}

		var commits []*conventional.Commit
		for _, c := range branch.Commits {
			if c.Hash == tag.Hash {
				break
			}

			if conventional.IsConventional(c) {
				con, _ := conventional.ParseConventional(c)
				commits = append(commits, con)
			}
		}

		template.RenderTemplate(&template.Data{
			Commits:   conventional.GroupByType(commits),
			Tag:       "new tag!",
		}, "default", os.Stdout)

		return nil
	},
}

func init() {
	generateCmd.Flags().StringVar(
		&from,
		"from",
		"",
		"Specific tag to use as start tag for changelog (defaults to latest none-pre release)",
	)
	generateCmd.Flags().StringVar(
		&extTemplate,
		"template",
		"",
		"Path to template to use instead of default",
	)
	generateCmd.Flags().BoolVar(
		&shouldPrepend,
		"prepend",
		true,
		`If true and a CHANGELOG.md file exists, the tool will prepend to the file, 
               else it will create or replace the CHANGELOG.md file (default true)`,
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
