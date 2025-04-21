package internal

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gensan0223/mm-cli/types"
)

func GetMemoFilePath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".mm-cli")
	os.MkdirAll(dir, os.ModePerm)
	return filepath.Join(dir, "memos.json")
}

func SaveMemos(memos []types.Memo) error {
	bytes, err := json.MarshalIndent(memos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetMemoFilePath(), bytes, 0644)
}

func LoadMemos() ([]types.Memo, error) {
	file := GetMemoFilePath()
	bytes, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return []types.Memo{}, nil // ファイルがなければ空でOK
		}
		return nil, err
	}

	var memos []types.Memo
	err = json.Unmarshal(bytes, &memos)
	if err != nil {
		return nil, err
	}
	return memos, nil
}
