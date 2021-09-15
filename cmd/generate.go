package cmd

import (
	"github.com/jitesoft/cc-gen/gitwrapper"
	"github.com/jitesoft/cc-gen/gitwrapper/conventional"
	"github.com/jitesoft/cc-gen/internal/template"
	"github.com/spf13/cobra"
	"log"
	"os"
	"io/ioutil"
)

var (
	title string     // Title of the release, defaults to no title.
	commitUri string // Uri to where commits should be linked.
	full bool        // If true, _all_ tags should be generated in the changelog.
	stdout bool      // If output is to go to stdout instead of file.
	prepend bool     // If changes should be prepended in case a changelog already exist.
	untilTag string  // Tag name to stop at, will override 'full'.
	fromTag string   // Tag name to use as start.
)

var generateCmd = &cobra.Command{
	Use:     "generate [tag name]",
	Aliases: []string{ "gen", "build" },
	Short:   "Generate a changelog.",
	Long:    `Generates a changelog in the defined project.
Passing a tag name (optional) will mark the latest as under the given tag.
If the latest commit is tagged, that will be used instead of the tag name argument.
If the latest commit is not tagged and tag name is excluded 'latest' will be used instead.`,
	Example: "cc-gen generate 1.2.3 --title=latest-awesome-release",
	Args:    cobra.MaximumNArgs(1),
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		branch, err := gitwrapper.GetCurrentBranch()
		latestTag := "latest"
		if len(args) > 0 {
			latestTag = args[0]
		}

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
			latestTag = tags[1].Name
		}

		firstHash := branch.Commits[0].Hash // Current commit.
		lastHash  := branch.Commits[len(branch.Commits) - 1].Hash // Last of all!

		if !full {
			if len(fromTag) > 0 {
				firstHash = gitwrapper.FindTag(tags, fromTag).Hash
			}
			if len(untilTag) > 0 {
				lastHash = gitwrapper.FindTag(tags, untilTag).Hash
			}
		}

		tags := gitwrapper.GetTagRangeBetween(tags, firstHash, lastHash)




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

		output := os.Stdout
		extra := ""
		if !stdout {
			if !prepend {
				output, _ = os.Create("CHANGELOG.md")
			} else {
				output, _ = os.OpenFile("CHANGELOG.md", os.O_RDWR | os.O_CREATE, 0755)
				fileData, err := ioutil.ReadFile("CHANGELOG.md")
				if err != nil {
					extra = ""
				} else {
					extra = string(fileData)
				}
			}
		}

		err = template.RenderTemplate(&template.Data{
			CommitUri: "",
			Extra:     extra,
			Tags: []*template.TagData{
				{
					Commits:   conventional.GroupByType(commits),
					Name:      latestTag,
				},
			},
		}, "default", output)

		if err != nil {
			log.Panic(err)
		}



		return nil
	},
}

func init() {
	generateCmd.Flags().StringVar(
		&fromTag,
		"from",
		"",
		"Specific tag to use as start tag for changelog (uses latest commit as default)",
	)
	generateCmd.Flags().StringVar(
		&untilTag,
		"until",
		"",
		"Specific tag to stop at (defaults to latest tag or, in case it shares hash with latest commit, the one before)",
	)
	generateCmd.Flags().BoolVar(
		&full,
		"full",
		false,
		"Setting this to true indicates that a full changelog should be generated instead of just since last tag (defaults to false)",
	)
	generateCmd.Flags().BoolVar(
		&prepend,
		"prepend",
		true,
		"If true and a CHANGELOG.md file exists, the tool will prepend to the file, else it will create or replace the CHANGELOG.md file (default true)",
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
