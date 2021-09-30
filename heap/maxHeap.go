package heap

import "fmt"

type maxHeap struct {
	HeapArr  []int
	Capacity int
	Size     int
}

func (m *maxHeap) GetRootValue() (int, error) {
	if m.Size <= 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	return m.HeapArr[0], nil
}

func (m *maxHeap) PopRootNode() (int, error) {
	if m.Size <= 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	root := m.HeapArr[0]
	err := m.DeleteNode(0)
	if err != nil {
		return 0, err
	}

	return root, nil
}

func (m *maxHeap) InsertNode(val int) error {
	if m.Size == m.Capacity {
		return fmt.Errorf("overflow, heap is full")
	}

	i := m.Size
	m.Size++
	m.HeapArr[i] = val
	for i != 0 && m.HeapArr[m.getParentNode(i)] < m.HeapArr[i] {
		parent := m.getParentNode(i)
		m.swap(i, parent)
		i = parent
	}

	return nil
}

func (m *maxHeap) DeleteNode(node int) error {
	if m.Size <= node {
		return fmt.Errorf("node not exist")
	}

	lastNode := m.Size-1
	m.HeapArr[node] = m.HeapArr[lastNode]
	m.HeapArr[lastNode] = 0
	m.Size--
	m.heapify(node)
	return  nil
}

func (m *maxHeap) getParentNode(node int) int {
	return (node - 1) / 2
}

func (m *maxHeap) getLeftChild(node int) int {
	return (2 * node) + 1
}

func (m *maxHeap) getRightChild(node int) int {
	return (2 * node) + 2
}

func (m *maxHeap) swap(from, to int) {
	m.HeapArr[from], m.HeapArr[to] = m.HeapArr[to], m.HeapArr[from]
}

func (m *maxHeap) heapify(node int) {
	left := m.getLeftChild(node)
	right := m.getRightChild(node)
	largest := node
	if left < m.Size && m.HeapArr[left] > m.HeapArr[largest] {
		largest = left
	}

	if left < m.Size && m.HeapArr[right] > m.HeapArr[largest] {
		largest = right
	}

	if largest != node {
		m.swap(node, largest)
		m.heapify(largest)
	}
}

func (m *maxHeap) ToString() string {
	return fmt.Sprintf("%v", m.HeapArr)
}

func NewMaxHeap(cap int) HeapInterface {
	arr := make([]int, cap)
	return &maxHeap{
		HeapArr:  arr,
		Capacity: cap,
		Size:     0,
	}
}
