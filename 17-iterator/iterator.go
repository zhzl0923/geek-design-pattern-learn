package iterator

type Iterator interface {
	Current() interface{}
	Next()
	HasNext() bool
}

type ArrayIterator struct {
	cursor    int
	arrayList []interface{}
}

func NewArrayIterator(arrList []interface{}) *ArrayIterator {
	return &ArrayIterator{cursor: 0, arrayList: arrList}
}

func (i *ArrayIterator) HasNext() bool {
	return i.cursor != len(i.arrayList)
}

func (i *ArrayIterator) Next() {
	i.cursor++
}

func (i *ArrayIterator) Current() interface{} {
	if i.cursor >= len(i.arrayList) {
		return nil
	}
	return i.arrayList[i.cursor]
}
