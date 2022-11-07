package main

import (
	"fmt"
)

var parent []int

func make_set(v int) {
	parent[v] = v
}

func find_set(v int) int {
	if v == parent[v] {
		return v
	}
	return find_set(parent[v])
}

func union_set(a int, b int) {
	a = find_set(a)
	b = find_set(b)

	if a != b {
		parent[b] = a
	}
}

func initialize(parent []int, N int) {
	for i := 0; i <= N; i++ {
		parent[i] = i
	}
}

func main() {
	fmt.Println("This is Disjoint Set")
	fmt.Println("Enter maximum value no: ")
	var n int
	fmt.Scan(&n)
	initialize(parent, n)

	union_set(1, 2)
	union_set(3, 4)
	union_set(2, 3)

	var x = find_set(1)
	var y = find_set(3)
	if x == y {
		fmt.Println("ok! They are same set")
	} else {
		fmt.Println("owf!, they are not in same set")
	}

}
