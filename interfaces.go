package main

type I interface {
	get() int
	set(int)
}

type S struct {
	value int
}

func (s *S) get() int {
	return s.value
}

func (s *S) set(v int) {
	s.value = v
}
