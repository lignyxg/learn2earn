package learn2earn

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p *People) Run() {
	fmt.Println("running")
}

func (p *People) Talk() {
	fmt.Println("Muhahahaha")
}
