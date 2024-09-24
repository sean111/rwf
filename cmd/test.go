/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"rwf/pkg/raider"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := raider.GetStaticData(10)
		if err != nil {
			log.Error("Error: ", err)
		}
		log.Info("Raid Data :", "name", data.Raid[0].Name, "slug", data.Raid[0].Slug, "start_time", data.Raid[0].StartTimes.US)
		//log.Info("Slug :", currentRaid.Slug)
		//log.Info("US Start Time :", currentRaid.StartTimes.US)
		//log.Info("US End Time :", currentRaid.EndTimes.US)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
