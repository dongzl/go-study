package uber_go_guide

// https://github.com/uber-go/guide/blob/master/style.md#receivers-and-interfaces
type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}
