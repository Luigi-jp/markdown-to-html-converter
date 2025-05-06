package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"flag"
	"log"
	"os"
)

func main() {
	input := flag.String("input", "", "入力となるMarkdownファイルパス")
	output := flag.String("output", "", "HTMLの出力先ファイルパス")
	flag.Parse()

	if *input == "" || *output == "" {
		flag.Usage()
		os.Exit(1)
	}

	md, err := os.ReadFile(*input)
	if err != nil {
		log.Fatalf("ファイルの読み込みに失敗しました。: %v", err)
	}

	html := convertMdToHtml(md)

	err = os.WriteFile(*output, html, 0644)
	if err != nil {
		log.Fatalf("ファイルへの書き込みに失敗しました。: %v", err)
	}

	log.Printf("%sをHTMLに変換して%sに出力しました。", *input, *output)
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
