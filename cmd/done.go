/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/nrakochy/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
	  log.Fatalln("Problem reading ye olde file\n", err)
        }

	i, err := strconv.Atoi(args[0])

	if err != nil {
	  log.Fatalln(args[0], "is not a valid label\n", err)
	  return
        }

	if i > 0 && i < len(items){
		items[i - 1].Done = true
		fmt.Printf("%q %v \n", items[i-1].Text, "marked done")
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "does not match any items")
	}
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
