package main

type Stack interface {
	Push(val interface{})
	Pop() interface{}
	Size() int
	Peek() interface{}
}

func NewStack() Stack {
	return &stack{}
}

// Don't use slice to store value internally, since scaling slice requires full copy
type stack struct {
	top  *element
	size int
}

type element struct {
	val  interface{}
	prev *element
	next *element
}

func (s *stack) Push(val interface{}) {
	e := &element{
		val:  val,
		prev: s.top,
		next: nil,
	}
	if s.size == 0 {
		s.size++
		s.top = e
		return
	}
	s.size++
	s.top.next = e
	s.top = e
	return
}

func (s *stack) Pop() interface{} {
	top := s.top
	if top == nil {
		return nil
	}
	s.size--
	s.top = top.prev
	if s.top != nil {
		s.top.next = nil
	}
	return top.val
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) Peek() interface{} {
	return s.top.val
}

var pairMap = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func isOpen(s string) bool {
	return s == "{" || s == "[" || s == "("
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := NewStack()
	for _, c := range s {
		cs := string(c)
		if isOpen(cs) {
			stack.Push(cs)
			continue
		}
		if stack.Size() == 0 {
			return false
		}
		ocs := stack.Pop()
		if pairMap[ocs.(string)] != cs {
			return false
		}
	}
	return stack.Size() == 0
}
