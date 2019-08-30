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

type DefaultDict interface {
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Delete(interface{})
}

func NewDefaultDict(defaultVal interface{}) DefaultDict {
	return &defaultDict{defaultVal: defaultVal, m: make(map[interface{}]interface{})}
}

type defaultDict struct {
	defaultVal interface{}
	m          map[interface{}]interface{}
}

func (d *defaultDict) Get(key interface{}) interface{} {
	if v, ok := d.m[key]; ok {
		return v
	}
	return d.defaultVal
}

func (d *defaultDict) Set(key interface{}, val interface{}) {
	d.m[key] = val
}

func (d *defaultDict) Delete(key interface{}) {
	delete(d.m, key)
}

func longestValidParentheses(s string) int {
	stack := NewStack()
	m := NewDefaultDict(0)

	maxLength := 0
	for idx := range s {
		switch string(s[idx]) {
		case ")":
			if stack.Size() == 0 || stack.Peek() != "(" {
				stack.Push(")")
				continue
			}
			stack.Pop()
			length := m.Get(stack.Size()).(int) + 2
			if m.Get(stack.Size()+1) != 0 {
				length += m.Get(stack.Size() + 1).(int)
				m.Set(stack.Size()+1, 0)
			}
			m.Set(stack.Size(), length)
			if length > maxLength {
				maxLength = length
			}
		case "(":
			stack.Push("(")
		}
	}
	return maxLength
}
