package jsonutil

import "strings"

func Get(dict map[string]interface{}, keypath string) (value interface{}, ok bool) {
	if keypath == "/" {
		return dict, true
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return nil, false
	}

	lastIdx := len(keys) - 1
	tmp := dict
	for i, k := range keys {
		if k == "" {
			continue
		}

		v, ok := tmp[k]
		if ok && i == lastIdx {
			return v, true
		}

		if !ok {
			return nil, false
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return nil, false
		}
	}
	return nil, false
}

func Set(dict map[string]interface{}, keypath string, value interface{}) (err error) {
	ppath := ParentKeypath(keypath)
	sub, e := MkSubobj(dict, ppath)
	if e != nil {
		return e
	}

	bkey := BaseKey(keypath)
	if bkey == "/" {
		return ErrInvalidKey
	}

	sub[bkey] = value
	return nil
}

func TrySet(dict map[string]interface{}, keypath string, value interface{}) (err error) {
	if IsExist(dict, keypath) == false {
		return ErrInvalidKeyPath
	}

	ppath := ParentKeypath(keypath)
	sub, e := MkSubobj(dict, ppath)
	if e != nil {
		return e
	}

	bkey := BaseKey(keypath)
	if bkey == "/" {
		return ErrInvalidKey
	}

	if _, ok := sub[bkey]; !ok {
		return ErrInvalidKey
	}

	sub[bkey] = value
	return nil
}

func Delete(dict map[string]interface{}, keypath string) {
	if keypath == "/" {
		for k, _ := range dict {
			delete(dict, k)
		}
		return
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return
	}

	lastIdx := len(keys) - 1
	tmp := dict
	for i, k := range keys {
		if k == "" {
			continue
		}

		v, ok := tmp[k]
		if ok && i == lastIdx {
			delete(tmp, k)
			return
		}

		if !ok {
			return
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return
		}
	}
}
