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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new To-Do item",
	Long:  `Add will create a new item to your To-Do list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := list.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := list.Item{ID: len(items) + 1, Text: x}
		items = append(items, item)
	}

	err = list.SaveItems(viper.GetString("datafile"), items)

	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
