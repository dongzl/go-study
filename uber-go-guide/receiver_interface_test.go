package uber_go_guide

import "testing"

func TestS_Read(t *testing.T) {
	sVals := map[int]S{1: {"A"}}
	// You can only call Read using a value
	sVals[1].Read()

	// This will not compile:
	// sVals[1].Write("test")
}

func TestS_Write(t *testing.T) {
	sPtrs := map[int]*S{1: {"A"}}

	// You can call both Read and Write using a pointer
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

func TestF(t *testing.T) {
	s1Val := S1{}
	s1Ptr := &S1{}
	//s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	i.f()

	// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
	//   i = s2Val
}
