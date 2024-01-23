package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// https://zhuanlan.zhihu.com/p/416054707
// https://www.cnblogs.com/huageyiyangdewo/p/17357767.html

// JustErrors illustrates the use of a Group in place of a sync.WaitGroup to
// simplify goroutine counting and error handling. This example is derived from
// the sync.WaitGroup example at https://golang.org/pkg/sync/#example_WaitGroup.
func exampleGroup_justErrors() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.qq.com/",
	}

	//var urls2 = []string{
	//	"http://www.qq.com/",
	//	"http://www.163.com/",
	//	"http://www.qq.com/",
	//}

	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
}

func main() {
	exampleGroup_justErrors()
}
