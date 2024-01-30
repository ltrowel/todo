/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/ltrowel/todo/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneOpt bool
var allOpt bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in your To-Do list",
	Long:  `Cleanly display the items you have in your To-Do list`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := list.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(list.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
		}
	}

	w.Flush()

}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Items")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show All Items")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
