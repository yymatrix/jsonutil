package jsonutil

import (
	"encoding/json"
	"log"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type KeypathSuite struct{}

var _ = check.Suite(&KeypathSuite{})

func (s *KeypathSuite) SetUpTest(c *check.C) {
}

func (s *KeypathSuite) TestParentKeypath(c *check.C) {
	p1 := ParentKeypath("/")
	c.Assert(p1, check.Equals, "/")

	p2 := ParentKeypath("/A")
	c.Assert(p2, check.Equals, "/")

	p3 := ParentKeypath("/A/A2")
	c.Assert(p3, check.Equals, "/A")

	p4 := ParentKeypath("/A/A2/A3")
	c.Assert(p4, check.Equals, "/A/A2")
}

func (s *KeypathSuite) TestBaseKey(c *check.C) {
	b1 := BaseKey("/")
	c.Assert(b1, check.Equals, "/")

	b2 := BaseKey("/A")
	c.Assert(b2, check.Equals, "A")

	b3 := BaseKey("/A/A2")
	c.Assert(b3, check.Equals, "A2")

	b4 := BaseKey("/A/A2/A3")
	c.Assert(b4, check.Equals, "A3")
}

func (s *KeypathSuite) TestIsExist(c *check.C) {
	str := `
	{
		"A" : "1",
		"B" : {
			"B1" : "b1",
			"B2" : "b2"
		}
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)

	ok1 := IsExist(dict, "/")
	c.Assert(ok1, check.Equals, true)

	ok2 := IsExist(dict, "/A")
	c.Assert(ok2, check.Equals, true)

	ok3 := IsExist(dict, "/B/B1")
	c.Assert(ok3, check.Equals, true)

	ok4 := IsExist(dict, "/A/A1")
	c.Assert(ok4, check.Equals, false)

	ok5 := IsExist(dict, "/B/B1/B2")
	c.Assert(ok5, check.Equals, false)
}

func (s *KeypathSuite) TestSub(c *check.C) {
	str := `
	{
		"A" : "1",
		"B" : {
			"B1" : "b1",
			"B2" : "b2"
		}
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)

	log.Printf("%#v\n", dict)

	sub, e := Subobj(dict, "/A")
	c.Assert(e, check.Equals, ErrInvalidJsonObject)
	c.Assert(sub, check.IsNil)

	sub, e = Subobj(dict, "/B")
	c.Assert(e, check.IsNil)
	c.Assert(sub, check.NotNil)

	sub, e = Subobj(dict, "/C")
	c.Assert(e, check.Equals, ErrInvalidKeyPath)
	c.Assert(sub, check.IsNil)
}

func (s *KeypathSuite) TestMkSubobj(c *check.C) {
	str := `
	{
		"A" : "1",
		"B" : {
			"B1" : "b1",
			"B2" : "b2"
		}
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)

	sub, err := MkSubobj(dict, "/A")
	c.Assert(err, check.Equals, ErrInvalidJsonObject)
	c.Assert(sub, check.IsNil)

	sub2, err2 := MkSubobj(dict, "/B")
	c.Assert(err2, check.IsNil)
	c.Assert(sub2, check.NotNil)
	c.Assert(len(sub2), check.Equals, 2)

	sub3, err3 := MkSubobj(dict, "/C")
	c.Assert(err3, check.IsNil)
	c.Assert(sub3, check.NotNil)
	c.Assert(len(sub3), check.Equals, 0)

	sub4, err4 := MkSubobj(dict, "/D/DD")
	c.Assert(err4, check.IsNil)
	c.Assert(sub4, check.NotNil)
	c.Assert(len(sub4), check.Equals, 0)
}
