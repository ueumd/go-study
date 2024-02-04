package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func PoolTest(w io.Writer, key, val string) {
	b, _ := bufPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().UTC().Format("2006-01-02 15:04:05"))
	b.WriteByte('|')

	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)

	w.Write(b.Bytes())
	w.Write([]byte("\n"))

	bufPool.Put(b)
}

func main() {
	PoolTest(os.Stdout, "dablelv", "monkey")
}
