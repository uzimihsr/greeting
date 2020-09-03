package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var globalFlags struct {
	language string
}

var rootCmd = &cobra.Command{
	Use:   "greeting",
	Short: "greeting command",
	Long: `greeting command.
say hello to your friend.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// サブコマンドでも共通で利用できるグローバルな引数(language)の設定
	rootCmd.PersistentFlags().StringVarP(&globalFlags.language, "lang", "l", "en", "Language : en, ja")
}
