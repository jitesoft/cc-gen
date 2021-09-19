package cmd

import (
	"github.com/jitesoft/cc-gen/gitwrapper"
	"github.com/jitesoft/cc-gen/gitwrapper/conventional"
	"github.com/jitesoft/cc-gen/internal/template"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var (
	commitUri string // Uri to where commits should be linked.
	full      bool   // If true, _all_ tags should be generated in the changelog.
	stdout    bool   // If output is to go to stdout instead of file.
	prepend   bool   // If changes should be prepended in case a changelog already exist.
	untilTag  string // Tag name to stop at, will override 'full'.
	fromTag   string // Tag name to use as start.
	fileName  string // Name of file, if not stdout.
)

var generateCmd = &cobra.Command{
	Use:     "generate [tag name]",
	Aliases: []string{"gen", "build"},
	Short:   "Generate a changelog.",
	Long: `Generates a changelog in the current project.
Passing a tag name (optional) will mark the latest as under the given tag.
If the latest commit is tagged, that will be used instead of the tag name argument.
If the latest commit is not tagged and tag name is excluded 'latest' will be used instead.`,
	Example: "cc-gen generate 1.2.3",
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

		if branch.Commits[0].Hash == tags[0].Hash {
			latestTag = tags[0].Name
		}

		firstHash := branch.Commits[0].Hash
		lastHash := branch.Commits[len(branch.Commits)-1].Hash

		if !full {
			if len(fromTag) > 0 {
				t := gitwrapper.FindTag(tags, fromTag)
				latestTag = t.Name
				firstHash = t.Hash
			}

			if len(untilTag) > 0 {
				lastHash = gitwrapper.FindTag(tags, untilTag).Hash
			}
		}

		tagData := make(map[string][]*conventional.Commit)
		currentTag := latestTag
		hasFirst := false
		tagData[currentTag] = []*conventional.Commit{}
		for _, c := range branch.Commits {
			// Check if the hash is a tag, if it is, we want to use another 'current tag'.
			// Have to be done even if we haven't reached the first hash.
			if tag := gitwrapper.FindTagByHash(tags, c.Hash); tag != nil {
				currentTag = tag.Name
			}

			// Find first hash or continue.
			if !hasFirst && c.Hash != firstHash {
				continue
			}
			hasFirst = true

			// Last hash? then break!
			if c.Hash == lastHash {
				break
			}

			// If the commit is conventional, we add it to the list.
			if conventional.IsConventional(c) {
				cc, _ := conventional.ParseConventional(c)
				// currentTag is the tag name, we add the commit to its array.
				tagData[currentTag] = append(tagData[currentTag], cc)
			}
		}

		// After collecting all commits, we have to create the tag data and group
		// the commits under their 'type'. Can't be done above, as we don't have
		// all in the list at that point.
		var outputData []*template.TagData
		for t, c := range tagData {
			outputData = append(outputData, &template.TagData{
				Commits: conventional.GroupByType(c),
				Name:    t,
			})
		}

		output := os.Stdout
		extra := ""
		if !stdout {
			if !prepend {
				output, _ = os.Create(fileName)
			} else {
				output, _ = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
				fileData, err := ioutil.ReadFile(fileName)
				if err != nil {
					extra = ""
				} else {
					extra = string(fileData)
				}
			}
		}

		err = template.RenderTemplate(&template.Data{
			CommitUri: commitUri,
			Extra:     extra,
			Tags:      outputData,
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
		"to",
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
		"If true and a [output] file exists, the tool will prepend to the file, else it will create or replace the [output] file (default true)",
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
		"Sets base URI to use as link to commits. Uses sprintf format with a single %s where hash will be inserted",
	)
	generateCmd.Flags().StringVarP(
		&fileName,
		"output",
		"o",
		"CHANGELOG.md",
		"Filename to write the generated changelog to. Not used in `stdout` mode (defaults to CHANGELOG.md)",
	)

	rootCmd.AddCommand(generateCmd)
}
