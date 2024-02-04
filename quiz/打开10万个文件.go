package main

import (
	"os"
	"strconv"
	"sync"
)

/**
打开10万个文件
*/

func processFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

/*
*
异步并发模式
*/
func syncProcessFile() error {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			//
			// 如果文件非常大或同时处理的文件很多，那么会存在携程堆积的情况，会有大量的携程，导致内存耗尽
			// 可以考虑携程池
			processFile("file" + strconv.Itoa(i) + ".txt")
		}(i)
	}
	wg.Wait()
	return nil
}
