package support

type Items []Value

type Slice struct {
	Items Items
}

func NewSlice(items ...interface{}) Slice {
	slice := Slice{}

	for _, item := range items {
		slice.Items = append(slice.Items, NewValue(item))
	}

	return slice
}

func (c *Slice) Push(item interface{}) *Slice {
	c.Items = append(c.Items, NewValue(item))

	return c
}

func (c Slice) Reverse() Slice {
	items := c.Items
	for left, right := 0, len(items)-1; left < right; left, right = left+1, right-1 {
		items[left], items[right] = items[right], items[left]
	}

	c.Items = items

	return c
}

func (c Slice) ToSlice() Items {
	return c.Items
}

// Determine if an item exists in the collection by a string
func (c Slice) Contains(search interface{}) bool {
	for _, item := range c.Items {
		if item.Source() == NewValue(search).Source() {
			return true
		}
	}

	return false
}
