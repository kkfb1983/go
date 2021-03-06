package main

import (
	"demo/src"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

type Api struct {
	Message           string `json:"message"`
	Documentation_url string `json:"documentation_url"`
}

func main() {
	//for _, arg := range os.Args[1:]{
	//
	//}
	/** add func */
	//var c int
	// c = add(3, 8)
	//src.Test()
	//src.Slics()
	//src.Dedup()
	//var v = []int{1, 3, 2}
	//src.Treesort(v)]
	/** Json And Unjson */
	//data, err := json.Marshal(src.Movies)
	//if err != nil{
	//	log.Fatalf("JSON marshaling failed: %s", err)
	//}
	//fmt.Printf("%s\n", data)
	//var titles []struct{Title string}
	//if err := json.Unmarshal(data, &titles); err != nil{
	//	log.Fatalf("JSON unmarshaling failed: %s", err)
	//}
	//fmt.Println(titles)
	/** http json */
	//result, err := src.Github(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
	//param := src.Add(5, 15)
	//fmt.Println("hello word", param)
	// go fmt.Println("hello word", param)
	/** tempconv func */
	//var c src.Celsius
	//var f src.Fahrenheit
	//c = int(src.CTof(12.13))
	//fmt.Println(c)
	//fmt.Printf("%g\n",src.BoilingC-src.FreezingC)
	//boilingF := src.CTof(src.BoilingC)
	//fmt.Printf("%g\n",boilingF-src.CTof(src.FreezingC))
	//fmt.Printf("%g\n",boilingF-src.FreezingC)
	//fmt.Println(c==0)
	//fmt.Println(f>=0)
	////fmt.Println(c==f)
	//fmt.Println(c == src.Celsius(f))
	//c := src.FToC(212.0)
	//fmt.Println(c.String())
	//fmt.Printf("%v\n", c)
	//fmt.Printf("%s\n", c)
	//fmt.Println(c)
	//fmt.Printf("%g\n", c)
	//fmt.Println(float64(c))
	//c := src.CTok(1)
	//fmt.Println(c)
	//fmt.Println(float64(c))
	//fmt.Println(src.PopCount(9))

	//getUrl()
	fmt.Println("\n#######################\n")
	//src.Squares()
	//src.Topsort()
	//breadthFirst(crawl, os.Args[1:])
	//src.Title("http://www.baidu.com")
	//src.Title(os.Args[1])
	//defer src.Trace("Title")
	// ...lots of work…
	//time.Sleep(5 * time.Second) // simulate slow
	//var double_val int = src.Deuble(4)
	//fmt.Printf("res d:%d",double_val)
	//src.Test()
	//defer printStack()
	//f(3)
	//src.Soletitle(os.Args[1])
	fmt.Println("\n########### interface ############\n")
	src.InterfaceDemo()
	//src.Interface()

}

func getUrl() {
	fmt.Println("\n#######################\n")
	url := src.IssuesUrl + "q=decoder"
	resp, _ := http.Get(url)
	//s := Api{}
	s := make(map[string]string)

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &s)

	//fmt.Println(string(body))
	for key, val := range s {
		fmt.Println(key+" ->", val+"\n")
	}

}
func crawl(url string) []string {
	fmt.Println(url)
	list, err := src.Links(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
