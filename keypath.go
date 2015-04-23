package jsonutil

import (
	"path/filepath"
	"strings"
)

func ParentKeypath(keypath string) string {
	s := filepath.Dir(keypath)
	return filepath.ToSlash(s)
}

func BaseKey(keypath string) string {
	s := filepath.Base(keypath)
	return filepath.ToSlash(s)
}

func IsExist(dict map[string]interface{}, keypath string) (ok bool) {
	if keypath == "/" {
		return true
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return true
	}

	lastIdx := len(keys) - 1
	tmp := dict
	for i, k := range keys {
		if k == "" {
			continue
		}

		v, ok := tmp[k]
		if ok && i == lastIdx {
			return true
		}

		if !ok {
			return false
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return false
		}
	}
	return false
}

func Subobj(dict map[string]interface{}, keypath string) (sub map[string]interface{}, err error) {
	if keypath == "/" {
		return dict, nil
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return dict, nil
	}

	tmp := dict
	for _, k := range keys {
		if k == "" {
			continue
		}

		v, ok := tmp[k]
		if !ok {
			return nil, ErrInvalidKeyPath
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return nil, ErrInvalidJsonObject
		}
	}
	return tmp, nil
}

func MkSubobj(dict map[string]interface{}, keypath string) (sub map[string]interface{}, err error) {
	if keypath == "/" {
		return dict, nil
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return dict, nil
	}

	tmp := dict
	for _, k := range keys {
		if k == "" {
			continue
		}

		v, ok := tmp[k]
		if !ok {
			tmpobj := make(map[string]interface{})
			tmp[k] = tmpobj
			tmp = tmpobj
			continue
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return nil, ErrInvalidJsonObject
		}
	}
	return tmp, nil
}
