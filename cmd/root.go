/*
Copyright © 2023 Digger.dev
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type Backend struct {
	name   string
	bucket string
}

const metadataTableName = "digger-tbm-metadata"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tbm",
	Short: "Terraform Backend Manager",
	Long:  `TBM is a CLI to manage S3 backends for Terraform on AWS`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tbm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
