/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/gensan0223/mm-cli/types"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "メモを追加します",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("メモの内容を入力してください")
			return
		}

		content := args[0]

		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		newMemo := types.Memo{
			ID:      len(memos) + 1,
			Content: content,
		}

		memos = append(memos, newMemo)

		if err := internal.SaveMemos(memos); err != nil {
			fmt.Println("保存エラー:", err)
			return
		}

		fmt.Printf("メモを保存しました: [%d] %s\n", newMemo.ID, newMemo.Content)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
