package vectorStorage

import "HexStringToInt64/converter/storage"

type VectorStorage struct {
	size    int
	storage []int
}

func (v *VectorStorage) GetSize() int {
	return v.size
}

func (v *VectorStorage) Begin() storage.StorageIterator {
	return &VectorIterator{
		vectorStorage: v,
		current:       0,
	}
}

func (v *VectorStorage) End() storage.StorageIterator {
	return &VectorIterator{
		vectorStorage: v,
		current:       v.size - 1,
	}
}

func (v *VectorStorage) PushBack(new int) {
	v.size += 1
	v.storage = append(v.storage, new)
}

func (v *VectorStorage) Clear() {
	v.size = 0
	v.storage = nil
}

func (v *VectorStorage) get(i int) int {
	return v.storage[i]
}

func (v *VectorStorage) inRange(i int) bool {
	if i >= 0 && i < v.size {
		return true
	}
	return false
}

func (v *VectorStorage) set(i int, new int) {
	v.storage[i] = new
}

func NewVectorStorage(size int) *VectorStorage {
	return &VectorStorage{
		size:    0,
		storage: make([]int, size),
	}
}
