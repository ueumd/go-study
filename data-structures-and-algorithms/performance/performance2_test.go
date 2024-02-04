package performance

import (
	"strconv"
	"testing"
)

type person2 struct {
	name string
	age  int
}

func Benchmark_PointerWithValue2(b *testing.B) {
	b.StopTimer()

	// 初始化数据
	persons := make([]person2, 10000)
	for i := range persons {
		persons[i] = person2{
			name: strconv.Itoa(i),
			age:  i,
		}
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		// 切片存储结构体的指针
		clonedPersons := make([]*person2, 10000)
		for i := range persons {
			clonedPersons[i] = &(persons[i])
		}
	}
}
