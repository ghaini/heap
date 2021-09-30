package heap

type HeapInterface interface {
	GetRootValue() (int, error)
	PopRootNode() (int, error)
	InsertNode(val int) error
	DeleteNode(node int) error

	getParentNode(node int) int
	getLeftChild(node int) int
	getRightChild(node int) int
	swap(from, to int)
	heapify(node int)

	ToString() string
}
