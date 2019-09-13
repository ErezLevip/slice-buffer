package slice_buffer

import (
	"sync"
	"testing"
)

var pool *sync.Pool
func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return NewSliceBuffer(50000, 2)
		},
	}

}

func get() *SliceBuffer {
	return pool.Get().(*SliceBuffer)
}
func free(sb *SliceBuffer) {
	sb.Reset()
	pool.Put(sb)
}

func BenchmarkNewSliceBuffer(t *testing.B) {
	t.ReportAllocs()

	sb := NewSliceBuffer(50000, 2)
	for i := 0; i < 100000000; i++ {
		sb.Append(i)
	}
}

func BenchmarkAppend(t *testing.B) {
	t.ReportAllocs()
	s := make([]interface{}, 0)
	for i := 0; i < 100000000; i++ {
		s = append(s, i)
	}
}
