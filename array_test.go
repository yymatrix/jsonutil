package jsonutil

import (
	"encoding/json"

	"gopkg.in/check.v1"
)

type ArraySuite struct{}

var _ = check.Suite(&ArraySuite{})

func (s *ArraySuite) SetUpTest(c *check.C) {
}

func (s *ArraySuite) TestGetArray(c *check.C) {
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

	array1, ok1 := GetArray(dict, "/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(array1, check.NotNil)
	c.Assert(len(array1), check.Equals, 3)

	array2, ok2 := GetArray(dict, "/subobj/subarray")
	c.Assert(ok2, check.Equals, true)
	c.Assert(array2, check.NotNil)
	c.Assert(len(array2), check.Equals, 3)

	array3, ok3 := GetArray(dict, "/subobj/subsubobj/array")
	c.Assert(ok3, check.Equals, true)
	c.Assert(array3, check.NotNil)
	c.Assert(len(array3), check.Equals, 2)

	array4, ok4 := GetArray(dict, "/foo")
	c.Assert(ok4, check.Equals, false)
	c.Assert(array4, check.IsNil)

	array5, ok5 := GetArray(dict, "/foox")
	c.Assert(ok5, check.Equals, false)
	c.Assert(array5, check.IsNil)
}

func (s *ArraySuite) TestGetAt(c *check.C) {
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

	e1, ok1 := GetAt(dict, "/array", 0)
	c.Assert(ok1, check.Equals, true)
	c.Assert(e1, check.NotNil)

	obj1, ok1 := e1.(map[string]interface{})
	c.Assert(ok1, check.Equals, true)
	c.Assert(obj1, check.NotNil)
	c.Assert(len(obj1), check.Equals, 1)

	obj1_foo, ok1 := obj1["foo"]
	c.Assert(ok1, check.Equals, true)
	c.Assert(obj1_foo.(float64), check.Equals, float64(1))

	e2, ok2 := GetAt(dict, "/subobj/subarray", 2)
	c.Assert(ok2, check.Equals, true)
	c.Assert(e2, check.NotNil)
	c.Assert(e2.(float64), check.Equals, float64(3))

	e3, ok3 := GetAt(dict, "/subobj/subsubobj/array", 1)
	c.Assert(ok3, check.Equals, true)
	c.Assert(e3, check.NotNil)
	c.Assert(e3.(string), check.Equals, "world")

	e4, ok4 := GetAt(dict, "/array", 3)
	c.Assert(ok4, check.Equals, false)
	c.Assert(e4, check.IsNil)

	e5, ok5 := GetAt(dict, "/foo", 3)
	c.Assert(ok5, check.Equals, false)
	c.Assert(e5, check.IsNil)
}

func (s *ArraySuite) TestAppend(c *check.C) {
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

	ok1 := Append(dict, "/array", "xx")
	c.Assert(ok1, check.Equals, true)

	array1, ok1 := GetArray(dict, "/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(len(array1), check.Equals, 4)

	e1, ok1 := GetAt(dict, "/array", 3)
	c.Assert(ok1, check.Equals, true)
	c.Assert(e1.(string), check.Equals, "xx")

	ok2 := Append(dict, "/foo", "xx")
	c.Assert(ok2, check.Equals, false)
}

func (s *ArraySuite) TestDeleteAt(c *check.C) {
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

	array1, ok1 := GetArray(dict, "/subobj/subsubobj/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(len(array1), check.Equals, 2)

	DeleteAt(dict, "/subobj/subsubobj/array", 0)
	array1, ok1 = GetArray(dict, "/subobj/subsubobj/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(len(array1), check.Equals, 1)

	DeleteAt(dict, "/foo", 0)
	e2, ok2 := dict["foo"]
	c.Assert(ok2, check.Equals, true)
	c.Assert(e2.(float64), check.Equals, float64(1))
}

func (s *ArraySuite) TestClearArray(c *check.C) {
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

	array1, ok1 := GetArray(dict, "/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(array1, check.NotNil)
	c.Assert(len(array1), check.Equals, 3)

	ClearArray(dict, "/array")
	array1, ok1 = GetArray(dict, "/array")
	c.Assert(ok1, check.Equals, true)
	c.Assert(array1, check.NotNil)
	c.Assert(len(array1), check.Equals, 0)

	array2, ok2 := GetArray(dict, "/subobj/subsubobj/array")
	c.Assert(ok2, check.Equals, true)
	c.Assert(array2, check.NotNil)
	c.Assert(len(array2), check.Equals, 2)

	ClearArray(dict, "/subobj/subsubobj/array")
	array2, ok2 = GetArray(dict, "/subobj/subsubobj/array")
	c.Assert(ok2, check.Equals, true)
	c.Assert(array2, check.NotNil)
	c.Assert(len(array2), check.Equals, 0)

	ClearArray(dict, "/foo")
	e3, ok3 := dict["foo"]
	c.Assert(ok3, check.Equals, true)
	c.Assert(e3.(float64), check.Equals, float64(1))
}
