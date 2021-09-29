package main

import "fmt"

type QuickUnionUF struct {
	id []int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickUnionUF{id: id}
}

func (qf *QuickUnionUF) root(i int) int {
	for i != qf.id[i] {
		i = qf.id[i]
	}

	return i
}

func (qf *QuickUnionUF) Connected(p int, q int) bool {
	return qf.root(qf.id[p]) == qf.root(qf.id[q])
}

func (qf *QuickUnionUF) Union(p int, q int) {
	i := qf.root(p)
	j := qf.root(q)
	qf.id[i] = j
}

func (qf *QuickUnionUF) GetAll() []int {
	return qf.id
}

func main() {
	fmt.Println("Quick Find UF")
	qf := NewQuickUnionUF(10)
	fmt.Println(qf.Connected(0, 1))
	fmt.Println(qf.GetAll())
	qf.Union(0, 1)
	fmt.Println(qf.Connected(0, 1))
	fmt.Println(qf.GetAll())
	qf.Union(1, 2)
	fmt.Println(qf.GetAll())
	qf.Union(2, 9)
	fmt.Println(qf.GetAll())
	qf.Union(2, 7)
	fmt.Println(qf.GetAll())
	qf.Union(9, 3)
	fmt.Println(qf.GetAll())
}
