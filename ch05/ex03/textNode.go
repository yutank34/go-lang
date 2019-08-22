package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getTextNodes: %v\n", err)
		os.Exit(1)
	}

	texts := getTextNodes(nil, doc)

	for _, t := range texts {
		fmt.Println(t)
	}

}

func getTextNodes(texts []string, n *html.Node) []string {

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}
		texts = getTextNodes(texts, c)
	}
	return texts
}
