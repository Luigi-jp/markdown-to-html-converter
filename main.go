package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
	"os"
)

func main() {
	// 引数を受け取る
	args := os.Args

	// 引数のバリデーション
	if len(args) < 3 {
		fmt.Fprint(os.Stderr, "エラー: 引数を正しく指定してください。")
		os.Exit(1)
	}

	input, output := args[1], args[2]

	md, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v", err)
		os.Exit(1)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	html := markdown.Render(doc, renderer)

	err = os.WriteFile(output, html, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%sをHTMLに変換して%sに出力しました。", input, output)
}
