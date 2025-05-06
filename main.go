package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"flag"
	"fmt"
	"os"
)

func main() {
	// args := os.Args

	// if len(args) < 3 {
	// 	fmt.Fprint(os.Stderr, "エラー: 引数を正しく指定してください。")
	// 	os.Exit(1)
	// }

	// input, output := args[1], args[2]

	input := flag.String("input", "", "入力となるMarkdownファイルパス")
	output := flag.String("output", "", "HTMLの出力先ファイルパス")
	flag.Parse()

	if *input == "" || *output == "" {
		flag.Usage()
		os.Exit(1)
	}

	md, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v", err)
		os.Exit(1)
	}

	html := convertMdToHtml(md)

	err = os.WriteFile(*output, html, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%sをHTMLに変換して%sに出力しました。", *input, *output)
}

func convertMdToHtml(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	htmlData := markdown.Render(doc, renderer)
	return htmlData
}
