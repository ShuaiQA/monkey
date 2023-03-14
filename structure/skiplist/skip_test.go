package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	skip := Constructor()
	skip.Add(10)
	skip.Add(20)
	skip.Add(30)
	skip.Add(40)
	fmt.Println(skip)

}
