package cmd

import (
	"fmt"
	"strings"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "保存されたメモを一覧表示します",
	Run: func(cmd *cobra.Command, args []string) {
		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		if len(memos) == 0 {
			fmt.Println("まだメモがありません")
			return
		}

		tag := strings.ToLower(tagFilter)
		for _, memo := range memos {
			if tag == "" {
				fmt.Println(memo.Format())
			} else {
				for _, t := range memo.Tags {
					if strings.ToLower(t) == tag {
						fmt.Println(memo.Format())
						break
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
