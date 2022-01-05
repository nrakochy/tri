/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/nrakochy/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int


func addRun(cmd *cobra.Command, args []string){
  items, err := todo.ReadItems(viper.GetString("datafile"))
  if err != nil {
	  log.Printf("%v", err)
  }
  for _, x := range args {
	  item := todo.Item{Text: x}
	  item.SetPriority(priority)
	  items = append(items, item)
  }

 writeErr := todo.SaveItems(viper.GetString("datafile"), items);
 if writeErr != nil {
	 fmt.Errorf("%v", err)
 }
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addRun,
}


func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	 addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
}
