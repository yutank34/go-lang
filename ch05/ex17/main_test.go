package ex17

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	input := `
	<html>
		<head>
			<title></title>
		</head>
		<body>
			<div>
				<h1></h1>
				<img />
				<div>
					<h2></h2>
				</div>
				<div>
					<h2></h2>
					<h3></h3>
				</div>
			</div>
		</body>
	</html>
	`
	doc, _ := html.Parse(strings.NewReader(input))

	tests := []struct {
		tags []string
		want []string
	}{
		{[]string{"img"}, []string{"img"}},
		{[]string{"h1", "h2", "h3"}, []string{"h1", "h2", "h2", "h3"}},
		{[]string{"h4"}, nil},
	}

	for _, test := range tests {
		output := ElementsByTagName(doc, test.tags...)
		if !reflect.DeepEqual(tagNames(output), test.want) {
			t.Errorf("unexpected value: expected: %v result: %v", test.want, tagNames(output))
		}
	}
}

func tagNames(nodes []*html.Node) (names []string) {
	for _, n := range nodes {
		names = append(names, n.Data)
	}
	return
}
