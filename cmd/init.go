/*
Copyright © 2023 Digger.dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize TBM metadata on AWS",
	Long:  `TBM creates a DynamoDB table to store S3 bucket IDs of backends and other metadata`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stub for the init command")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
