package ex17

import "golang.org/x/net/html"

func ElementsByTagName(n *html.Node, names ...string) (nodes []*html.Node) {
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, n)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := ElementsByTagName(c, names...)
		nodes = append(nodes, child...)
	}
	return nodes
}
