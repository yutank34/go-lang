package main

import (
	"bytes"
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

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {

			if n.FirstChild != nil {

				attr := getAttr(n.Attr)
				fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attr)

			}
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// func startElement(n *html.Node) {
// 	if n.Type == html.ElementNode {

// 		if n.FirstChild != nil {

// 			attr := getAttr(n.Attr)
// 			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attr)

// 		}
// 	}
// }

// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		depth--
// 		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }

func getAttr(attr []html.Attribute) string {
	var bf bytes.Buffer
	for _, a := range attr {

		bf.WriteByte(' ')

		bf.WriteString(fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	return bf.String()
}
