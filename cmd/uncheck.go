/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/ltrowel/todo/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uncheckCmd represents the uncheck command
var uncheckCmd = &cobra.Command{
	Use:   "uncheck",
	Short: "Mark an item as incomplete",
	Long:  `Move a To-Do item back to not Done`,
	Run:   uncheckRun,
}

func uncheckRun(cmd *cobra.Command, args []string) {
	filename := viper.GetString("datafile")
	position, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(position, "is not a valid label\n", err)
	}

	err = list.ToggleChecked(filename, position, false)

	if err != nil {
		log.Fatalln(err, position)
	}
}

func init() {
	rootCmd.AddCommand(uncheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uncheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uncheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
