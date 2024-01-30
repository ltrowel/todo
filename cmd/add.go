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

var priority int

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
		item := list.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = list.SaveItems(viper.GetString("datafile"), items)

	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
