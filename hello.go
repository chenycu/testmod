package testmod

import "fmt"

// Hi 我哈哈哈
func Hi(name string) string {
	fmt.Println("我是master修改的")
	return fmt.Sprintf("Hi, %s", name)
}
