/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"rwf/pkg/raider"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info (expansions|info)",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:       info,
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
	ValidArgs: []string{"expansions", "raids"},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().IntP("expansion", "e", raider.CurrentExpansion, "Get information for an expansion")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func info(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "expansions":
		fmt.Println("Expansion information")
		// Loop over raider.Expansion and print out information
		break
	case "raids":
		fmt.Println("Raid information")
		break
	}
}
