package list

import (
	"testing"
)

type Integer int

func TestNew(t *testing.T) {
	l := New()
	a := l.InsertFront(10)
	b := l.InsertFront(11)
	c := l.InsertFront(12)
	d := l.InsertFront(13)
	l.Print()
	if l.Len() != 4 {
		t.Error("size error")
	}
	l.MoveFront(c)
	l.MoveToBack(a)
	l.Print()
	l.Remove(d)
	l.Print()
	l.InsertAfter(20, b)
	l.Print()
}
