#Slice buffer

Slice buffer saves allocations and ops by reusing the same underlying slice.

## Installation
go get -u "github.com/erezlevip/slice-buffer"

## Quick Start
Create a slice buffer and pass initial size and a multiplier for the resize.
* the initial size will affect the resizing of the buffer but it can be pretty big because it is reused. 
````go
sb := NewSliceBuffer(50000, 2)
````

Add items using the Append method
````go
for i := 0; i < 100000000; i++ {
	sb.Append(i)
}
````

To get the slice use the Slice method:
````go
for _, item := range sb.Slice() {
	fmt.Println(item.(string))
}
````

To reset the slice buffer use the Reset:
````go
sb.Reset()
````

To get the length of the slice use the Length method:
````go
sb.Length()
````

If you want to use Sync.Pool with it:
````go
var pool *sync.Pool
func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return NewSliceBuffer(50000, 2)
		},
	}
}

func getSlice() *SliceBuffer {
	return pool.Get().(*SliceBuffer)
}
func freeSlize(sb *SliceBuffer) {
	sb.Reset() //reset before putting it back
	pool.Put(sb)
}
````

