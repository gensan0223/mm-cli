package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "メモにタグを追加・管理します",
}

var tagAddCmd = &cobra.Command{
	Use:   "add [id] [tag]",
	Short: "指定したメモにタグを追加します",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Println("正しいIDを入力してください")
			os.Exit(1)
		}

		tag := strings.ToLower(args[1])

		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			os.Exit(1)
		}

		found := false
		for i, m := range memos {
			if m.ID == id {
				// 重複チェック
				for _, t := range m.Tags {
					if t == tag {
						fmt.Printf("ID [%d] のメモには既にタグ「%s」があります\n", id, tag)
						return
					}
				}
				memos[i].Tags = append(memos[i].Tags, tag)
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("ID [%d] のメモが見つかりませんでした\n", id)
			return
		}

		if err := internal.SaveMemos(memos); err != nil {
			fmt.Println("保存エラー:", err)
			os.Exit(1)
		}

		fmt.Printf("ID [%d] にタグ「%s」を追加しました\n", id, tag)
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove [id] [tag]",
	Short: "指定したメモからタグを削除します",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Println("正しいIDを入力してください")
			os.Exit(1)
		}

		removeTag := strings.ToLower(args[1])

		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			os.Exit(1)
		}

		found := false
		for i, m := range memos {
			if m.ID == id {
				newTags := []string{}
				removed := false
				for _, t := range m.Tags {
					if strings.ToLower(t) == removeTag {
						removed = true
						continue
					}
					newTags = append(newTags, t)
				}
				if removed {
					memos[i].Tags = newTags
					found = true
				} else {
					fmt.Printf("ID [%d] にタグ「%s」は存在しませんでした\n", id, removeTag)
					return
				}
				break
			}
		}

		if !found {
			fmt.Printf("ID [%d] のメモが見つかりませんでした\n", id)
			return
		}

		if err := internal.SaveMemos(memos); err != nil {
			fmt.Println("保存エラー:", err)
			os.Exit(1)
		}

		fmt.Printf("ID [%d] からタグ「%s」を削除しました\n", id, removeTag)
	},
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "すべてのメモに含まれるタグを一覧表示します",
	Run: func(cmd *cobra.Command, args []string) {
		memos, err := internal.LoadMemos()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		tagSet := map[string]bool{}

		for _, memo := range memos {
			for _, tag := range memo.Tags {
				tagSet[tag] = true
			}
		}

		if len(tagSet) == 0 {
			fmt.Println("タグはまだありません")
			return
		}

		fmt.Println("使われているタグ一覧:")

		for tag := range tagSet {
			fmt.Println(" -", tag)
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(tagAddCmd)
	tagCmd.AddCommand(tagRemoveCmd)
	tagCmd.AddCommand(tagListCmd)
}
