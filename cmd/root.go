package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:     "cc-gen",
    Short:   "Generate a changelog from conventional commits.",
    Long:    `
Generates a changelog from conventional commits. 

Depending on flags this application create a changelog spanning between the last two releases (skipping pre-release) 
and output it in a markdown file or to stdout.`,
    Version: "0.0.1",
}

func init() {

}

func Execute() error {
    return rootCmd.Execute()
}
