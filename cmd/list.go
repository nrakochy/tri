/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/nrakochy/tri/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt bool
)

func runList(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	sort.Sort(todo.ByPri(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
	  if allOpt || i.Done == doneOpt {
		msg := i.Label() + "\t" +  i.PrettyDone() + "\t" + i.PrettyP() + "\t" + i.Text + "\t"
		fmt.Fprintln(w, msg)
	  }
	}
	w.Flush()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	 listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	 listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
}
