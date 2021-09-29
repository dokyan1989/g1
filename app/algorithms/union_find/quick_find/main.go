package main

import "fmt"

type QuickFindUF struct {
	id []int
}

func NewQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickFindUF{id: id}
}

func (qf *QuickFindUF) Connected(p int, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFindUF) Union(p int, q int) {
	pid := qf.id[p]
	qid := qf.id[q]
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}

func (qf *QuickFindUF) GetAll() []int {
	return qf.id
}

func main() {
	fmt.Println("Quick Find UF")
	qf := NewQuickFindUF(10)
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
