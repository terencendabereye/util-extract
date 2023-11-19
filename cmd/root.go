/*
Copyright © 2023 Terence Ndabereye ndabereye@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "extract [flags] filename...",
	Short: "uncompress compressed files",
	Example: "extract file.zip",
	// Args: cobra.ExactArgs(1),
	Args:  cobra.MinimumNArgs(1),
	ValidArgs: []string{"filename ..."},
	Run: func(cmd *cobra.Command, args []string) {
		runExtract(args)
	},
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
	rootCmd.Version = AppVersion
	
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.extract.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
