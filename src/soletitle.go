package src

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Soletitle(url string) (title string, err error) {
	var doc *html.Node
	type bailout struct{}
	resp, _ := http.Get(url)
	doc, _ = html.Parse(resp.Body)
	defer resp.Body.Close()
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
