// 这段代码演示了如何使用堆接口构建一个整数堆。
package main

import "fmt"

// 这个示例会将一些整数插入到堆里面， 接着检查堆中的最小值，
// 之后按顺序从堆里面移除各个整数。
func main() {
	ar := make([]int, 3)
	ar[1] = 3
	fmt.Println(ar, len(ar))
}
