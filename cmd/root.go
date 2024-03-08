package cmd

import (
	"github.com/spf13/cobra"
)

var (
	loadCfg bool
	rootCmd = &cobra.Command{
		Use: "app",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().
		BoolVar(&loadCfg, "load-env", false, "load-env in order to load local .env file")
}
