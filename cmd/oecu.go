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
var oecuCmd = &cobra.Command{
	Use:   "oecu",
	Short: "Print the ascii art of cat",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := aa.AaOecu.Open("oecu.txt")
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
	rootCmd.AddCommand(oecuCmd)
}
