package main

import (
	"custom-structures/utils"
	"fmt"
)

func main() {
	list := utils.LinkedList[string]{}

	//add first value to empty utils
	fmt.Println(list.Size) //should be 0
	utils.Push(&list, "hello")

	//add second value to utils
	fmt.Println(list.Size) //should be 1
	utils.Push(&list, "world")
	fmt.Println(list.Size) //should be 2

	fmt.Println(utils.Pop(&list)) //should pop "world"
	fmt.Println(list.Size)        // should be 1

	fmt.Println(utils.Pop(&list)) //should pop "hello"
	fmt.Println(list.Size)        // should be 0

	fmt.Println(utils.Pop(&list)) //TODO find a way to throw error in this case
}
