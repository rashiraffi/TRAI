/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"strings"

	"github.com/rashiraffi/trai/internal/ask"

	"github.com/spf13/cobra"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   `ask "question"`,
	Short: "Ask any question and get an AI-powered response",
	Long: `The 'ask' command within the 'trai' application enables users to ask any question and receive an AI-powered response.
This command seamlessly integrates AI capabilities into the terminal environment, allowing users to obtain information or
assistance on various topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		ask.Ask(context.Background(), strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
