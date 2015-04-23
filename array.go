package jsonutil

import "strings"

func GetArray(dict map[string]interface{}, keypath string) (array []interface{}, ok bool) {
	if keypath == "/" {
		return nil, false
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

		v, aok := tmp[k]
		if aok && i == lastIdx {
			array, ok = v.([]interface{})
			return
		}

		if !aok {
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

func GetAt(dict map[string]interface{}, keypath string, idx int) (element interface{}, ok bool) {
	array, ok := GetArray(dict, keypath)
	if !ok {
		return nil, false
	}

	if idx < 0 || idx > len(array)-1 {
		return nil, false
	}

	return array[idx], true
}

func Append(dict map[string]interface{}, keypath string, element interface{}) (ok bool) {
	if keypath == "/" {
		return false
	}

	pureKeyPath := strings.Trim(keypath, "/")
	keys := strings.Split(pureKeyPath, "/")
	if len(keys) == 0 {
		return false
	}

	lastIdx := len(keys) - 1
	tmp := dict
	for i, k := range keys {
		if k == "" {
			continue
		}

		v, aok := tmp[k]
		if aok && i == lastIdx {
			array, xok := v.([]interface{})
			if !xok {
				return false
			}
			tmp[k] = append(array, element)
			return true
		}

		if !aok {
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

func DeleteAt(dict map[string]interface{}, keypath string, idx int) {
	if keypath == "/" {
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

		v, aok := tmp[k]
		if aok && i == lastIdx {
			array, xok := v.([]interface{})
			if !xok {
				return
			}
			tmp[k] = append(array[:idx], array[idx+1:]...)
			return
		}

		if !aok {
			return
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return
		}
	}
	return
}

func ClearArray(dict map[string]interface{}, keypath string) {
	if keypath == "/" {
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

		v, aok := tmp[k]
		if aok && i == lastIdx {
			array, xok := v.([]interface{})
			if !xok {
				return
			}
			tmp[k] = array[0:0]
			return
		}

		if !aok {
			return
		}

		switch v.(type) {
		case map[string]interface{}:
			tmp = v.(map[string]interface{})

		default:
			return
		}
	}
	return
}
