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

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Mark an Item as Done",
	Aliases: []string{"do"},
	Long:    `Mark a To-Do Item as completed`,
	Run:     checkRun,
}

func checkRun(cmd *cobra.Command, args []string) {
	filename := viper.GetString("datafile")
	position, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(position, "is not a valid label\n", err)
	}

	err = list.ToggleChecked(filename, position, true)

	if err != nil {
		log.Fatalln(err, position)
	}
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
