package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var goodbyeCmd = &cobra.Command{
	Use:   "goodbye [NAME]",
	Args:  cobra.MinimumNArgs(1), // 引数が1つ未満の場合はエラー
	Short: "goodbye command",
	Long: `goodbye command.
say goodbye to your friends!`,
	Run: func(cmd *cobra.Command, args []string) {
		switch globalFlags.language {
		case "ja":
			sayGoodbyeInJapanese(args)
		case "en":
			sayGoodbyeInEnglish(args)
		default:
			fmt.Println("please specify the language (en/ja).")
			os.Exit(1)
		}
		fmt.Println(helloFlags.message)
	},
}

func sayGoodbyeInJapanese(args []string) {
	for _, name := range args {
		fmt.Printf("%sさん, ", name)
	}
	fmt.Println("さようなら!")
}

func sayGoodbyeInEnglish(args []string) {
	fmt.Print("Goodbye,")
	for i, name := range args {
		fmt.Print(" " + name)
		if len(args) != 2 && i != len(args)-1 {
			fmt.Print(",")
		}
		if i == len(args)-2 {
			fmt.Print(" and")
		}
	}
	fmt.Println("!")
}

func init() {
	rootCmd.AddCommand(goodbyeCmd)
}
