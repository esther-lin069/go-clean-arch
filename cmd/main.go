package main

import (
	"go-clean-arch/app/websocketcilent"
	"os"

	"github.com/spf13/cobra"
)

// root
var rootCmd = &cobra.Command{
	Long: "game core service.",
}

// event
var eventCmd = &cobra.Command{
	Use:       "event websocket-client",
	Short:     "event",
	ValidArgs: []string{},
	Run: func(cmd *cobra.Command, args []string) {
		websocketcilent.Start()
	},
}

func main() {
	// 禁止生成預設的次要指令 `completion`
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// 綁定 eventCmd
	rootCmd.AddCommand(eventCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
