package main

import "fmt"

type node struct {
	value int
}

type graph struct {
	nodes, edges int
	s            []int // <= there was var here
}

func main() {
	graphCreate() // <= g wasn't used
}

func input(tname string) (number int) {
	fmt.Println("input a number of " + tname)
	fmt.Scan(&number)
	return
}

func graphCreate() (g graph) { // <= g is declared here
	g = graph{nodes:input("nodes"), edges:input("edges")} // <= name the fields
	g.s = make([]int, 100) // <= g.s is already a known name
	return
}