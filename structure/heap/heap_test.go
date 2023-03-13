package heap

import (
	"container/heap"
	"math/rand"
	"testing"
	"time"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 随机化进行测试
func TestHeap(t *testing.T) {
	h := &IntHeap{}
	r := &IntHeap{}
	for i := 0; i < 40; i++ {
		a := gen()
		*h = append(*h, a)
		*r = append(*r, a)
	}
	Init(h)
	heap.Init(r)
	testHR(t, h, r)
	for i := 0; i < 30; i++ {
		a := gen() % 2
		if a == 0 {
			a = gen()
			h.Push(a)
			r.Push(a)
			testHR(t, h, r)
		} else {
			a1 := h.Pop().(int)
			a2 := r.Pop().(int)
			if a1 != a2 {
				t.Error("a1 is ", a1, "a2 is ", a2)
			}
			testHR(t, h, r)
		}
	}
}

func testHR(t *testing.T, h, r *IntHeap) {
	for i := 0; i < h.Len(); i++ {
		if (*h)[i] != (*r)[i] {
			t.Error(i, *h, *r)
		}
	}
}

func gen() int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(200)
}
