package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/spf13/cobra"
)

var tagFilter string

var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "メモの中からキーワードで検索します",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			os.Exit(1)
		}

		var results []string

		// タグ検索優先
		if tagFilter != "" {
			tag := strings.ToLower(tagFilter)
			for _, memo := range memos {
				for _, t := range memo.Tags {
					if strings.ToLower(t) == tag {
						results = append(results, memo.Format())
						break
					}
				}
			}
		} else if len(args) == 1 {
			keyword := strings.ToLower(args[0])
			for _, memo := range memos {
				if strings.Contains(strings.ToLower(memo.Content), keyword) {
					results = append(results, fmt.Sprintf("[%d] %s", memo.ID, memo.Content))
				}
			}
		} else {
			fmt.Println("キーワードか --tag のどちらかを指定してください")
			return
		}

		if len(results) == 0 {
			fmt.Println("条件に一致するメモは見つかりませんでした")
			return
		}

		for _, r := range results {
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&tagFilter, "tag", "", "タグで検索します")
}
