/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/gensan0223/mm-cli/types"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "指定したIDのメモを削除します",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// IDを文字列→整数に変換
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Println("正しいIDを入力してください")
			os.Exit(1)
		}

		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			os.Exit(1)
		}

		var updated []types.Memo
		found := false

		for _, m := range memos {
			if m.ID != id {
				updated = append(updated, m)
			} else {
				found = true
			}
		}

		if !found {
			fmt.Printf("ID [%d] のメモは見つかりませんでした\n", id)
			return
		}

		// ID再振り分け（任意）
		for i := range updated {
			updated[i].ID = i + 1
		}

		if err := internal.SaveMemos(updated); err != nil {
			fmt.Println("保存エラー:", err)
			os.Exit(1)
		}

		fmt.Printf("ID [%d] のメモを削除しました\n", id)
	}}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
