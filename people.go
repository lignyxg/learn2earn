package main

import "fmt"

type People struct {
	name string
	age  int
}

func (p *People) run() {
	fmt.Println("running")
}
