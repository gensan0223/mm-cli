# mm-cli

🎯 シンプル・高速な CLI メモ管理ツール  
Golang + Cobra で開発した、ターミナル完結のメモアプリです。

---

## 🚀 機能一覧

- 📥 `add`: メモを追加
- 📋 `list`: メモを一覧表示（タグで絞り込み可能）
- 🗃 `edit`: メモを `$EDITOR` で編集
- 🧹 `delete`: メモを削除
- 🏷 `tag add/remove/list`: タグの追加・削除・一覧
- 🔍 `search`: キーワード or タグで検索

---

## 🖥 使用例

```bash
mm add "Dockerのコマンドまとめ"
mm tag add 1 golang
mm list
mm search --tag golang
mm edit 1
mm delete 1
```

## 📁 データ保存先
メモは ~/.mm-cli/memos.json に保存されます

JSON形式なので、ポータブルで管理も楽

⚙ 開発環境・技術スタック
Go 1.21+

spf13/cobra - CLIフレームワーク

fatih/color - ターミナル色出力

## 🌱 こんな人におすすめ
ターミナルだけでメモ管理したい人

Gitログや開発メモをすぐ取りたい人

Go + CLI開発を学びたい人（ポートフォリオにも！）

## 🛠 今後のアイデア（実装予定）
 export 機能（Markdown / CSV）

 pin / archive 機能

 list --group-by-tag 表示

 bubbletea による TUI UI

## ✍️ 開発者
Author: @gensan0223

Feel free to fork or contribute!

## 🛠 インストール方法

### 方法①：Goがインストールされている場合（推奨）

```bash
go install github.com/gensan0223/mm-cli@latest
```
その後、以下で確認：

```bash
mm-cli --help
```
### 方法②：エイリアスで短縮
```bash
alias mm=mm-cli
```

### 方法③：バイナリから直接使う（Go不要）

GitHub Releases から OS に応じた実行ファイルをダウンロードしてください：

👉 [ダウンロードはこちら](https://github.com/gensan0223/mm-cli/releases)

```bash
chmod +x mm-linux-amd64
./mm-linux-amd64 list
```
