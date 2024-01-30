/*
Copyright © 2024 LIAM TROWEL <trowel.liam@gmail.com>
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
}
