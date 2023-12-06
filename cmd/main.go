package main

import (
	"go-clean-arch/app/websocketcilent"
	"go-clean-arch/app/websocketserver"
	"os"

	"github.com/spf13/cobra"
)

// root
var rootCmd = &cobra.Command{
	Long: "start service.",
}

// websocket-client
var websocketClientCmd = &cobra.Command{
	Use:       "websocket-client",
	Short:     "websocket-client",
	ValidArgs: []string{},
	Run: func(cmd *cobra.Command, args []string) {
		websocketcilent.Start()
	},
}

// websocket-server
var websocketServerCmd = &cobra.Command{
	Use:       "websocket-server",
	Short:     "websocket-server",
	ValidArgs: []string{},
	Run: func(cmd *cobra.Command, args []string) {
		websocketserver.Start()
	},
}

func main() {
	// 禁止生成預設的次要指令 `completion`
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// 綁定 eventCmd
	rootCmd.AddCommand(websocketClientCmd)
	rootCmd.AddCommand(websocketServerCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
