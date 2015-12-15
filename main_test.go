package main

import (
	"testing"

	. "github.com/pingcap/check"
)

func TestMain(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testMainSuite{})

type testMainSuite struct {
}

func (t *testMainSuite) TestHehe(c *C) {
	c.Assert(hehe(), Equals, 2)
}
