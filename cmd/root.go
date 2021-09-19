package cmd

import "github.com/spf13/cobra"

var Version string = "unknown"

var rootCmd = &cobra.Command{
	Use:     "cc-gen",
	Version: Version,
	Short:   "cc-gen is a change log generator for conventional commits.",
	Long: `cc-gen is a change log generator for conventional commits.
It scans the directory where it runs for a git repository and
uses the tags and commits to generate a changelog based on
the conventional commit standard.`,
}

func Execute() error {
	return rootCmd.Execute()
}
