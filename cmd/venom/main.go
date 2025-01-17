package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/luphaz/venom/cmd/venom/run"
	"github.com/luphaz/venom/cmd/venom/update"
	"github.com/luphaz/venom/cmd/venom/version"
)

var rootCmd = &cobra.Command{
	Use:   "venom",
	Short: "Venom aim to create, manage and run your integration tests with efficiency",
}

func main() {
	addCommands()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Err:%s", err)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(run.Cmd)
	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(update.Cmd)
}
