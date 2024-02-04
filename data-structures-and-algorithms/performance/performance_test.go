package performance

import (
	"strconv"
	"testing"
)

type person struct {
	name string
	age  int
}

func Benchmark_PointerWithValue(b *testing.B) {
	b.StopTimer()

	// 初始化数据
	persons := make([]person, 10000)
	for i := range persons {
		persons[i] = person{
			name: strconv.Itoa(i),
			age:  i,
		}
	}

	b.StartTimer()

	clonedPersons := make([]person, 10000)
	for n := 0; n < b.N; n++ {
		// 切片存储结构体的值

		for i := range persons {
			clonedPersons[i] = persons[i]
		}
	}
}
