package cmd

import "github.com/spf13/cobra"

var Version string

var rootCmd = &cobra.Command{
    Use:   "cc-gen",
    Short: "cc-gen is a change log generator for conventional commits.",
    Long:  `cc-gen is a change log generator for conventional commits.
            It scans the directory where it runs for a git repository and
            uses the tags and commits to generate a changelog based on
            the conventional commit standard.`,
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func Execute() error {
    return rootCmd.Execute()
}
