/*
Copyright © 2024 LIAM TROWEL <trowel.liam@gmail.com>
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
}
