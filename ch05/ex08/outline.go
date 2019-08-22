package ex08

import (
	"fmt"

	"golang.org/x/net/html"
)

var depth int

func ElementById(doc *html.Node, id string) *html.Node {
	// TODO wip
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
