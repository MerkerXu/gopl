package main

import "fmt"


func Push(stack []int, v int) []int {
    stack = append(stack, v)
    return stack
}

func Top(stack []int) int {
    return stack[len(stack) - 1]
}

func Pop(stack []int) []int {
    stack = stack[:len(stack)-1]
    return stack
}

func main() {
    stack := []int{}
    fmt.Printf("%v\n", stack)
    stack = Push(stack, 1)
    stack = Push(stack, 2)
    stack = Push(stack, 3)
    stack = Push(stack, 4)
    fmt.Printf("%v\n", stack)
    fmt.Printf("%v\n", Top(stack))
    stack = Pop(stack)
    stack = Pop(stack)
    fmt.Printf("%v\n", stack)
}
