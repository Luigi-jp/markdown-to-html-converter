# Markdown to HTML コンバーター

シンプルで使いやすいMarkdownファイルをHTMLに変換するコマンドラインツールです。

## 機能

- Markdownファイルを読み込み、HTMLファイルに変換
- コマンドライン引数によるファイルパスの指定
- 一般的なMarkdown記法のサポート（見出し、リスト、テーブル、太字など）

## 要件

- [Go](https://golang.org/dl/) 1.24.2以上

## インストール

```bash
# リポジトリをクローン
git clone https://github.com/Luigi-jp/markdown-to-html-converter.git
cd markdown-to-html-converter

# 依存関係のインストール
go mod download
```

## 使い方

```bash
# コンパイルして実行
go run main.go -input=<マークダウンファイルのパス> -output=<出力するHTMLファイルのパス>

# 例
go run main.go -input=sample.md -output=output.html
```

### コマンドラインオプション

- `-input`: 変換するMarkdownファイルのパス（必須）
- `-output`: 出力先のHTMLファイルのパス（必須）

## サンプル

リポジトリ内のsample.mdファイルを使って変換を試すことができます：

```bash
go run main.go -input=sample.md -output=output.html
```

## 使用ライブラリ

このプロジェクトは以下のライブラリを使用しています：

- [gomarkdown/markdown](https://github.com/gomarkdown/markdown) - Markdown解析・変換ライブラリ
