package stack

type Stack struct {
	data []string
}

func New() *Stack {
	return &Stack{data: []string{}}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	outputString := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return outputString
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Peek() string {
	if len(s.data) == 0 {
		return ""
	}

	return s.data[len(s.data)-1]
}
