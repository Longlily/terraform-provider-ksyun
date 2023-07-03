// this package will handle mainly that convert interface{} to anything

package ksyun

import "fmt"

func If2Slice(any interface{}) (ret []interface{}, err error) {
	if any == nil {
		return ret, err
	}
	switch any.(type) {
	case []interface{}:
		return any.([]interface{}), nil
	default:
		return nil, fmt.Errorf("got a unexpected type, expecte []interface{}")
	}
}

func If2Map(any interface{}) (ret map[string]interface{}, err error) {
	if any == nil {
		return ret, err
	}
	switch any.(type) {
	case map[string]interface{}:
		ret = any.(map[string]interface{})
		return
	default:
		err = fmt.Errorf("got a unexpected type, expecte map[string]interface{}")
	}
	return ret, err
}

func If2String(any interface{}) (ret string, err error) {
	if any == nil {
		return ret, err
	}
	switch any.(type) {
	case string:
		ret = any.(string)
		return
	default:
		err = fmt.Errorf("got a unexpected type, expecte string")
	}
	return
}
