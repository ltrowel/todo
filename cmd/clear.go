/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/ltrowel/todo/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear out your todo list",
	Long:  `Remove all items from todo list`,
	Run:   clearRun,
}

func clearRun(cmd *cobra.Command, args []string) {
	err := list.ClearItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Cleared To-Do list!")
	}
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
