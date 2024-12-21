/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"github.com/kasshii28/go-cli/aa"
	"github.com/spf13/cobra"
)

// prokenCmd represents the proken command
var prokenCmd = &cobra.Command{
	Use:   "proken",
	Short: "Print the ascii art of cat",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := aa.Aa.Open("proken.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Print(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(prokenCmd)
}
