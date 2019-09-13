package slice_buffer

type SliceBuffer struct {
	data       []interface{}
	last       int
	multiplier int
}

func NewSliceBuffer(initSize, multiplier int) *SliceBuffer {
	return &SliceBuffer{
		last:       0,
		data:       make([]interface{}, initSize),
		multiplier: multiplier,
	}
}

func (sb *SliceBuffer) Reset() {
	for i := range sb.data {
		sb.data[i] = nil
	}
	sb.last = 0
}
func (sb *SliceBuffer) Append(items ...interface{}) {
	if items == nil || len(items) == 0 {
		return
	}

	if len(items)+sb.last > len(sb.data) {
		sb.resize(sb.calcResize(len(items)))
	}

	for i, item := range items {
		sb.data[i+sb.last] = item
	}
	sb.last += len(items)
}

func (sb *SliceBuffer) resize(multiplier int) {
	res := make([]interface{}, len(sb.data)*multiplier)
	for i, v := range sb.data {
		res[i] = v
	}
	sb.data = res
}

func (sb *SliceBuffer) Slice() []interface{} {
	return sb.data[:sb.last+1]
}

func (sb *SliceBuffer) Length() int {
	return sb.last + 1
}

func (sb *SliceBuffer) calcResize(lenItems int) int {
	lenData := len(sb.data)
	targetSize := lenItems + sb.last
	m := targetSize / lenData
	if targetSize%lenData > 0 {
		m++
	}

	if m < sb.multiplier {
		m = sb.multiplier
	}
	return m
}
