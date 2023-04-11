package vectorStorage

import (
	"HexStringToInt64/converter/storage"
	"fmt"
)

type VectorIterator struct {
	vectorStorage *VectorStorage
	current       int
}

func (v *VectorIterator) Get() int {
	if v.vectorStorage.inRange(v.current) {
		return v.vectorStorage.get(v.current)
	}
	panic(fmt.Errorf("INTERNAL ERROR:VectorStorage[%d] out of range. Storage size = %d",
		v.current, v.vectorStorage.GetSize()))
}

func (v *VectorIterator) Set(new int) {
	if v.vectorStorage.inRange(v.current) {
		v.vectorStorage.set(v.current, new)
		return
	}
	panic(fmt.Errorf("INTERNAL ERROR:VectorStorage[%d] out of range. Storage size = %d",
		v.current, v.vectorStorage.GetSize()))
}

func (v *VectorIterator) InRange() bool {
	if v.vectorStorage.inRange(v.current) {
		return true
	}
	return false
}

func (v1 *VectorIterator) IsEqual(v2 *storage.StorageIterator) bool {
	v, ok := (*v2).(*VectorIterator)
	if !ok {
		return false
	}
	if *v == *v1 {
		return true
	}
	return false
}

func (v *VectorIterator) Next() {
	v.current += 1
}

func (v *VectorIterator) Prev() {
	v.current -= 1
}
