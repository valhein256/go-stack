package main

type Stack struct {
	items []item
}

func (s *Stack) Add(item Item) {
	s.items = append(s, item)
}

func (s *Stack) Pop() {

}

type Item struct {
	value int
}

func main() {
	s := Stack{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	println(s)
}
