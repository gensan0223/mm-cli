package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gensan0223/mm-cli/internal"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "指定したメモをエディタで編集します",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		var index = -1
		for i, m := range memos {
			if m.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Printf("ID [%d] のメモが見つかりませんでした\n", id)
			return
		}

		original := memos[index].Content

		// 一時ファイルに書き出し
		tmpDir := os.TempDir()
		tmpFile := filepath.Join(tmpDir, fmt.Sprintf("mm-memo-%d.txt", time.Now().UnixNano()))
		if err := os.WriteFile(tmpFile, []byte(original), 0644); err != nil {
			fmt.Println("一時ファイル作成失敗:", err)
			os.Exit(1)
		}

		// $EDITORで開く
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi" // デフォルト
		}

		cmdEditor := exec.Command(editor, tmpFile)
		cmdEditor.Stdin = os.Stdin
		cmdEditor.Stdout = os.Stdout
		cmdEditor.Stderr = os.Stderr

		if err := cmdEditor.Run(); err != nil {
			fmt.Println("エディタの起動に失敗:", err)
			os.Exit(1)
		}

		// 編集後の内容を読み込み
		editedBytes, err := os.ReadFile(tmpFile)
		if err != nil {
			fmt.Println("編集後のファイル読み込み失敗:", err)
			os.Exit(1)
		}

		newContent := strings.TrimSpace(string(editedBytes))
		if newContent == "" {
			fmt.Println("空の内容は保存できません")
			return
		}

		memos[index].Content = newContent

		if err := internal.SaveMemos(memos); err != nil {
			fmt.Println("保存エラー:", err)
			os.Exit(1)
		}

		fmt.Printf("メモ [%d] を更新しました\n", id)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
