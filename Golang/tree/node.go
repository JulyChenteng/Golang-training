package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) SetValue(value int) {
	if nil == node {
		fmt.Println("Setting Value to nil node, ignored!")
		return
	}

	node.Value = value
}

/*
 * Go 语言中没有构造和析构函数，因此一般都是通过普通函数来作为工厂函数创建结构体
 * 注意该函数返回了局部变量的地址，这在Go语言中是允许的。
 * 此时就需要考虑该局部变量是存在栈上还是堆上？
 * 内存分配的位置是有编译器和运行环境决定的，在本环境中返回局部变量和返回局部变量的地址均可。
 */
func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
