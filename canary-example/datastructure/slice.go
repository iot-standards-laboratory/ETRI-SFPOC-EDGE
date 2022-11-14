package main

import "fmt"

func RemoveItemFromList(list []int, item int) []int {
	for i, e := range list {
		if e == item {
			list[i] = list[len(list)-1]
			return list[:len(list)-1]
		}
	}
	return list
}
func main() {
	list := []int{1, 2, 3, 4, 5}
	list = RemoveItemFromList(list, 1)

	fmt.Println(list)
	fmt.Println(cap(list))
	fmt.Println(len(list))
}
