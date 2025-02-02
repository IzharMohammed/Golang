package main

import "fmt"

// Without generics

/* 
func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
*/

// With generics

// Method - 1: Using type constraints (int | string) to allow specific types
func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Method - 2: Using "comparable" constraint to allow any comparable type
func printSlice1[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Without Generics in struct

/* 
type stack struct{
	elements []int
}
*/

// With Generics in struct
// "any" is a type parameter that allows any type
// This allows us to create a stack of any data type

type stack[T any] struct {
	elements []T
}

func main() {
	// Example usage of generic functions
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"golang", "TS"}
	printSlice(nums)
	printSlice(names)

	/*
	// Without generics in struct
	myStack := stack{
		elements: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(myStack)
	*/

	// With generics in struct
	myStack := stack[string]{
		elements: []string{"golang", "TS", "T3Stack"},
	}
	fmt.Println(myStack)

	// printStringSlice(names)
}

/*
Why Use Generics in Go?
1. **Code Reusability**: Instead of writing multiple versions of the same function for different types, generics allow us to write one function that works with multiple types.
2. **Type Safety**: Generics provide compile-time type checking, reducing the risk of runtime errors.
3. **Improved Readability**: Reduces code duplication and makes code more maintainable.
4. **Flexibility**: Allows defining generic data structures like stacks, queues, and linked lists that can handle any data type.

How to Use Generics in Go?
- Use `[T any]` for completely flexible types.
- Use `[T comparable]` when types must support equality comparisons.
- Use `[T int | string]` to restrict allowed types explicitly.
- Apply generics in both functions and struct definitions to create flexible utilities.
*/