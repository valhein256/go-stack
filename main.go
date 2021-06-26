package main

type Stack struct {
	items []int
}

func (s *Stack) Add(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	last_item_index := len(s.items)
	if last_item_index == 0 {
		return -1
	}
	item, items := s.items[last_item_index-1], s.items[:last_item_index-1]
	s.items = items
	return item
}

type Item struct {
	value int
}

func main() {
	s := Stack{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	for i, e := range s.items {
		println(i, e)
	}
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
