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

	links := visit(nil, doc)

	for _, l := range links {
		fmt.Println(l)
	}
}

func visit(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

		if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
