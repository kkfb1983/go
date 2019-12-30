package src

import (
	"fmt"
	//"html"
	//"net/http"
	"os"
)

func FindlinkMain() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks:%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
func findlinks(url string) ([]string, error) {
	//resp, err := http.Get(url)
	//if err != nil{
	//	return nil,err
	//}
	//if resp.StatusCode != http.StatusOK {
	//	resp.Body.Close()
	//	return nil,fmt.Errorf("getting %s:%s",url,resp)
	//}
	//doc, err  := html.Parse(resp.Body)
	//resp.Body.Close()
	//if err != nil {
	//	return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	//}
	//return visit(nil, doc),nil
	return nil, nil
}
