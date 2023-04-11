package storage

type Storage interface {
	GetSize() int
	Begin() StorageIterator
	End() StorageIterator
	PushBack(int)
	Clear()
}

type StorageIterator interface {
	Next()
	Prev()
	Get() int
	Set(int)
	InRange() bool
	IsEqual(iterator *StorageIterator) bool
}
