package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	// remove timestamp from log
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

// Execute executes kubefun.
func Execute(version string, gitCommit string, buildTime string) {
	rootCmd := newRoot(version, gitCommit, buildTime)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRoot(version string, gitCommit string, buildTime string) *cobra.Command {
	rootCmd := newKubefunCmd(version)
	rootCmd.AddCommand(newVersionCmd(version, gitCommit, buildTime))

	return rootCmd
}
