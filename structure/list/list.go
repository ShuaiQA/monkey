package list

import (
	"fmt"
)

// 使用双链表进行表示,一个节点含有值和前后指针
type Node struct {
	next, prev *Node
	Value      any
}

type List struct {
	root Node
	len  int
}

// 声明一个空的节点,设置初始的next prev len
func New() *List {
	l := new(List)
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) FrontNode() *Node {
	if l.Len() == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) BackNode() *Node {
	if l.Len() == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) InsertFront(val any) *Node {
	return l.insertValue(val, &l.root)
}

func (l *List) InsertLast(val any) *Node {
	return l.insertValue(val, l.root.prev)
}

func (l *List) InsertBefore(val any, at *Node) *Node {
	return l.insertAfter(&Node{Value: val}, at.prev)
}

func (l *List) InsertAfter(val any, at *Node) *Node {
	return l.insertAfter(&Node{Value: val}, at)
}

func (l *List) insertValue(val any, at *Node) *Node {
	return l.insertAfter(&Node{Value: val}, at)
}

func (l *List) Remove(n *Node) any {
	l.remove(n)
	return n.Value
}

func (l *List) MoveFront(n *Node) {
	if l.root.next == n {
		return
	}
	l.move(n, &l.root)
}

func (l *List) MoveToBack(n *Node) {
	if l.root.prev == n {
		return
	}
	l.move(n, l.root.prev)
}

// 注意修改指针都是使用的是n前缀的地方,这样方便代码的编写,防止边界if判断
// 没有基于at进行写相应的指针操作
func (l *List) insertAfter(n, at *Node) *Node {
	n.next = at.next
	n.prev = at
	n.next.prev = n
	n.prev.next = n
	l.len++
	return n
}

func (l *List) remove(n *Node) {
	n.next.prev = n.prev
	n.prev.next = n.next
	n.prev = nil
	n.next = nil
	l.len--
}

func (l *List) move(n, at *Node) {
	if n == at {
		return
	}
	l.remove(n)
	l.insertAfter(n, at)
}

func (l *List) Print() {
	c := l.root.next
	for {
		fmt.Println(c.Value)
		c = c.next
		if c == &l.root {
			break
		}
	}
}
