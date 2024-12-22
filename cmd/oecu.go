/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kasshii28/go-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// prokenCmd represents the proken command
var oecuCmd = &cobra.Command{
	Use:   "oecu",
	Short: "Print the ascii art of cat",
	Run: func(cmd *cobra.Command, args []string) {
		selectText()
	},
}

func init() {
	rootCmd.AddCommand(oecuCmd)
}

func selectText() {
	text := promptui.Select{
		Label: "What do you want to print Text?",
		Items: []string{"Oecu", "Logo", "Gakutyou"},
		CursorPos: 0,
	}

	idx, res, err := text.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose no.%d %q\n", idx, res)

	switch res {
		case "Oecu":
			utils.PrintAaFromText("oecu.txt")
		case "Logo":
			utils.PrintAaFromText("logo.txt")
		case "Gakutyou":
			utils.PrintAaFromText("gakutyou.txt")
	}
}