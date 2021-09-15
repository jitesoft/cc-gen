package cmd
/*
import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initLong = `Initializes a configuration file for cc-gen in current directory,
the file is used by cc-gen to store repository configuration.
A configuration file is not required, arguments can be used to customize instead.`

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initializes a configuration file for cc-gen in current directory",
	Long:  initLong,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()

		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("initializing cc-gen in directory %s", path)
	},
}
*/
