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

var isCmd bool
var runCmd bool

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
		ask.Ask(context.Background(), ask.AskParams{
			Query: strings.Join(args, " "),
			IsCmd: isCmd,
			Run:   runCmd,
		})
	},
}

func init() {
	askCmd.PersistentFlags().BoolVarP(&isCmd, "cmd", "c", false, "Ask for command")
	askCmd.PersistentFlags().BoolVarP(&runCmd, "run", "r", false, "Run the asked command")

	rootCmd.AddCommand(askCmd)

}
