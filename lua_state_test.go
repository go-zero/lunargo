package lunargo_test

import (
	"runtime"
	"testing"

	. "github.com/go-zero/lunargo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type LuaStateTestSuite struct {
	Suite
	state *LuaState
}

func (s *LuaStateTestSuite) SetupTest() {
	s.state = NewLuaState()
}

func (s *LuaStateTestSuite) TearDownTest() {
	runtime.GC()
}

func (s *LuaStateTestSuite) TestToNotBeNil() {
	Expect(s.state).ToNot(BeNil())
}

func (s *LuaStateTestSuite) TestRunAPieceOfCode() {
	_, err := s.state.DoString("print 'hello world from Lua'")
	Expect(err).ToNot(HaveOccurred())
}

func TestLuaStateTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(LuaStateTestSuite))
}
