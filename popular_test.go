package jsonutil

import (
	"encoding/json"

	"gopkg.in/check.v1"
)

type PopularSuite struct{}

var _ = check.Suite(&PopularSuite{})

func (s *PopularSuite) SetUpTest(c *check.C) {
}

func (s *PopularSuite) TestGet(c *check.C) {
	str := `
	{
	    "foo": 1,
	    "bar": 2,
	    "test": "Hello, world!",
	    "baz": 123.1,
	    "array": [
	        {"foo": 1},
	        {"bar": 2},
	        {"baz": 3}
	    ],
	    "subobj": {
	        "foo": 1,
	        "subarray": [1,2,3],
	        "subsubobj": {
	            "bar": 2,
	            "baz": 3,
	            "array": ["hello", "world"]
	        }
	    },
	    "bool": true
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)

	val1, ok1 := Get(dict, "/")
	c.Assert(ok1, check.Equals, true)
	c.Assert(val1, check.NotNil)

	val2, ok2 := Get(dict, "/foo")
	c.Assert(ok2, check.Equals, true)
	c.Assert(val2, check.NotNil)
	c.Assert(val2.(float64), check.Equals, float64(1))

	val3, ok3 := Get(dict, "/test")
	c.Assert(ok3, check.Equals, true)
	c.Assert(val3, check.NotNil)
	c.Assert(val3.(string), check.Equals, "Hello, world!")

	val4, ok4 := Get(dict, "/subobj/subsubobj/bar")
	c.Assert(ok4, check.Equals, true)
	c.Assert(val4, check.NotNil)
	c.Assert(val4.(float64), check.Equals, float64(2))

	val5, ok5 := Get(dict, "/aa")
	c.Assert(ok5, check.Equals, false)
	c.Assert(val5, check.IsNil)

	val6, ok6 := Get(dict, "/subobj/subsubobj/bar/xx")
	c.Assert(ok6, check.Equals, false)
	c.Assert(val6, check.IsNil)
}

func (s *PopularSuite) TestSet(c *check.C) {
	dict := make(map[string]interface{})

	err1 := Set(dict, "/", "root")
	c.Assert(err1, check.Equals, ErrInvalidKey)

	err2 := Set(dict, "/A", "a")
	c.Assert(err2, check.IsNil)
	c.Assert(len(dict), check.Equals, 1)

	err3 := Set(dict, "/A", 22)
	c.Assert(err3, check.IsNil)

	valA, ok := dict["A"]
	c.Assert(ok, check.Equals, true)
	c.Assert(valA, check.Equals, 22)

	err4 := Set(dict, "/A/A2", "a2")
	c.Assert(err4, check.Equals, ErrInvalidJsonObject)
	c.Assert(len(dict), check.Equals, 1)

	err5 := Set(dict, "/B/B2", "b2")
	c.Assert(err5, check.IsNil)
	c.Assert(len(dict), check.Equals, 2)

	ok = IsExist(dict, "/B/B2")
	c.Assert(ok, check.Equals, true)
}

func (s *PopularSuite) TestTrySet(c *check.C) {
	str := `
	{
	    "foo": 1,
	    "bar": 2,
	    "test": "Hello, world!",
	    "baz": 123.1,
	    "array": [
	        {"foo": 1},
	        {"bar": 2},
	        {"baz": 3}
	    ],
	    "subobj": {
	        "foo": 1,
	        "subarray": [1,2,3],
	        "subsubobj": {
	            "bar": 2,
	            "baz": 3,
	            "array": ["hello", "world"]
	        }
	    },
	    "bool": true
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)

	err1 := TrySet(dict, "/foo", "www")
	c.Assert(err1, check.IsNil)

	val1, ok1 := Get(dict, "/foo")
	c.Assert(ok1, check.Equals, true)
	c.Assert(val1.(string), check.Equals, "www")

	err2 := TrySet(dict, "/subobj/subsubobj/bar", "good")
	c.Assert(err2, check.IsNil)

	val2, ok2 := Get(dict, "/subobj/subsubobj/bar")
	c.Assert(ok2, check.Equals, true)
	c.Assert(val2.(string), check.Equals, "good")

	err3 := TrySet(dict, "/aaaa", "www")
	c.Assert(err3, check.Equals, ErrInvalidKeyPath)

	err4 := TrySet(dict, "/subobj/subsubobj/bar/xxx", "www")
	c.Assert(err4, check.Equals, ErrInvalidKeyPath)
}

func (s *PopularSuite) TestDelete(c *check.C) {
	str := `
	{
	    "foo": 1,
	    "bar": 2,
	    "test": "Hello, world!",
	    "baz": 123.1,
	    "array": [
	        {"foo": 1},
	        {"bar": 2},
	        {"baz": 3}
	    ],
	    "subobj": {
	        "foo": 1,
	        "subarray": [1,2,3],
	        "subsubobj": {
	            "bar": 2,
	            "baz": 3,
	            "array": ["hello", "world"]
	        }
	    },
	    "bool": true
	}
	`

	dict := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &dict)
	c.Assert(err, check.IsNil)
	c.Assert(len(dict), check.Equals, 7)

	Delete(dict, "/foo")
	c.Assert(len(dict), check.Equals, 6)

	_, ok := dict["foo"]
	c.Assert(ok, check.Equals, false)

	sub, err := Subobj(dict, "/subobj/subsubobj")
	c.Assert(err, check.IsNil)
	c.Assert(sub, check.NotNil)
	c.Assert(len(sub), check.Equals, 3)

	Delete(dict, "/subobj/subsubobj/array")
	c.Assert(len(sub), check.Equals, 2)

	_, ok = sub["array"]
	c.Assert(ok, check.Equals, false)
}
